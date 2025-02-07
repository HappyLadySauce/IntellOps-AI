# PostgreSQL数据库备份工具

一个简单的PostgreSQL数据库备份工具，可以自动将数据库导出为SQL文件。

## 功能介绍

- 自动备份所有非系统数据库
- 备份文件自动添加时间戳
- 自动创建dumps目录存放备份文件
- 支持yaml配置文件管理数据库连接信息

## 使用前准备

1. 安装必要的软件：
   - Python 3.6+
   - PostgreSQL客户端工具(pg_dump)
   
2. 安装Python依赖：
```bash
pip install psycopg2-binary pyyaml
```

## 配置说明

在`config.yaml`文件中配置数据库连接信息：
```yaml
postgresql:
  host: your_host          # 数据库服务器地址
  database: your_database  # 数据库名称
  user: your_user         # 用户名
  password: your_password # 密码
  port: 5432             # 端口号
```

## 使用方法

1. 运行备份命令：
```bash
python dump-postgresql.py
```

2. 备份文件将保存在`dumps`目录下，文件名格式为：
```
数据库名_YYYYMMDD_HHMMSS.sql
```

## 注意事项

1. 确保运行环境已安装pg_dump工具
2. 确保有足够的磁盘空间
3. 确保数据库用户有备份权限
4. 定期清理旧的备份文件

## 常见问题

如果遇到错误，请检查：
- 数据库连接配置是否正确
- pg_dump命令是否可用
- 用户权限是否足够
- 磁盘空间是否充足