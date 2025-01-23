package install

import (
	"archive/zip"
	"bufio"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"gofly/utils/gf"
	"gofly/utils/gform"
	"gofly/utils/tools/gcfg"
	"gofly/utils/tools/gconv"
	"gofly/utils/tools/gctx"
	"gofly/utils/tools/grand"
	"gofly/utils/tools/gstr"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

/**
* 项目安装
 */
type Index struct{ NoNeedLogin []string }

func init() {
	fpath := Index{NoNeedLogin: []string{"*"}}
	gf.Register(&fpath, fpath)
}

// 安装页面
func (api *Index) Index(c *gf.GinCtx) {
	path, err := os.Getwd()
	if err != nil {
		gf.Failed().SetMsg("项目路径获取失败").SetData(err.Error()).Regin(c)
		return
	}
	ctx := gctx.New()
	appConf, _ := gcfg.Instance().Get(ctx, "app")
	appConf_arr := gconv.Map(appConf)
	dbConf, _ := gcfg.Instance().Get(ctx, "database.default")
	dbConf_arr := gconv.Map(dbConf)
	if appConf_arr["runEnv"] == "release" {
		gf.Failed().SetMsg("线上环境无法进行安装操作,如下安装请在在runEnv=debug环境下安装！").Regin(c)
		return
	}
	filePath := filepath.Join(path, "/devsource/developer/install/install.lock")
	if _, err := os.Stat(filePath); err == nil {
		c.HTML(http.StatusOK, "isinstall.html", gf.Map{"title": "已经安装页面"})
	} else {
		views_path := gconv.String(appConf_arr["vueobjroot"])
		if views_path == "" {
			views_path = filepath.Join(path, "/views")
		}
		c.HTML(http.StatusOK, "install.html", gf.Map{
			"title":      "安装页面",
			"views_path": gstr.Replace(views_path, "\\", "/"),
			"hostname":   dbConf_arr["hostname"],
			"hostport":   dbConf_arr["hostport"],
			"username":   dbConf_arr["username"],
			"password":   dbConf_arr["password"],
			"dbname":     dbConf_arr["dbname"],
			"prefix":     dbConf_arr["prefix"],
		})
	}

}

// 链接数据库配置
func dsn(username, password, hostname, hostport, dbName interface{}) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=1000ms&sql_mode=%s&multiStatements=true", username, password, hostname, hostport, dbName, "NO_ENGINE_SUBSTITUTION")
}

// 安装
func (api *Index) Save(c *gf.GinCtx) {
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	path, err := os.Getwd() //获取当前路径
	if err != nil {
		gf.Failed().SetMsg("项目路径获取失败").SetData(err.Error()).Regin(c)
		return
	}
	//1链接数据
	db, err := sql.Open("mysql", dsn(parameter["username"], parameter["password"], parameter["hostname"], parameter["hostport"], ""))
	if err != nil {
		gf.Failed().SetMsg(fmt.Sprintf("链接数据库Error %s when opening DB\n", err)).Regin(c)
	}
	defer db.Close()
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	//2创建数据库
	sqlstr := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci", parameter["dbname"])
	dbres, adderr := db.ExecContext(ctx, sqlstr)
	if adderr != nil {
		gf.Failed().SetMsg(fmt.Sprintf("创建数据库Error %s when creating DB\n", adderr)).Regin(c)
		return
	}
	_, rferr := dbres.RowsAffected()
	if rferr != nil {
		gf.Failed().SetMsg(fmt.Sprintf("创建数据库链接Error %s when fetching rows", rferr)).Regin(c)
		return
	}
	db.Close()
	db, err = sql.Open("mysql", dsn(parameter["username"], parameter["password"], parameter["hostname"], parameter["hostport"], parameter["dbname"]))
	if err != nil {
		gf.Failed().SetMsg(fmt.Sprintf("链接数据库%v，失败Error %s when opening DB", parameter["dbname"], err)).Regin(c)
		return
	}
	defer db.Close()
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)
	ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err = db.PingContext(ctx)
	if err != nil {
		gf.Failed().SetMsg(fmt.Sprintf("检查重链接数据库%v，失败Errors %s pinging DB", parameter["dbname"], err)).Regin(c)
		return
	}
	// db.Exec("SET @@sql_mode='NO_ENGINE_SUBSTITUTION';")
	//4导入数据库表
	//4.1导入基础数据库配置
	SqlPath := filepath.Join(path, "/devsource/developer/install/business_db.sql")
	sqls, sqlerr := os.ReadFile(SqlPath)
	if sqlerr != nil {
		gf.Failed().SetMsg("数据库文件不存在：" + SqlPath).SetData(sqlerr).Regin(c)
		return
	}
	db.Exec(string(sqls))
	// 4.2 安装admin端
	if parameter["isInstalladmin"] == "install" {
		adminSqlPath := filepath.Join(path, "/devsource/developer/install/admin_db.sql")
		adminsqls, adminsqlerr := os.ReadFile(adminSqlPath)
		if adminsqlerr != nil {
			gf.Failed().SetMsg("数据库文件不存在：" + adminSqlPath).SetData(adminsqlerr).Regin(c)
			return
		}
		db.Exec(string(adminsqls))
	}
	//5.修改后台账号
	salt := grand.Str("123456789", 6)
	businesspass := fmt.Sprintf("%v%v", gf.Md5(parameter["businessPassword"].(string)), salt)
	_, upberr := db.Exec("update business_account set username = ?, password = ?, salt = ? where id = ?", parameter["businessUsername"], gf.Md5(businesspass), salt, 1)
	if upberr != nil {
		gf.Failed().SetMsg("更新商务端后台账号密码失败：").SetData(upberr).Regin(c)
		return
	}
	//安装admin端
	if parameter["isInstalladmin"] == "install" {
		adminpass := fmt.Sprintf("%v%v", gf.Md5(parameter["adminPassword"].(string)), salt)
		_, upaerr := db.Exec("update admin_account set username = ?, password = ?, salt = ? where id = ?", parameter["adminUsername"], gf.Md5(adminpass), salt, 1)
		if upaerr != nil {
			gf.Failed().SetMsg("更新admin端后台账号密码失败：").SetData(upaerr).Regin(c)
			return
		}
	}
	//修改表前缀
	var table_prefix string = ""
	if parameter["prefix"] != "" {
		table_prefix = gconv.String(parameter["prefix"])
		rows, _ := db.Query("SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_TYPE = 'BASE TABLE' AND TABLE_SCHEMA='" + gconv.String(parameter["dbname"]) + "'")
		defer rows.Close()
		var tablename_str string
		var sql_str string
		for rows.Next() {
			rows.Scan(&tablename_str)
			sql_str = "ALTER TABLE " + tablename_str + " RENAME TO " + table_prefix + tablename_str + ";"
			_, upferr := db.Exec(sql_str)
			if upferr != nil {
				gf.Failed().SetMsg("修改表前缀失败：").SetData(upferr).Regin(c)
				return
			}
		}
	}
	//6.创建安装锁文件
	filePath := filepath.Join(path, "/devsource/developer/install/install.lock")
	os.Create(filePath)
	//7.安装前端页面
	if _, ok := parameter["vuepath"]; ok && parameter["vuepath"] != "" {
		//parameter["vueobjroot"] = filepath.Join(gconv.String(parameter["vuepath"]), "/business") //更新前端路径
		//7.1 如果没有filepath文件目录就创建一个
		file_path := fmt.Sprintf("%v", parameter["vuepath"])
		if _, err := os.Stat(file_path); err != nil {
			if !os.IsExist(err) {
				os.MkdirAll(file_path, os.ModePerm)
			}
		}
		//7.2 复制前端文件到指定位置
		vuesoure_path := filepath.Join(path, "/devsource/developer/install/vuecode/")
		CopyDir(vuesoure_path, file_path)
		//7.3 解压文件
		business_vue_path := filepath.Join(file_path, "/business.zip")
		admin_vue_path := filepath.Join(file_path, "/admin.zip")
		Unzip(business_vue_path, file_path)
		//安装admin端
		if parameter["isInstalladmin"] == "install" { //解压admin
			Unzip(admin_vue_path, file_path)
		} else { //不安装admin则把admin的后端代码删除
			app_admin_path := filepath.Join(path, "/app/admin")
			os.RemoveAll(app_admin_path)
			ChecAdminRemoveController()
		}
		//7.4 删除zip文件
		os.RemoveAll(business_vue_path)
		os.RemoveAll(admin_vue_path)
	}
	//修改配置文件-数据库
	upDbconf := gf.Map{
		"hostname": parameter["hostname"],
		"hostport": parameter["hostport"],
		"username": parameter["username"],
		"password": parameter["password"],
		"dbname":   parameter["dbname"],
		"prefix":   table_prefix,
	}
	cferr := gf.UpConfFieldData(path, upDbconf, "    ")
	if cferr != nil {
		gf.Failed().SetMsg("修改数据库配置失败").SetData(err.Error()).Regin(c)
		return
	}
	//修改应用配置
	upAppconf := gf.Map{
		"vueobjroot":   parameter["vuepath"],
		"busDirName":   "business",
		"adminDirName": "admin",
	}
	cferr = gf.UpConfFieldData(path, upAppconf, "  ")
	if cferr != nil {
		gf.Failed().SetMsg("修改前端路径失败").SetData(err.Error()).Regin(c)
		return
	}
	// 热更数据库配置
	gform.SetConfig(gform.Config{
		"default": gform.ConfigGroup{
			gform.ConfigNode{
				Hostname:  gconv.String(parameter["hostname"]),
				Hostport:  gconv.String(parameter["hostport"]),
				Username:  gconv.String(parameter["username"]),
				Password:  gconv.String(parameter["password"]),
				Dbname:    gconv.String(parameter["dbname"]),
				Prefix:    table_prefix,
				Type:      "mysql",
				Role:      "master",
				Weight:    100,
				CreatedAt: "createtime",
				UpdatedAt: "updatetime",
				DeletedAt: "deletetime",
			},
		},
	})
	gf.Success().SetMsg("安装成功，去安装前端并运行试试！如果数据库报错可重启服务再试").SetData(parameter).Regin(c)
}

// 移除admin控制器 判断存在则移除
func ChecAdminRemoveController() {
	filePath := filepath.Join("app/controller.go")
	con_path := "gofly/app/admin"
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	var result = ""
	for {
		a, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		if strings.Contains(string(a), con_path) || strings.Contains(string(a), fmt.Sprintf("admin.RouterHandler(c, \"%v\")", "admin")) { //存在路由则移除
			continue
		} else {
			result += string(a) + "\n"
		}
	}
	fw, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //os.O_TRUNC清空文件重新写入，否则原文件内容可能残留
	w := bufio.NewWriter(fw)
	w.WriteString(result)
	if err != nil {
		panic(err)
	}
	w.Flush()
}

// DeCompress 解压文件 返回解压的目录
// zipFile 完整文件路径，dest文件目录
func DeCompress(zipFile, dest string) (string, error) {
	// 打开zip文件
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return "", err
	}
	defer func() {
		err := reader.Close()
		if err != nil {
			// global.App.Log.Info(fmt.Sprintf("解压文件关闭失败: %v\n", err.Error()))
		}
	}()
	var (
		first string // 记录第一次的解压的名字
		order int    = 0
	)
	for _, file := range reader.File {
		rc, err := file.Open()
		if err != nil {
			return "", err
		}
		filename := filepath.Join(dest, file.Name)
		//记录第一次的名字
		if order == 0 {
			first = filename
		}
		order += 1
		if file.FileInfo().IsDir() {
			err = os.MkdirAll(filename, 0755)
			if err != nil {
				return "", err
			}
		} else {
			w, err := os.Create(filename)
			if err != nil {
				return "", err
			}
			//defer w.Close()
			_, err = io.Copy(w, rc)
			if err != nil {
				return "", err
			}
			iErr := w.Close()
			if iErr != nil {
				// global.App.Log.Info(fmt.Sprintf("[unzip]: close io %s\n", iErr.Error()))
			}
			fErr := rc.Close()
			if fErr != nil {
				// global.App.Log.Info(fmt.Sprintf("[unzip]: close io %s\n", fErr.Error()))
			}
		}
	}
	return first, nil
}

// Unzip decompresses a zip file to specified directory.
// Note that the destination directory don't need to specify the trailing path separator.
// If the destination directory doesn't exist, it will be created automatically.
func Unzip(zipath, dir string) error {
	// Open zip file.
	reader, err := zip.OpenReader(zipath)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, file := range reader.File {
		if err := unzipFile(file, dir); err != nil {
			return err
		}
	}
	return nil
}

func unzipFile(file *zip.File, dir string) error {
	// Prevent path traversal vulnerability.
	// Such as if the file name is "../../../path/to/file.txt" which will be cleaned to "path/to/file.txt".
	name := strings.TrimPrefix(filepath.Join(string(filepath.Separator), file.Name), string(filepath.Separator))
	filePath := path.Join(dir, name)

	// Create the directory of file.
	if file.FileInfo().IsDir() {
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
			return err
		}
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	// Open the file.
	r, err := file.Open()
	if err != nil {
		return err
	}
	defer r.Close()

	// Create the file.
	w, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer w.Close()

	// Save the decompressed file content.
	_, err = io.Copy(w, r)
	return err
}

// 2复制整个文件夹下文件到另个文件夹 targetPath文件夹，destPath复制的文件
func CopyDir(targetPath string, destPath string) error {
	err := filepath.Walk(targetPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		destPath := filepath.Join(destPath, path[len(targetPath):])
		//如果是个文件夹则创建这个文件夹
		if info.IsDir() {
			return os.MkdirAll(destPath, info.Mode())
		}
		//如果是文件则生成这个文件
		return copyFile(path, destPath)

	})
	return err
}

// 复制单个文件
func copyFile(srcFile, destFile string) error {
	src, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer src.Close()
	//创建复制的文件
	dest, err := os.Create(destFile)
	if err != nil {
		return err
	}
	defer dest.Close()
	//复制内容到文件
	_, err = io.Copy(dest, src)
	if err != nil {
		return err
	}
	//让复制的文件将内容存到硬盘而不是缓存
	err = dest.Sync()
	if err != nil {
		return err
	}
	return nil
}
