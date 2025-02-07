import psycopg2
import yaml
import os
from datetime import datetime
from decimal import Decimal

def load_config():
    """
    加载配置文件
    
    返回:
        dict: 包含PostgreSQL连接参数的字典
    """
    try:
        config_path = os.path.join(os.path.dirname(__file__), 'config.yaml')
        with open(config_path, 'r', encoding='utf-8') as file:
            config = yaml.safe_load(file)
        return config['postgresql']
    except Exception as e:
        print(f"读取配置文件时发生错误: {e}")
        return None

class PostgresqlExporter:
    def __init__(self, db_config):
        self.db_config = db_config
        self.connection = None
        self.cursor = None

    def connect(self, database=None):
        """建立数据库连接"""
        try:
            # 如果指定了数据库，就使用指定的数据库
            if database:
                self.db_config['database'] = database
            
            self.connection = psycopg2.connect(**self.db_config)
            self.cursor = self.connection.cursor()
            return True
        except Exception as e:
            print(f"连接数据库失败: {e}")
            return False

    def close(self):
        """关闭数据库连接"""
        if self.cursor:
            self.cursor.close()
        if self.connection:
            self.connection.close()
            print("数据库连接已关闭")

    def get_databases(self):
        """获取所有数据库列表"""
        try:
            # 连接到默认数据库
            self.db_config['database'] = 'postgres'
            if not self.connect():
                return []
                
            self.cursor.execute("""
                SELECT datname 
                FROM pg_database 
                WHERE datistemplate = false 
                AND datname NOT IN ('postgres', 'template0', 'template1')
                ORDER BY datname;
            """)
            return [row[0] for row in self.cursor.fetchall()]
        except Exception as e:
            print(f"获取数据库列表失败: {e}")
            return []
        finally:
            self.close()

    def get_tables(self):
        """获取所有表名"""
        try:
            self.cursor.execute("""
                SELECT table_name 
                FROM information_schema.tables 
                WHERE table_schema = 'public'
                AND table_type = 'BASE TABLE'
                ORDER BY table_name;
            """)
            return [row[0] for row in self.cursor.fetchall()]
        except Exception as e:
            print(f"获取表列表失败: {e}")
            return []

    def get_create_table_sql(self, table_name):
        """获取创建表的SQL语句"""
        try:
            # 获取列信息
            self.cursor.execute("""
                SELECT 
                    column_name, 
                    data_type,
                    character_maximum_length,
                    is_nullable,
                    column_default
                FROM information_schema.columns
                WHERE table_schema = 'public' AND table_name = %s
                ORDER BY ordinal_position;
            """, (table_name,))
            columns = self.cursor.fetchall()

            # 获取主键信息
            self.cursor.execute("""
                SELECT a.attname
                FROM pg_index i
                JOIN pg_attribute a ON a.attrelid = i.indrelid AND a.attnum = ANY(i.indkey)
                WHERE i.indrelid = %s::regclass AND i.indisprimary;
            """, (table_name,))
            primary_keys = [row[0] for row in self.cursor.fetchall()]

            # 构建CREATE TABLE语句
            sql_parts = [f"CREATE TABLE {table_name} ("]
            column_definitions = []

            for col_name, data_type, max_length, is_nullable, default in columns:
                # 处理特殊的列名（如果列名是PostgreSQL关键字或包含大写字母）
                if col_name.lower() in ['collation'] or any(c.isupper() for c in col_name):
                    col_name = f'"{col_name}"'

                # 构建列定义
                col_def = [f"{col_name} {data_type}"]
                
                # 添加长度限制（如果有）
                if max_length is not None:
                    col_def[0] = f"{col_name} {data_type}({max_length})"
                
                # 添加NULL/NOT NULL
                if is_nullable == 'NO':
                    col_def.append('NOT NULL')
                
                # 添加默认值（如果有）
                if default is not None:
                    # 处理特殊的默认值
                    if default.startswith("nextval"):
                        col_def.append(f"DEFAULT nextval('{table_name}_{col_name}_seq'::regclass)")
                    else:
                        col_def.append(f"DEFAULT {default}")
                
                column_definitions.append(" ".join(col_def))

            # 添加主键约束（如果有）
            if primary_keys:
                column_definitions.append(f"PRIMARY KEY ({', '.join(primary_keys)})")

            sql_parts.append(",\n    ".join(column_definitions))
            sql_parts.append(");")

            # 如果有序列，添加序列创建语句
            sequences_sql = []
            for col_name, data_type, max_length, is_nullable, default in columns:
                if default and default.startswith("nextval"):
                    seq_name = f"{table_name}_{col_name}_seq"
                    sequences_sql.append(f"CREATE SEQUENCE IF NOT EXISTS {seq_name};")

            if sequences_sql:
                return "\n".join(sequences_sql + [""] + ["\n".join(sql_parts)])
            return "\n".join(sql_parts)

        except Exception as e:
            print(f"获取表 {table_name} 的创建语句失败: {e}")
            return None

    def format_value(self, value):
        """格式化SQL值"""
        if value is None:
            return 'NULL'
        elif isinstance(value, (int, float, Decimal)):
            return str(value)
        elif isinstance(value, datetime):
            return f"'{value.isoformat()}'"
        elif isinstance(value, (bytes, bytearray)):
            return f"'\\x{value.hex()}'"
        else:
            # 使用更安全的字符串处理方式
            return "'" + str(value).replace("'", "''") + "'"

    def get_insert_data_sql(self, table_name):
        """获取插入数据的SQL语句"""
        try:
            # 获取列信息
            self.cursor.execute("""
                SELECT column_name 
                FROM information_schema.columns 
                WHERE table_schema = 'public' 
                AND table_name = %s 
                ORDER BY ordinal_position;
            """, (table_name,))
            columns = [row[0] for row in self.cursor.fetchall()]
            
            # 处理特殊列名和大小写
            quoted_columns = []
            for col in columns:
                if col.lower() in ['collation'] or any(c.isupper() for c in col):
                    quoted_columns.append(f'"{col}"')
                else:
                    quoted_columns.append(col)

            # 构建查询语句，使用带引号的列名
            select_cols = ', '.join(quoted_columns)
            self.cursor.execute(f"SELECT {select_cols} FROM {table_name};")
            rows = self.cursor.fetchall()

            if not rows:
                return None

            insert_statements = []
            for row in rows:
                values = [self.format_value(val) for val in row]
                insert_sql = f"INSERT INTO {table_name} ({', '.join(quoted_columns)}) VALUES ({', '.join(values)});"
                insert_statements.append(insert_sql)

            return "\n".join(insert_statements)

        except Exception as e:
            print(f"获取表 {table_name} 的数据失败: {e}")
            return None

    def get_sequences(self):
        """获取所有序列"""
        try:
            self.cursor.execute("""
                SELECT sequence_name
                FROM information_schema.sequences
                WHERE sequence_schema = 'public'
                ORDER BY sequence_name;
            """)
            return [row[0] for row in self.cursor.fetchall()]
        except Exception as e:
            print(f"获取序列列表失败: {e}")
            return []

    def get_sequence_value(self, sequence_name):
        """获取序列当前值"""
        try:
            self.cursor.execute(f"SELECT last_value, is_called FROM {sequence_name};")
            last_value, is_called = self.cursor.fetchone()
            return last_value if is_called else last_value - 1
        except Exception as e:
            print(f"获取序列 {sequence_name} 的值失败: {e}")
            return None

    def export_tables(self, database, tables, output_file):
        """导出指定的表"""
        try:
            with open(output_file, 'w', encoding='utf-8') as f:
                # 写入文件头部注释，使用标准的时间格式
                f.write(f"-- 数据库导出时间: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}\n")
                f.write(f"-- 数据库: {database}\n\n")
                
                # 写入SET语句
                f.write("SET statement_timeout = 0;\n")
                f.write("SET client_encoding = 'UTF8';\n")
                f.write("SET standard_conforming_strings = on;\n\n")

                # 导出序列
                sequences = self.get_sequences()
                if sequences:
                    f.write("-- 导出序列\n")
                    for seq_name in sequences:
                        current_value = self.get_sequence_value(seq_name)
                        if current_value is not None:
                            f.write(f"DROP SEQUENCE IF EXISTS {seq_name} CASCADE;\n")
                            f.write(f"CREATE SEQUENCE {seq_name};\n")
                            f.write(f"SELECT pg_catalog.setval('{seq_name}', {current_value}, true);\n")
                    f.write("\n")

                total_tables = len(tables)
                for i, table_name in enumerate(tables, 1):
                    print(f"正在导出表 {table_name} ({i}/{total_tables})")
                    
                    # 写入删除表语句
                    f.write(f"\n-- 删除表 {table_name}\n")
                    f.write(f"DROP TABLE IF EXISTS {table_name} CASCADE;\n\n")
                    
                    # 写入创建表语句
                    f.write(f"-- 创建表 {table_name}\n")
                    create_sql = self.get_create_table_sql(table_name)
                    if create_sql:
                        f.write(create_sql + "\n\n")
                    
                    # 写入数据
                    f.write(f"-- 插入表 {table_name} 数据\n")
                    insert_sql = self.get_insert_data_sql(table_name)
                    if insert_sql:
                        f.write(insert_sql + "\n")

            print(f"\n数据库导出成功！文件保存在: {output_file}")
            return True

        except Exception as e:
            print(f"导出数据库时发生错误: {e}")
            return False

def print_menu(options, title):
    """打印选择菜单"""
    print(f"\n{title}")
    print("=" * len(title))
    for i, option in enumerate(options, 1):
        print(f"{i}. {option}")
    print("0. 返回上级菜单")

def get_user_choice(options, prompt="请输入你的选择: "):
    """获取用户输入的选择"""
    while True:
        try:
            choice = input(prompt)
            if choice == "0":
                return None
            choice = int(choice)
            if 1 <= choice <= len(options):
                return choice - 1
            print("无效的选择，请重试")
        except ValueError:
            print("请输入数字")

def select_multiple_items(options, title):
    """选择多个项目"""
    print(f"\n{title}")
    print("=" * len(title))
    for i, option in enumerate(options, 1):
        print(f"{i}. {option}")
    print("输入序号(用逗号分隔)选择多个项目，或者输入 'all' 选择所有项目")
    print("例如: 1,3,5 或 all")
    print("输入 0 返回上级菜单")

    while True:
        choice = input("请输入你的选择: ").strip()
        if choice == "0":
            return None
        if choice.lower() == "all":
            return options
        try:
            indices = [int(x.strip()) - 1 for x in choice.split(",")]
            if all(0 <= i < len(options) for i in indices):
                return [options[i] for i in indices]
            print("无效的选择，请重试")
        except ValueError:
            print("格式错误，请重试")

def main():
    """主函数"""
    config = load_config()
    if not config:
        return

    exporter = PostgresqlExporter(config)
    
    while True:
        print("\n=== PostgreSQL数据库导出工具 ===")
        print("1. 导出指定数据库")
        print("2. 导出所有数据库")
        print("0. 退出程序")
        
        choice = input("请输入你的选择: ")
        
        if choice == "0":
            print("程序已退出")
            break
            
        elif choice == "1":
            # 获取所有数据库
            databases = exporter.get_databases()
            if not databases:
                print("没有找到任何数据库")
                continue
                
            # 选择数据库
            print_menu(databases, "选择要导出的数据库")
            db_choice = get_user_choice(databases)
            if db_choice is None:
                continue
                
            selected_db = databases[db_choice]
            
            # 连接到选中的数据库
            if not exporter.connect(selected_db):
                continue
                
            # 获取所有表
            tables = exporter.get_tables()
            if not tables:
                print("没有找到任何表")
                exporter.close()
                continue
                
            # 选择表
            selected_tables = select_multiple_items(tables, f"选择要导出的表 (数据库: {selected_db})")
            if selected_tables is None:
                exporter.close()
                continue
                
            # 生成输出文件名
            timestamp = datetime.now().strftime('%Y年%m月%d日%H时%M分%S秒')
            dumps_dir = os.path.join(os.path.dirname(__file__), 'dumps')
            os.makedirs(dumps_dir, exist_ok=True)
            output_file = os.path.join(dumps_dir, f"{selected_db}_{timestamp}.sql")
            
            # 导出选中的表
            exporter.export_tables(selected_db, selected_tables, output_file)
            exporter.close()
            
        elif choice == "2":
            # 获取所有数据库
            databases = exporter.get_databases()
            if not databases:
                print("没有找到任何数据库")
                continue
                
            # 导出每个数据库
            for db_name in databases:
                print(f"\n正在处理数据库: {db_name}")
                
                # 连接到当前数据库
                if not exporter.connect(db_name):
                    continue
                    
                # 获取所有表
                tables = exporter.get_tables()
                if not tables:
                    print(f"数据库 {db_name} 中没有找到任何表")
                    exporter.close()
                    continue
                    
                # 生成输出文件名
                timestamp = datetime.now().strftime('%Y年%m月%d日%H时%M分%S秒')
                dumps_dir = os.path.join(os.path.dirname(__file__), 'dumps')
                os.makedirs(dumps_dir, exist_ok=True)
                output_file = os.path.join(dumps_dir, f"{db_name}_{timestamp}.sql")
                
                # 导出所有表
                exporter.export_tables(db_name, tables, output_file)
                exporter.close()
                
        else:
            print("无效的选择，请重试")

if __name__ == "__main__":
    main()
