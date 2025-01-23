package developer

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gofly/utils/gf"
	"gofly/utils/tools/gcfg"
	"gofly/utils/tools/gconv"
	"gofly/utils/tools/gfile"
	"gofly/utils/tools/gjson"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// 代码仓安装操作
type Codestoreoption struct{ NoNeedAuths []string }

func init() {
	fpath := Codestoreoption{NoNeedAuths: []string{"*"}}
	gf.Register(&fpath, fpath)
}

// 安装前下载代码
func (api *Codestoreoption) DownCode(c *gf.GinCtx) {
	if appConf_arr["runEnv"] == "release" {
		gf.Failed().SetMsg("生产环境禁止操作，请在开发环境下操作").Regin(c)
		return
	}
	parameter, perr := gf.PostParam(c)
	packName := gconv.String(parameter["name"])
	if packName == "" || perr != nil {
		gf.Failed().SetMsg("安装包名称不能为空").SetData(perr).Regin(c)
		return
	}
	//判断如果本地存在代码包就不用下载
	path, _ := os.Getwd()
	install_apppath := filepath.Join(path, "/devsource/codemarket/install", packName)
	if _, err := os.Stat(install_apppath); err == nil {
		gf.Success().SetMsg("本地代码已存在直接安装").SetData(true).Regin(c)
		return
	}
	baseurl := c.DefaultQuery("baseurl", "")
	if baseurl == "" {
		gf.Failed().SetMsg("私有仓不存在！").Regin(c)
	} else {
		ref, resErro := gf.HttpGet(baseurl+"/goflycode/content/getDownUrl", parameter)
		if resErro != nil {
			gf.Failed().SetMsg("请求下载地址失败").SetData(resErro).Regin(c)
		} else {
			if gconv.Int(ref["code"]) == 0 && ref["data"] != nil {
				downdir := filepath.Join(path, "/devsource/codemarket/install", gconv.String(parameter["name"])+".zip")
				downstatus, downdir_zippath := DownFileToDir(gconv.String(ref["data"]), downdir)
				if downstatus {
					dezipdir := filepath.Join(path, "/devsource/codemarket/install")
					err := Unzip(downdir_zippath, dezipdir)
					if err == nil {
						os.Remove(downdir_zippath)
						gf.Success().SetMsg("下载代码成功").SetData(true).Regin(c)
					} else {
						gf.Failed().SetMsg("下载代码失败").SetData(resErro).Regin(c)
					}
				} else {
					gf.Failed().SetMsg("下载代码失败").SetData(resErro).Regin(c)
				}
			} else {
				gf.Failed().SetMsg("下载代码失败").SetData(resErro).Regin(c)
			}
		}
	}
}

// 安装代码
func (api *Codestoreoption) InstallCode(c *gf.GinCtx) {
	prefix := dbConf_arr["prefix"]
	if prefix == nil {
		prefix = ""
	}
	userID := c.GetInt64("userID") //当前用户ID
	//更新配置
	cf, _ := gcfg.New()
	data := cf.MustGet(ctx, "app")
	appConf_arr = gconv.Map(data)
	parameter, perr := gf.PostParam(c)
	packName := gconv.String(parameter["name"])
	if appConf_arr["runEnv"] == "release" {
		gf.Failed().SetMsg("生产环境禁止操作，请在开发环境下操作").Regin(c)
		return
	}
	if packName == "" || perr != nil {
		gf.Failed().SetMsg("安装包名称不能为空").Regin(c)
		return
	}
	path, err := os.Getwd()
	if err != nil {
		gf.Failed().SetMsg("项目路径获取失败").Regin(c)
		return
	}
	//1.安装后端go
	CopyAllDir(filepath.Join(path, "/devsource/codemarket/install", packName, "/go/app"), filepath.Join(path, "/app"))
	//utils扩展
	CopyAllDir(filepath.Join(path, "/devsource/codemarket/install", packName, "/go/plugin"), filepath.Join(path, "/utils/plugin"))
	//2.导入数据表
	SqlPath := filepath.Join(path, "/devsource/codemarket/install", packName, "/install.sql")
	aqlerr := ImportSqlFile(SqlPath)
	if aqlerr != nil {
		gf.Failed().SetMsg("导入插件sql数据文件失败").SetData(aqlerr).Regin(c)
		return
	}
	//代码包配置
	install_cofig, _ := GetInstallConfig(filepath.Join(path, "/devsource/codemarket/install", packName))
	//如果存在前缀-添加导入表前缀
	if prefix != "" && install_cofig.Sqldb.Packtables != "" {
		var tableNames_arr = strings.Split(install_cofig.Sqldb.Packtables, ",")
		for _, tableName := range tableNames_arr {
			if tableName != "" {
				gf.DB().Exec(c, "ALTER TABLE "+tableName+" RENAME TO "+gf.String(prefix)+tableName+";")
			}
		}
	}
	//3.导入菜单
	businessmenuPath := filepath.Join(path, "/devsource/codemarket/install", packName, "/businessmenu.json")
	bmenufile, _ := os.Open(businessmenuPath)
	bmenubytes, bmenuerr := io.ReadAll(bmenufile)
	bnenuids := ""
	if bmenuerr == nil && bmenubytes != nil {
		var data interface{}
		json.Unmarshal(bmenubytes, &data)
		if install_cofig.Sqldb.Businessmenuids != "" { //删除菜单
			gf.DB().Exec(c, "DELETE FROM "+gf.String(prefix)+"business_auth_rule WHERE id IN ("+install_cofig.Sqldb.Businessmenuids+")")
		}
		m_nenuids := Insertmenu(userID, data, 0, "business_auth_rule")
		bnenuids = strings.Join(m_nenuids, ",")
	}
	//关闭menu.json读取
	bmenufile.Close()
	//admin端
	adminmenuPath := filepath.Join(path, "/devsource/codemarket/install", packName, "/adminmenu.json")
	amenufile, _ := os.Open(adminmenuPath)
	amenubytes, amenuerr := io.ReadAll(amenufile)
	anenuids := ""
	if amenuerr == nil && amenubytes != nil {
		var data interface{}
		json.Unmarshal([]byte(amenubytes), &data)
		if install_cofig.Sqldb.Adminmenuids != "" { //删除菜单
			gf.DB().Exec(c, "DELETE FROM "+gf.String(prefix)+"admin_auth_rule WHERE id IN ("+install_cofig.Sqldb.Adminmenuids+")")
		}
		m_nenuids := Insertmenu(userID, data, 0, "admin_auth_rule")
		anenuids = strings.Join(m_nenuids, ",")
	}
	//关闭menu.json读取
	amenufile.Close()
	//安装前端代码
	vueobjrootPath := filepath.Join(gconv.String(appConf_arr["vueobjroot"]))
	vueobjrootPathadmin := filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(appConf_arr["adminDirName"]))
	//获取文件路径
	vue_views, _ := GetAllFile(filepath.Join(path, "/devsource/codemarket/install", packName, "/vue/business/views"))
	vue_components, _ := GetAllFile(filepath.Join(path, "/devsource/codemarket/install", packName, "/vue/business/components"))
	if _, err := os.Stat(vueobjrootPath); !os.IsNotExist(err) {
		if install_cofig.App.Installcover { //覆盖安装删除原来代码
			b_views := strings.Split(gf.String(vue_views["folders"]), ",")
			for _, fpath := range b_views {
				os.RemoveAll(filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(appConf_arr["busDirName"]), "/src/views", fpath))
			}
		}
		//处理business
		CopyAllDir(filepath.Join(path, "/devsource/codemarket/install/", packName, "/vue/business/views"), filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(appConf_arr["busDirName"]), "/src/views"))
		CopyAllDir(filepath.Join(path, "/devsource/codemarket/install/", packName, "/vue/business/components"), filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(appConf_arr["busDirName"]), "/src/components"))
		//判断是否存在admin
		if _, err := os.Stat(vueobjrootPathadmin); !os.IsNotExist(err) {
			CopyAllDir(filepath.Join(path, "/devsource/codemarket/install/", packName, "/vue/admin/views"), filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(appConf_arr["adminDirName"]), "/src/views"))
			CopyAllDir(filepath.Join(path, "/devsource/codemarket/install/", packName, "/vue/admin/components"), filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(appConf_arr["adminDirName"]), "/src/components"))
		}
	}
	//5.更新配置文件
	go_app_str, go_app_arr, _ := GetAllFileApp(filepath.Join(path, "/devsource/codemarket/install", packName, "/go/app"))
	utilsfiles, _ := GetAllFile(filepath.Join(path, "/devsource/codemarket/install", packName, "/go/plugin"))
	upconf := gf.Map{"viewsfilesbusiness": vue_views["files"], "viewsfoldersbusiness": vue_views["folders"], "componentfilesbusiness": vue_components["files"], "componentfoldersbusiness": vue_components["folders"],
		"appfolders": go_app_str, "utilsfiles": utilsfiles["files"], "isinstall": true, "businessmenuids": bnenuids, "adminmenuids": anenuids}
	//如果存在admin
	if _, err := os.Stat(vueobjrootPathadmin); !os.IsNotExist(err) {
		vue_views, _ := GetAllFile(filepath.Join(path, "/devsource/codemarket/install", packName, "/vue/admin/views"))
		vue_components, _ := GetAllFile(filepath.Join(path, "/devsource/codemarket/install", packName, "/vue/admin/components"))
		upconf["viewsfilesadmin"] = vue_views["files"]
		upconf["viewsfoldersadmin"] = vue_views["folders"]
		upconf["componentfilesadmin"] = vue_components["files"]
		upconf["componentfoldersadmin"] = vue_components["folders"]
	}
	cferr := UpConfFieldData(path+"/devsource/codemarket/install/"+packName, upconf)
	if cferr != nil {
		gf.Success().SetMsg("更新配置失败，请重新安装").SetData(packName).Regin(c)
		return
	} else {
		//判断模块控制器是否存在
		for _, go_app_dir := range go_app_arr {
			modelname := ""
			//添加根模块-在app下的控制器
			if strings.Contains(go_app_dir, "/") {
				go_path_arr := strings.Split(go_app_dir, "/")
				modelname = go_path_arr[0]
				//判断安装模块下是否存在控制器controller.go文件，haseMoleCtr=true是存在
				var haseMoleCtr bool = false
				if _, err := os.Stat(filepath.Join(path, "/devsource/codemarket/install", packName, "/go/app/", modelname, "controller.go")); !os.IsNotExist(err) {
					haseMoleCtr = true
				}
				CheckIsAddController("", modelname, haseMoleCtr)
			}
			//过滤lib-添加类型模块控制-例如business下面的控制
			if !strings.HasSuffix(go_app_dir, "lib") {
				CheckIsAddController(modelname, go_app_dir, false)
			}
		}
		if install_cofig.Installgo.NoVerifyAPIRoot != "" {
			//同时判断模块是否添加token和合法性验证
			AddOrRemoveValidity(install_cofig.Installgo.NoVerifyAPIRoot, "install", path, appConf_arr)
		}
		//分别复制配置文件到resource的的配置文件夹
		if !install_cofig.App.Installcover {
			gfile.CopyFile(filepath.Join(path, "/devsource/codemarket/install/", packName, "config.yml"), filepath.Join(path, "/resource/codeinstall", packName, "config.yml"))
		}
		//复制包配置到/resource/config下
		confFilePath := filepath.Join(path, "/devsource/codemarket/install/", packName, packName+".yaml")
		if _, err := os.Stat(confFilePath); !os.IsNotExist(err) { //存在
			gfile.CopyFile(confFilePath, filepath.Join(path, "/resource/config", packName+".yaml"))
		}
		// 如存在附件资源存在则复制到 /resource/static下
		staticFilePath := filepath.Join(path, "/devsource/codemarket/install/", packName, packName)
		if _, err := os.Stat(staticFilePath); !os.IsNotExist(err) { //存在
			CopyAllDir(staticFilePath, filepath.Join(path, "/resource/static", packName))
		}
		//删除安装文件包
		os.RemoveAll(filepath.Join(path, "/devsource/codemarket/install/", packName))
		gf.Success().SetMsg("安装成功").SetData(packName).Regin(c)
	}
}

// 卸载代码
func (api *Codestoreoption) UninstallCode(c *gf.GinCtx) {
	if appConf_arr["runEnv"] == "release" {
		gf.Failed().SetMsg("生产环境禁止操作，请在开发环境下操作").Regin(c)
		return
	}
	//前缀
	prefix := dbConf_arr["prefix"]
	if prefix == nil {
		prefix = ""
	}
	//更新配置
	cf, _ := gcfg.New()
	data := cf.MustGet(ctx, "app")
	appConf_arr = gconv.Map(data)

	parameter, perr := gf.PostParam(c)
	packName := gconv.String(parameter["name"])
	if packName == "" || perr != nil {
		gf.Failed().SetMsg("安装包名称不能为空").Regin(c)
		return
	}
	path, err := os.Getwd()
	if err != nil {
		gf.Failed().SetMsg("项目路径获取失败").Regin(c)
		return
	}
	//获取配置文件
	apppath := filepath.Join(path, "/resource/codeinstall", packName)
	install_cofig, _ := GetInstallConfig(apppath)
	//1.卸载数据库
	if install_cofig.Sqldb.Packtables != "" {
		var tableNames_arr = strings.Split(install_cofig.Sqldb.Packtables, ",")
		for _, tableName := range tableNames_arr {
			if tableName != "" {
				gf.DB().Exec(c, "DROP TABLE IF EXISTS "+gf.String(prefix)+tableName)
			}
		}
	}
	//卸载菜单
	if install_cofig.Sqldb.Businessmenuids != "" {
		gf.DB().Exec(c, "DELETE FROM "+gf.String(prefix)+"business_auth_rule WHERE id IN ("+install_cofig.Sqldb.Businessmenuids+")")
		gf.Model("business_auth_rule").WhereIn("pid", strings.Split(install_cofig.Sqldb.Businessmenuids, ",")).Delete()
	}
	if install_cofig.Sqldb.Adminmenuids != "" {
		gf.DB().Exec(c, "DELETE FROM "+gf.String(prefix)+"admin_auth_rule WHERE id IN ("+install_cofig.Sqldb.Adminmenuids+")")
		gf.Model("admin_auth_rule").WhereIn("pid", strings.Split(install_cofig.Sqldb.Adminmenuids, ",")).Delete()
	}
	//2.卸载前端
	//2.1 Business端
	vue_viewsfiles_path := filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(appConf_arr["busDirName"]))
	if _, err := os.Stat(vue_viewsfiles_path); err == nil {
		if install_cofig.Installvue.Viewsfilesbusiness != "" {
			var files_arr = strings.Split(install_cofig.Installvue.Viewsfilesbusiness, ",")
			for _, filename := range files_arr {
				os.Remove(filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(appConf_arr["busDirName"]), "/src/views", filename))
			}
		}
		if install_cofig.Installvue.Viewsfoldersbusiness != "" {
			var folders_arr = strings.Split(install_cofig.Installvue.Viewsfoldersbusiness, ",")
			for _, foldername := range folders_arr {
				if foldername != "" {
					os.RemoveAll(filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(appConf_arr["busDirName"]), "/src/views", foldername))
				}
			}
		}
		if install_cofig.Installvue.Componentfilesbusiness != "" {
			var Cfiles_arr = strings.Split(install_cofig.Installvue.Componentfilesbusiness, ",")
			for _, cfilesname := range Cfiles_arr {
				os.Remove(filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(appConf_arr["busDirName"]), "/src/components", cfilesname))
			}
		}
		if install_cofig.Installvue.Componentfoldersbusiness != "" {
			var Cfolders_arr = strings.Split(install_cofig.Installvue.Componentfoldersbusiness, ",")
			for _, cfoldername := range Cfolders_arr {
				if cfoldername != "" {
					os.RemoveAll(filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(appConf_arr["busDirName"]), "/src/components", cfoldername))
				}
			}
		}
	}
	//2.2 admin端
	vue_viewsfiles_path_admin := filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(appConf_arr["adminDirName"]))
	if _, err := os.Stat(vue_viewsfiles_path_admin); err == nil {
		if install_cofig.Installvue.Viewsfilesadmin != "" {
			var files_arr = strings.Split(install_cofig.Installvue.Viewsfilesadmin, ",")
			for _, filename := range files_arr {
				os.Remove(filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(appConf_arr["adminDirName"]), "/src/views", filename))
			}
		}
		if install_cofig.Installvue.Viewsfoldersadmin != "" {
			var folders_arr = strings.Split(install_cofig.Installvue.Viewsfoldersadmin, ",")
			for _, foldername := range folders_arr {
				if foldername != "" {
					os.RemoveAll(filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(appConf_arr["adminDirName"]), "/src/views", foldername))
				}
			}
		}
		if install_cofig.Installvue.Componentfilesadmin != "" {
			var Cfiles_arr = strings.Split(install_cofig.Installvue.Componentfilesadmin, ",")
			for _, cfilesname := range Cfiles_arr {
				os.Remove(filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(appConf_arr["adminDirName"]), "/src/components", cfilesname))
			}
		}
		if install_cofig.Installvue.Componentfoldersadmin != "" {
			var Cfolders_arr = strings.Split(install_cofig.Installvue.Componentfoldersadmin, ",")
			for _, cfoldername := range Cfolders_arr {
				if cfoldername != "" {
					os.RemoveAll(filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(appConf_arr["adminDirName"]), "/src/components", cfoldername))
				}
			}
		}
	}
	//3.删除后端代码
	if install_cofig.Installgo.Appfolders != "" {
		var app_folders_arr = strings.Split(install_cofig.Installgo.Appfolders, ",")
		for _, appFolderName := range app_folders_arr {
			if appFolderName != "" {
				os.RemoveAll(filepath.Join(path, "app", appFolderName))
				modelname := ""
				if strings.Contains(appFolderName, "/") {
					go_path_arr := strings.Split(appFolderName, "/")
					modelname = go_path_arr[0]
				}
				CheckApiRemoveController(modelname, appFolderName)
				//判断是否要删除model目录
				mode_name := strings.Split(appFolderName, "/")
				model_dir := filepath.Join(path, "app", mode_name[0])
				hasefolders, _ := GetDirHasefolder(model_dir)
				if !hasefolders {
					os.RemoveAll(model_dir)
					//删除APP下的控制模块
					CheckApiRemoveController("", mode_name[0])
					// CheckApiRemoveAppController(mode_name[0])
				}
			}
		}
	}
	//卸载后端plugin组件
	if install_cofig.Installgo.Utilsfiles != "" {
		var utils_files_arr = strings.Split(install_cofig.Installgo.Utilsfiles, ",")
		for _, utilsFileName := range utils_files_arr {
			os.RemoveAll(filepath.Join(path, "utils/plugin", utilsFileName))
			//检查是否存在和文件名相同的文件夹
			filenames := strings.Split(utilsFileName, ".")
			if len(filenames) == 2 {
				plugin_folder := filepath.Join(path, "/utils/plugin", filenames[0])
				if _, err := os.Stat(plugin_folder); err == nil { //存在对应目录
					os.RemoveAll(plugin_folder)
				}
			}
		}
	}
	if install_cofig.Installgo.NoVerifyAPIRoot != "" {
		//同时删除合法性验证和token验证
		AddOrRemoveValidity(install_cofig.Installgo.NoVerifyAPIRoot, "uninstall", path, appConf_arr)
	}
	//4.删除配置文件
	os.RemoveAll(filepath.Join(path, "/resource/codeinstall", packName)) //删除配置文件
	//5.删除config配置文件
	ResourceConfig := filepath.Join(path, "/resource/config", packName+".yaml")
	if _, err := os.Stat(ResourceConfig); err == nil {
		os.Remove(ResourceConfig)
	}
	// 5.如果/resource/static下存在资源文件则删除
	staticFilePath := filepath.Join(path, "/resource/static", packName)
	if _, err := os.Stat(staticFilePath); !os.IsNotExist(err) { //存在
		os.RemoveAll(staticFilePath)
	}
	gf.Success().SetMsg("卸载成功").Regin(c)
}

// 打包
func (api *Codestoreoption) PackCode(c *gf.GinCtx) {
	if appConf_arr["runEnv"] == "release" {
		gf.Failed().SetMsg("生产环境禁止操作，请在开发环境下操作").Regin(c)
		return
	}
	//更新配置
	cf, _ := gcfg.New()
	data := cf.MustGet(ctx, "app")
	appConf_arr = gconv.Map(data)
	params, _ := gf.RequestParam(c)
	packName := gconv.String(params["name"])
	if packName == "" {
		gf.Failed().SetMsg("打包文件不存在").SetData(params).Regin(c)
		return
	}
	path, err := os.Getwd()
	if err != nil {
		gf.Failed().SetMsg("项目路径获取失败").Regin(c)
		return
	}
	pack_path := filepath.Join(path, "/devsource/codemarket/release/", packName)
	//1制作打包文件
	if _, err := os.Stat(pack_path); err != nil && !os.IsExist(err) {
		os.MkdirAll(pack_path, os.ModePerm)
	} else {
		os.RemoveAll(pack_path)
		os.MkdirAll(pack_path, os.ModePerm)
	}
	//2复制包模板
	CopyAllDir(filepath.Join(path, "/devsource/developer/codetpl/packcode"), pack_path)
	//更新基础配置
	upconf := map[string]interface{}{"version": params["version"], "title": params["title"], "installcover": params["installcover"], "des": params["des"], "name": params["name"], "packtables": params["packtables"], "noVerifyAPIRoot": params["noVerifyAPIRoot"]}
	UpConfFieldData(path+"/devsource/codemarket/release/"+packName, upconf)
	//3复制前端文件
	//3.1复制business
	if fmt.Sprintf("%T", params["viewsfiles_business"]) == "[]interface {}" {
		var viewsfile_arr = params["viewsfiles_business"].([]interface{})
		for _, viewsfile := range viewsfile_arr {
			if viewsfile != "" {
				viewsfile_dirname := strings.Trim(gf.JSONToString(viewsfile), "\"")
				CopyFile(filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(appConf_arr["busDirName"]), "/src/views", viewsfile_dirname), filepath.Join(path, "/devsource/codemarket/release/", packName, "/vue/business/views", viewsfile_dirname))
			}
		}
	}
	if fmt.Sprintf("%T", params["viewsfolders_business"]) == "[]interface {}" {
		var viewsfolders_arr = params["viewsfolders_business"].([]interface{})
		for _, viewsfolder := range viewsfolders_arr {
			if viewsfolder != "" {
				views_dirname := strings.Trim(gf.JSONToString(viewsfolder), "\"")
				CopyAllDir(filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(appConf_arr["busDirName"]), "/src/views", views_dirname), filepath.Join(path, "/devsource/codemarket/release/", packName, "/vue/business/views", views_dirname))
			}
		}
	}
	if fmt.Sprintf("%T", params["componentfiles_business"]) == "[]interface {}" {
		var componentfile_arr = params["componentfiles_business"].([]interface{})
		for _, componentfile := range componentfile_arr {
			if componentfile != "" {
				componentfile_dirname := strings.Trim(gf.JSONToString(componentfile), "\"")
				CopyFile(filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(appConf_arr["busDirName"]), "/src/components", componentfile_dirname), filepath.Join(path, "/devsource/codemarket/release/", packName, "/vue/business/components", componentfile_dirname))
			}
		}
	}
	if fmt.Sprintf("%T", params["componentfolders_business"]) == "[]interface {}" {
		var componentfolder_arr = params["componentfolders_business"].([]interface{})
		for _, componentfolder := range componentfolder_arr {
			if componentfolder != "" {
				componentfolder_dirname := strings.Trim(gf.JSONToString(componentfolder), "\"")
				CopyAllDir(filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(appConf_arr["busDirName"]), "/src/components", componentfolder_dirname), filepath.Join(path, "/devsource/codemarket/release/", packName, "/vue/business/components", componentfolder_dirname))
			}
		}
	}
	//3.2复制admin
	if fmt.Sprintf("%T", params["viewsfiles_admin"]) == "[]interface {}" {
		var viewsfile_arr = params["viewsfiles_admin"].([]interface{})
		for _, viewsfile := range viewsfile_arr {
			if viewsfile != "" {
				viewsfile_dirname := strings.Trim(gf.JSONToString(viewsfile), "\"")
				CopyFile(filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(appConf_arr["adminDirName"]), "/src/views", viewsfile_dirname), filepath.Join(path, "/devsource/codemarket/release/", packName, "/vue/admin/views", viewsfile_dirname))
			}
		}
	}
	if fmt.Sprintf("%T", params["viewsfolders_admin"]) == "[]interface {}" {
		var viewsfolders_arr = params["viewsfolders_admin"].([]interface{})
		for _, viewsfolder := range viewsfolders_arr {
			if viewsfolder != "" {
				views_dirname := strings.Trim(gf.JSONToString(viewsfolder), "\"")
				CopyAllDir(filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(appConf_arr["adminDirName"]), "/src/views", views_dirname), filepath.Join(path, "/devsource/codemarket/release/", packName, "/vue/admin/views", views_dirname))
			}
		}
	}
	if fmt.Sprintf("%T", params["componentfiles_admin"]) == "[]interface {}" {
		var componentfile_arr = params["componentfiles_admin"].([]interface{})
		for _, componentfile := range componentfile_arr {
			if componentfile != "" {
				componentfile_dirname := strings.Trim(gf.JSONToString(componentfile), "\"")
				CopyFile(filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(appConf_arr["adminDirName"]), "/src/components", componentfile_dirname), filepath.Join(path, "/devsource/codemarket/release/", packName, "/vue/admin/components", componentfile_dirname))
			}
		}
	}
	if fmt.Sprintf("%T", params["componentfolders_admin"]) == "[]interface {}" {
		var componentfolder_arr = params["componentfolders_admin"].([]interface{})
		for _, componentfolder := range componentfolder_arr {
			if componentfolder != "" {
				componentfolder_dirname := strings.Trim(gf.JSONToString(componentfolder), "\"")
				CopyAllDir(filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(appConf_arr["adminDirName"]), "/src/components", componentfolder_dirname), filepath.Join(path, "/devsource/codemarket/release/", packName, "/vue/admin/components", componentfolder_dirname))
			}
		}
	}
	//后端代码
	if fmt.Sprintf("%T", params["appfolders"]) == "[]interface {}" {
		var appfolder_arr = params["appfolders"].([]interface{})
		for _, appfolder := range appfolder_arr {
			if appfolder != "" {
				appfolder_dirname := strings.Trim(gf.JSONToString(appfolder), "\"")
				CopyAllDir(filepath.Join(path, "/app", appfolder_dirname), filepath.Join(path, "/devsource/codemarket/release/", packName, "/go/app", appfolder_dirname))
			}
		}
	}
	//打包plugin插件
	if fmt.Sprintf("%T", params["utilsfiles"]) == "[]interface {}" {
		var utilsfile_arr = params["utilsfiles"].([]interface{})
		for _, utilsfile := range utilsfile_arr {
			if utilsfile != "" {
				utilsfile_dirname := strings.Trim(gf.JSONToString(utilsfile), "\"")
				CopyAllDir(filepath.Join(path, "/utils/plugin", utilsfile_dirname), filepath.Join(path, "/devsource/codemarket/release/", packName, "/go/plugin", utilsfile_dirname))
				//检查是否存在和文件名相同的文件夹
				filenames := strings.Split(utilsfile_dirname, ".")
				if len(filenames) == 2 {
					plugin_folder := filepath.Join(path, "/utils/plugin", filenames[0])
					if _, err := os.Stat(plugin_folder); err == nil { //存在对应目录
						CopyAllDir(plugin_folder, filepath.Join(path, "/devsource/codemarket/release/", packName, "/go/plugin", filenames[0]))
					}

				}
			}
		}
	}
	//导出数据库表
	if params["packtables"] != "" {
		pathname := filepath.Join(path, "/devsource/codemarket/release", packName, "install.sql")
		tables := strings.Split(gconv.String(params["packtables"]), ",")
		ExecSqlFile(tables, pathname)
	}
	//把菜单数据写入menu.json文件
	if params["businessmenujson"] != nil {
		menujson := params["businessmenujson"]
		meni_json := filepath.Join(path, "/devsource/codemarket/release/", packName, "businessmenu.json")
		if _, err := os.Stat(meni_json); err != nil {
			if !os.IsExist(err) {
				os.MkdirAll(meni_json, os.ModePerm)
			}
		}
		menudata, _ := gjson.Marshal(menujson)
		os.WriteFile(meni_json, menudata, 0777)
	}
	if params["adminmenujson"] != nil {
		menujson := params["adminmenujson"]
		meni_json := filepath.Join(path, "/devsource/codemarket/release/", packName, "adminmenu.json")
		if _, err := os.Stat(meni_json); err != nil {
			if !os.IsExist(err) {
				os.MkdirAll(meni_json, os.ModePerm)
			}
		}
		menudata, _ := gjson.Marshal(menujson)
		os.WriteFile(meni_json, menudata, 0777)
	}
	// 查看/resource/config是否存在动态配置文件-存在则复制到包目录下
	confFilePath := filepath.Join(path, "/resource/config", packName+".yaml")
	if _, err := os.Stat(confFilePath); !os.IsNotExist(err) { //存在
		gfile.CopyFile(confFilePath, filepath.Join(pack_path, packName+".yaml"))
	}
	// 查看/resource/static是否存在静态文件-存在则复制到包目录下
	staticFilePath := filepath.Join(path, "/resource/static", packName)
	if _, err := os.Stat(staticFilePath); !os.IsNotExist(err) { //存在
		CopyAllDir(staticFilePath, filepath.Join(pack_path, packName))
	}
	//打包文件路径
	if _, err := os.Stat(pack_path); err == nil {
		defer os.RemoveAll(pack_path) //最后删除文件夹
		//压缩成zip
		packfile, err := os.Open(pack_path)
		if err != nil {
			gf.Failed().SetMsg("打包错误1").SetData(err).Regin(c)
			return
		}
		defer packfile.Close()
		dest := filepath.Join(path, "/devsource/codemarket/release", packName+".zip")
		err = gf.Compress(packfile, dest)
		if err != nil {
			gf.Failed().SetMsg("打包错误2").SetData(err).Regin(c)
			return
		}
		gf.Success().SetMsg("打包成功").SetData(packName).Regin(c)
	} else {
		gf.Failed().SetMsg("文件不存在").SetData(params).Regin(c)
	}
}

// 查找本地已经安装的包
func (api *Codestoreoption) GetInstallPack(c *gf.GinCtx) {
	path, _ := os.Getwd()
	pathname := filepath.Join(path, "/resource/codeinstall")
	rd, err := os.ReadDir(pathname)
	if err != nil {
		gf.Success().SetMsg("本地没有安装的包").SetData("").Regin(c)
		return
	}
	var folders = make([]string, 0)
	for _, fi := range rd {
		if fi.IsDir() {
			apppath := filepath.Join(path, "/resource/codeinstall", fi.Name())
			install_cofig, _ := GetInstallConfig(apppath)
			if install_cofig.App.Isinstall {
				folders = append(folders, fi.Name())
			}
		}
	}
	gf.Success().SetMsg("本地已经安装的包").SetData(strings.Join(folders, ",")).Regin(c)
}

// 获取打包文件目录
func (api *Codestoreoption) GetPackdirs(c *gf.GinCtx) {
	path, _ := os.Getwd()
	_, go_app_arr, _ := GetAppFileNoAdmin(filepath.Join(path, "/app"))
	go_app_dir, _ := gf.GetAllFileArray(filepath.Join(path, "/app"))
	utilstool, _ := gf.GetAllFileArray(filepath.Join(path, "/utils/plugin"))
	//前端
	vue_components_business, _ := gf.GetAllFileArray(filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(appConf_arr["busDirName"]), "/src/components"))
	vue_views_business, _ := gf.GetAllFileArray(filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(appConf_arr["busDirName"]), "/src/views"))
	rebdata := gf.Map{"applist": go_app_arr, "utilstool": utilstool["files"], "appDirlist": go_app_dir["folders"],
		"vue_business": gf.Map{
			"componentfiles": vue_components_business["files"], "componentfolders": vue_components_business["folders"],
			"viewsfiles": vue_views_business["files"], "viewsfolders": vue_views_business["folders"],
		},
		"vue_admin": gf.Map{
			"componentfiles": make([]interface{}, 0), "componentfolders": make([]interface{}, 0),
			"viewsfiles": make([]interface{}, 0), "viewsfolders": make([]interface{}, 0),
		},
	}
	// admin端
	vue_viewsfiles_path_admin := filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(appConf_arr["adminDirName"]))
	if _, err := os.Stat(vue_viewsfiles_path_admin); err == nil {
		vue_components_admin, _ := gf.GetAllFileArray(filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(appConf_arr["adminDirName"]), "/src/components"))
		vue_views_admin, _ := gf.GetAllFileArray(filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(appConf_arr["adminDirName"]), "/src/views"))
		madmin := gf.Map{
			"componentfiles": vue_components_admin["files"], "componentfolders": vue_components_admin["folders"],
			"viewsfiles": vue_views_admin["files"], "viewsfolders": vue_views_admin["folders"],
		}
		rebdata["vue_admin"] = madmin
	}
	gf.Success().SetMsg("获取打包文件目录").SetData(rebdata).Regin(c)
}

// 获取菜单列表
func (api *Codestoreoption) GetMenutreeData(c *gf.GinCtx) {
	business_menuList, _ := gf.Model("business_auth_rule").Fields("id,pid,title,locale").Order("weigh asc").Select()
	if business_menuList == nil {
		business_menuList = make(gf.OrmResult, 0)
	}
	for _, val := range business_menuList {
		if val["title"].String() == "" {
			val["title"] = val["locale"]
		}
	}
	business_menuList = gf.GetTreeArray(business_menuList, 0, "")
	//admin
	var admin_menuList = make(gf.OrmResult, 0)
	vue_viewsfiles_path_admin := filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(appConf_arr["adminDirName"]))
	if _, err := os.Stat(vue_viewsfiles_path_admin); err == nil {
		admin_menuList, _ = gf.Model("admin_auth_rule").Fields("id,pid,title,locale").Order("weigh asc").Select()
		if admin_menuList == nil {
			admin_menuList = make(gf.OrmResult, 0)
		}
		for _, val := range admin_menuList {
			if val["title"].String() == "" {
				val["title"] = val["locale"]
			}
		}
		admin_menuList = gf.GetTreeArray(admin_menuList, 0, "")
	}
	gf.Success().SetMsg("获取打包文件目录").SetData(gf.Map{"business_menuList": business_menuList, "admin_menuList": admin_menuList}).Regin(c)
}

// 获取菜单选择ID转json
func (api *Codestoreoption) MenuTreeToJson(c *gf.GinCtx) {
	params, _ := gf.RequestParam(c)
	tabename := "business_auth_rule"
	if params["formtable"] == "admin" {
		tabename = "admin_auth_rule"
	}
	rules := gf.GetRulesID(tabename, "pid", params["menu"]) //获取子菜单包含的父级ID
	menuList, _ := gf.Model(tabename).WhereIn("id", rules.([]interface{})).
		Fields("id,pid,title,locale,type,icon,routepath,routename,component,permission,path,redirect,isExt,keepalive,hideInMenu,activeMenu,noAffix,onlypage,requiresAuth").Order("weigh asc").Select()
	if menuList == nil {
		menuList = make(gf.OrmResult, 0)
	}
	for _, val := range menuList {
		if val["title"].String() == "" {
			val["title"] = val["locale"]
		}
	}
	menuList = gf.GetRuleTreeArrayByPack(menuList, 0)
	gf.Success().SetMsg("获取菜单选择ID转json").SetData(menuList).Regin(c)
}

// 上传附件
func (api *Codestoreoption) Upfile(c *gf.GinCtx) {
	params := map[string]string{}
	req, err := NewFileUploadRequest(c, params)
	if err != nil {
		gf.Failed().SetMsg(err.Error()).Regin(c)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		gf.Failed().SetMsg(err.Error()).Regin(c)
		return
	}
	body := &bytes.Buffer{}
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	result, _ := io.ReadAll(body)
	ref := string(result)
	var parameter map[string]interface{}
	if err := json.Unmarshal([]byte(ref), &parameter); err == nil {
		if parameter["status"] == "done" {
			c.JSON(200, parameter)
		} else {
			gf.Failed().SetMsg("附件上传失败").SetData(parameter).Regin(c)
		}
	}
}

// NewFileUploadRequest ...
func NewFileUploadRequest(c *gf.GinCtx, params map[string]string) (*http.Request, error) {
	domainurl := c.DefaultPostForm("domainurl", "")
	file, err := c.FormFile("file")
	if err != nil {
		return nil, errors.New("获取数据失败，")
	}
	body := &bytes.Buffer{}
	// 文件写入 body
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", file.Filename)
	if err != nil {
		return nil, err
	}
	filesrc, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer filesrc.Close()
	_, err = io.Copy(part, filesrc)
	// 其他参数列表写入 body
	for k, v := range params {
		if err := writer.WriteField(k, v); err != nil {
			return nil, err
		}
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, domainurl, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())
	verifytime := time.Now().Unix()
	mdsecret := gf.Md5(fmt.Sprintf("gofly@888%v", verifytime))
	req.Header.Add("Verify-Encrypt", mdsecret)
	req.Header.Add("Verify-Time", fmt.Sprintf("%v", verifytime))
	return req, err
}

// 安装本地代码包
func (api *Codestoreoption) InstallLocalCode(c *gf.GinCtx) {
	file, err := c.FormFile("file")
	if err != nil {
		gf.Failed().SetMsg("获取数据失败，").SetData(err).Regin(c)
		return
	}
	pathdir, _ := os.Getwd()
	downpathzip := filepath.Join(pathdir, "/devsource/codemarket/install", file.Filename)
	err = c.SaveUploadedFile(file, downpathzip)
	if err != nil {
		gf.Failed().SetMsg("上传失败").SetData(err).Regin(c)
	} else {
		dezipdir := filepath.Join(pathdir, "/devsource/codemarket/install")
		err := Unzip(downpathzip, dezipdir)
		if err == nil {
			os.Remove(downpathzip)
			filename_arr := strings.Split(file.Filename, ".")
			gf.Success().SetMsg("安装包解压成功").SetData(filename_arr[0]).Regin(c)
		} else {
			gf.Failed().SetMsg("安装包解压失败").SetData(err).Regin(c)
		}
	}

}
