package developer

import (
	"encoding/json"
	"fmt"
	"gofly/utils/gf"
	"gofly/utils/tools/gcfg"
	"gofly/utils/tools/gconv"
	"gofly/utils/tools/gmap"
	"gofly/utils/tools/gvar"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// 代码生成
type Generatecode struct{ NoNeedAuths []string }

func init() {
	fpath := Generatecode{NoNeedAuths: []string{"*"}}
	gf.Register(&fpath, fpath)
}

// 获取列表
func (api *Generatecode) GetList(c *gf.GinCtx) {
	pageNo := gconv.Int(c.DefaultQuery("page", "1"))
	pageSize := gconv.Int(c.DefaultQuery("pageSize", "10"))
	//搜索添条件
	param, _ := gf.RequestParam(c)
	whereMap := gmap.New()
	whereMap.Set("status", 0)
	if name, ok := param["name"]; ok && name != "" {
		whereMap.Set("comment like ? OR cate_tablename  like ?", gf.Slice{"%" + gconv.String(name) + "%", "%" + gconv.String(name) + "%"})
	}
	if status, ok := param["status"]; ok && status != "" {
		whereMap.Set("status", status)
	}
	if createtime, ok := param["createtime"]; ok && createtime != "" {
		datetime_arr := gf.SplitAndStr(gf.String(createtime), ",")
		whereMap.Set("createtime between ? and ?", gf.Slice{datetime_arr[0] + " 00:00", datetime_arr[1] + " 23:59"})
	}
	MDB := gf.Model("common_generatecode").Where(whereMap)
	totalCount, _ := MDB.Clone().Count()
	list, err := MDB.Page(pageNo, pageSize).Order("id desc").Select()
	if err != nil {
		gf.Failed().SetMsg(err.Error()).Regin(c)
	} else {
		gf.Success().SetMsg("获取全部列表").SetData(gf.Map{
			"page":     pageNo,
			"pageSize": pageSize,
			"total":    totalCount,
			"items":    list}).Regin(c)
	}
}

// 获取数据表字段
func (api *Generatecode) GetDbfield(c *gf.GinCtx) {
	tablename := c.DefaultQuery("tablename", "")
	if tablename == "" {
		gf.Failed().SetMsg("请传数据表名称").Regin(c)
	} else {
		tablename_arr := strings.Split(tablename, ",")
		//获取数据库名
		var dielddata_list []map[string]interface{}
		for _, Val := range tablename_arr {
			dielddata, _ := gf.DB().Query(c, "select COLUMN_NAME,COLUMN_COMMENT,DATA_TYPE from information_schema.columns where TABLE_SCHEMA='"+gconv.String(dbConf_arr["dbname"])+"' AND TABLE_NAME='"+Val+"'")
			for _, data := range dielddata {
				if data["COLUMN_COMMENT"].String() == "" && data["COLUMN_NAME"].String() == "id" {
					data["COLUMN_COMMENT"] = gvar.New("ID")
				}
				dielddata_list = append(dielddata_list, map[string]interface{}{"value": data["COLUMN_NAME"], "label": data["COLUMN_COMMENT"], "type": data["DATA_TYPE"]})
			}
		}
		gf.Success().SetMsg("获取数据表字段").SetData(dielddata_list).Regin(c)
	}
}

// 获取数据库列表
func (api *Generatecode) GetTablelist(c *gf.GinCtx) {
	id := c.DefaultQuery("id", "")
	dbtalbelist, _ := gf.Model("common_generatecode").Where("status", 0).Where("id !=", id).Fields("tablename as value,comment as label").Order("id desc").Select()
	gf.Success().SetMsg("获取数据库列表").SetData(dbtalbelist).Regin(c)
}

// 获取字典数据库
func (api *Generatecode) GetDicTable(c *gf.GinCtx) {
	dbtalbelist, _ := gf.Model("common_dictionary_group").Where("status", 0).Fields("id,title").Select()
	gf.Success().SetMsg("获取字典数据库").SetData(dbtalbelist).Regin(c)
}

// 添加生成代码数据表前检查是否已经存在
func (api *Generatecode) CheckedHaseTable(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	tablename_arr, err := gf.Model("common_generatecode").WhereIn("tablename", param["tablenames"]).Array("tablename")
	if err != nil {
		gf.Failed().SetMsg("检查失败").SetData(err).Regin(c)
	} else {
		gf.Success().SetMsg("检查成功！").SetData(tablename_arr).Regin(c)
	}
}

// 更新生成代码的数据表
func (api *Generatecode) UpCodeTable(c *gf.GinCtx) {
	if appConf_arr["runEnv"] == "release" {
		gf.Failed().SetMsg("生产环境禁止操作，请在开发环境下操作").Regin(c)
		return
	}
	//获取数据库名
	param, _ := gf.RequestParam(c)
	gf.DB().Exec(c, "SET @@sql_mode='NO_ENGINE_SUBSTITUTION';")
	var dielddata_list gf.OrmResult
	for _, Val := range param["tablenames"].([]interface{}) {
		dbtalbelist, _ := gf.DB().Query(c, "SELECT TABLE_NAME,TABLE_COMMENT,ENGINE,TABLE_ROWS,TABLE_COLLATION,AUTO_INCREMENT FROM information_schema.TABLES WHERE TABLE_SCHEMA='"+gconv.String(dbConf_arr["dbname"])+"' AND TABLE_NAME='"+gconv.String(Val)+"'")
		for _, val := range dbtalbelist {
			webb, _ := json.Marshal(val)
			var webjson map[string]interface{}
			_ = json.Unmarshal(webb, &webjson)
			if val["TABLE_COMMENT"] == nil {
				val["TABLE_COMMENT"] = gvar.New("")
			}
			if val["AUTO_INCREMENT"] == nil {
				val["AUTO_INCREMENT"] = gvar.New("")
			}
			midata := gf.OrmRecord{
				"tablename":      val["TABLE_NAME"],
				"comment":        val["TABLE_COMMENT"],
				"rule_name":      val["TABLE_COMMENT"],
				"engine":         val["ENGINE"],
				"table_rows":     val["TABLE_ROWS"],
				"collation":      val["TABLE_COLLATION"],
				"auto_increment": val["AUTO_INCREMENT"],
			}
			dielddata_list = append(dielddata_list, midata)
		}
	}
	_, err := gf.Model("common_generatecode").Data(dielddata_list).Save()
	if err != nil {
		gf.Failed().SetMsg("更新失败").SetData(err).Regin(c)
	} else {
		gf.Success().SetMsg("添加成功！").Regin(c)
	}
}

// 判断元素是否存在数组中-泛型val并返回id
func IsContainValGetId(items gf.OrmResult, item *gvar.Var) int64 {
	for _, eachItem := range items {
		if eachItem["tablename"].String() == item.String() {
			return eachItem["id"].Int64()
		}
	}
	return 0
}

// 保存-生成代码
func (api *Generatecode) Save(c *gf.GinCtx) {
	if appConf_arr["runEnv"] == "release" {
		gf.Failed().SetMsg("生产环境禁止操作，请在开发环境下操作").Regin(c)
		return
	}
	param, _ := gf.RequestParam(c)
	codedata := param["codedata"].(map[string]interface{})
	field_list := param["field_list"].([]interface{})
	list := param["list"].([]interface{})
	search_list := param["search_list"].([]interface{})
	filename_arr := strings.Split(codedata["api_filename"].(string), `.`) //文件名称
	//更新字段列表数据
	for _, fval := range field_list {
		item := fval.(map[string]interface{})
		gf.Model("common_generatecode_field").Data(item).Where("id", item["id"]).Update()
	}
	for _, lval := range list {
		item := lval.(map[string]interface{})
		gf.Model("common_generatecode_field").Data(item).Where("id", item["id"]).Update()
	}
	for _, sval := range search_list {
		item := sval.(map[string]interface{})
		gf.Model("common_generatecode_field").Data(item).Where("id", item["id"]).Update()
	}
	auth_rule_talbename := "business_auth_rule"
	if codedata["codelocation"] == "adminDirName" {
		auth_rule_talbename = "admin_auth_rule"
	}
	//1生成菜单
	user_id, _ := c.Get("userID")
	findrule, _ := gf.Model(auth_rule_talbename).Where("type", 1).Where("routepath", codedata["routepath"]).WhereOr("routename", codedata["routename"]).Fields("id").Find()
	var isok = false
	if findrule == nil {
		save_arr := map[string]interface{}{
			"title": codedata["rule_name"], "type": 1, "uid": user_id,
			"icon": codedata["icon"], "routepath": codedata["routepath"], "routename": codedata["routename"],
			"pid": codedata["pid"], "component": codedata["component"],
		}
		getId, err := gf.Model(auth_rule_talbename).Data(save_arr).InsertAndGetId()
		if err != nil {
			gf.Failed().SetMsg("添加菜单失败").SetData(err).Regin(c)
		} else { //更新排序
			gf.Model(auth_rule_talbename).Data(gf.Map{"weigh": getId}).Where("id", getId).Update()
			codedata["rule_id"] = getId
			isok = true
			//添加权限数据
			pathFull := fmt.Sprintf("/%v/%v", codedata["api_path"], filename_arr[0])
			var addAuthList gf.Slice
			addAuthList = append(addAuthList, gf.Map{"title": "查看", "type": 2, "uid": user_id, "weigh": 1, "pid": getId, "path": pathFull + "/getList", "permission": "view"})
			addAuthList = append(addAuthList, gf.Map{"title": "添加/编辑", "type": 2, "uid": user_id, "weigh": 2, "pid": getId, "path": pathFull + "/save", "permission": "add"})
			addAuthList = append(addAuthList, gf.Map{"title": "删除", "type": 2, "uid": user_id, "weigh": 3, "pid": getId, "path": pathFull + "/del", "permission": "del"})
			addAuthList = append(addAuthList, gf.Map{"title": "状态", "type": 2, "uid": user_id, "weigh": 4, "pid": getId, "path": pathFull + "/upStatus", "permission": "status"})
			addAuthList = append(addAuthList, gf.Map{"title": "详情", "type": 2, "uid": user_id, "weigh": 5, "pid": getId, "path": pathFull + "/getContent", "permission": "details"})
			if codedata["tpl_type"] == "contentcatelist" || codedata["tpl_type"] == "sitecatelist" { //添加分类权限
				addAuthList = append(addAuthList, gf.Map{"title": "添加分类", "type": 2, "uid": user_id, "weigh": 6, "pid": getId, "path": pathFull + "cate/save", "permission": "addcate"})
				addAuthList = append(addAuthList, gf.Map{"title": "删除分类", "type": 2, "uid": user_id, "weigh": 7, "pid": getId, "path": pathFull + "cate/del", "permission": "delcate"})
				addAuthList = append(addAuthList, gf.Map{"title": "分类状态", "type": 2, "uid": user_id, "weigh": 8, "pid": getId, "path": pathFull + "cate/upStatus", "permission": "catestatus"})
			}
			gf.Model(auth_rule_talbename).Insert(addAuthList)
		}
	} else {
		isok = true
		codedata["rule_id"] = findrule["id"]
	}
	//菜单添加好后添加代码
	if isok {
		/***************************后端**************************/
		file_path := filepath.Join("app/", gconv.String(codedata["api_path"]))
		//1. 如果没有filepath文件目录就创建一个
		if _, err := os.Stat(file_path); err != nil {
			if !os.IsExist(err) {
				os.MkdirAll(file_path, os.ModePerm)
			}
		}
		//2. 替换文件内容
		packgename_arr := strings.Split(codedata["api_path"].(string), `/`)
		//2.1 模块名称
		modelname := "business"
		if len(packgename_arr) > 0 {
			modelname = packgename_arr[0]
		}
		//2.2 文件名称
		filename := "index"
		if len(filename_arr) > 0 {
			filename = filename_arr[0]
		}
		//2.3 包名
		packageName := ""
		if len(packgename_arr) > 0 {
			packageName = packgename_arr[len(packgename_arr)-1]
		}
		//创建后端代码
		fields_inter, _ := gf.Model("common_generatecode_field").Where("generatecode_id", codedata["id"]).Where("islist", 1).
			Order("list_weigh asc,id asc").Array("field")
		if fields_inter != nil {
			var str_arr = make([]string, len(fields_inter))
			for k, v := range fields_inter {
				str_arr[k] = fmt.Sprintf("%v", v)
			}
			codedata["fields"] = strings.Join(str_arr, ",")
		} else {
			codedata["fields"] = ""
		}
		//获取表单字段
		formfield, _ := gf.Model("common_generatecode_field").Where("generatecode_id", codedata["id"]).Where("isform", 1).
			Fields("id,name,field,required,formtype,datatable,datatablename,dic_group_id,def_value,option_value,gridwidth").Order("field_weigh asc,id asc").Select()
		//列表字段
		listfield, _ := gf.Model("common_generatecode_field").Where("generatecode_id", codedata["id"]).Where("islist", 1).
			Fields("id,name,formtype,field,align,width,show_ui,option_value,datatable,datatablename,dic_group_id").Order("list_weigh asc,id asc").Select()
		//获取搜索字段
		search_field, _ := gf.Model("common_generatecode_field").Where("generatecode_id", codedata["id"]).Where("issearch", 1).
			Fields("id,field,name,searchway,searchtype,searchwidth,option_value,datatable,datatablename,dic_group_id").Order("search_weigh asc,id asc").Select()
		go MarkeGoCode(file_path, filename, packageName, codedata, formfield, listfield, search_field)
		//3. 查看是否添加文件到控制器文件
		go CheckIsAddController(modelname, gconv.String(codedata["api_path"]), false)
		/******************************前端******************************/
		//更新配置
		cf, _ := gcfg.New()
		data := cf.MustGet(ctx, "app")
		appConf_arr = gconv.Map(data)
		component_arr := strings.Split(codedata["component"].(string), `/`)
		componentpah_arr := strings.Split(codedata["component"].(string), (component_arr[len(component_arr)-1]))
		code_main_dir := appConf_arr["busDirName"]
		if codedata["codelocation"] == "adminDirName" {
			code_main_dir = appConf_arr["adminDirName"]
		}
		vue_path := filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(code_main_dir), "/src/views/", componentpah_arr[0]) //前端文件路径
		//1. 如果没有filepath文件目录就创建一个
		if _, err := os.Stat(vue_path); err != nil {
			if !os.IsExist(err) {
				os.MkdirAll(vue_path, os.ModePerm)
			}
		}
		//2. 复制前端模板到新创建文件夹下
		CopyAllDir(filepath.Join("devsource/developer/codetpl/vue/", gconv.String(codedata["tpl_type"])), vue_path)
		//3. 修改模板文件内容
		if codedata["tpl_type"] == "contentcatelist" { //如果是关联分类则更新分类api.ts
			ApitsReplay(filepath.Join(vue_path, "cate/api.ts"), packageName, filename+"cate")
		} else if codedata["tpl_type"] == "sitecatelist" { //左侧菜单数据
			ApitsReplay(filepath.Join(vue_path, "api/cate.ts"), packageName, filename+"cate")
		}
		//修改api/index.ts文件
		ApitsReplay(filepath.Join(vue_path, "api/index.ts"), packageName, filename)
		//替换data.ts
		UpFieldData(filepath.Join(vue_path, "data.ts"), listfield)                                                                //更新data.ts
		UpFieldAddForm(filepath.Join(vue_path, "AddForm.vue"), codedata["fields"], formfield)                                     //更新表单
		UpFieldIndex(filepath.Join(vue_path, "index.vue"), search_field, listfield, gconv.String(codedata["tpl_type"]), filename) //index.vue列表
		/*************最后更新代码生成表数据***************************/
		codedata["is_install"] = 1
		res, err := gf.Model("common_generatecode").Data(codedata).Where("id", codedata["id"]).Update()
		if err != nil {
			gf.Failed().SetMsg("更新失败").SetData(err).Regin(c)
		} else {
			gf.Success().SetMsg("更新成功！").SetData(res).Regin(c)
		}
	} else {
		gf.Failed().SetMsg("添加菜单失败").Regin(c)
	}

}

// 更新状态
func (api *Generatecode) UpStatus(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	res2, err := gf.Model("common_generatecode").Where("id", param["id"]).Data(map[string]interface{}{"status": param["status"]}).Update()
	if err != nil {
		gf.Failed().SetMsg("更新失败！").SetData(err).Regin(c)
	} else {
		msg := "更新成功！"
		if res2 == nil {
			msg = "暂无数据更新"
		}
		gf.Success().SetMsg(msg).SetData(res2).Regin(c)
	}
}

// 删除/卸载
func (api *Generatecode) Del(c *gf.GinCtx) {
	if appConf_arr["runEnv"] == "release" {
		gf.Failed().SetMsg("生产环境禁止操作，请在开发环境下操作").Regin(c)
		return
	}
	param, _ := gf.RequestParam(c)
	if param["is_install"] != nil && gconv.Int(param["is_install"]) == 1 { //卸载
		isok, err := common_uninstall(param["id"])
		if isok {
			gf.Model("common_generatecode").Where("id", param["id"]).Data(map[string]interface{}{"is_install": 2}).Update()
			gf.Success().SetMsg("卸载成功！").Regin(c)
		} else {
			gf.Failed().SetMsg("卸载失败").SetData(err).Regin(c)
		}
	} else { //删除
		res2, err := gf.Model("common_generatecode").Where("id", param["id"]).Delete()
		if err != nil {
			gf.Failed().SetMsg("删除失败").SetData(err).Regin(c)
		} else {
			gf.Model("common_generatecode_field").Where("generatecode_id", param["id"]).Delete()
			gf.Success().SetMsg("删除成功！").SetData(res2).Regin(c)
		}
	}
}

// 卸载
func (api *Generatecode) Uninstallcode(c *gf.GinCtx) {
	if appConf_arr["runEnv"] == "release" {
		gf.Failed().SetMsg("生产环境禁止操作，请在开发环境下操作").Regin(c)
		return
	}
	param, _ := gf.RequestParam(c)
	isok, err := common_uninstall(param["id"])
	if isok {
		gf.Success().SetMsg("卸载成功！").SetData(0).Regin(c)
	} else {
		gf.Failed().SetMsg("卸载失败").SetData(err).Regin(c)
	}
}

// 卸载通用方法
func common_uninstall(id interface{}) (bool, error) {
	data, err := gf.Model("common_generatecode").Where("id", id).Fields("rule_id,api_path,api_filename,cate_tablename,tpl_type,component,codelocation").Find()
	if err != nil {
		return false, err
	} else {
		file_path := filepath.Join("app/", gconv.String(data["api_path"]))
		//判断后端代码是否存在删除后端代码
		filego_path := filepath.Join(file_path, gconv.String(data["api_filename"]))
		if _, err := os.Stat(filego_path); err == nil {
			//删除菜单
			auth_rule_talbename := "business_auth_rule"
			if data["codelocation"].String() == "adminDirName" {
				auth_rule_talbename = "admin_auth_rule"
			}
			gf.Model(auth_rule_talbename).Where("id", data["rule_id"]).Delete()
			gf.Model(auth_rule_talbename).Where("pid", data["rule_id"]).Delete()
			gf.Model("common_generatecode").Data(map[string]interface{}{"is_install": 0}).Where("id", id).Update()
			go UnInstallCodeFile(data)
		}
		return true, nil
	}
}

// 获取内容
func (api *Generatecode) GetContent(c *gf.GinCtx) {
	id := c.DefaultQuery("id", "")
	if id == "" {
		gf.Failed().SetMsg("请传参数id").Regin(c)
	} else {
		data, err := gf.Model("common_generatecode").Fields("id,tablename,comment,pid,rule_id,rule_name,icon,is_install,routepath,routename,component,api_path,api_filename,cate_tablename,tpl_type,codelocation").Where("id", id).Find()
		if err != nil {
			gf.Failed().SetMsg("获取内容失败").SetData(err).Regin(c)
		} else {
			if data == nil {
				gf.Failed().SetMsg("生成数据表不存在").SetData(err).Regin(c)
			} else {
				var dielddata_list []map[string]interface{}
				var haseids []interface{}
				dielddata, _ := gf.DB().Query(c, "select COLUMN_NAME,COLUMN_COMMENT,DATA_TYPE,CHARACTER_MAXIMUM_LENGTH,COLUMN_DEFAULT from information_schema.columns where TABLE_SCHEMA='"+gconv.String(dbConf_arr["dbname"])+"' AND TABLE_NAME='"+data["tablename"].String()+"'")
				for _, data := range dielddata {
					if data["COLUMN_COMMENT"].String() == "" && data["COLUMN_NAME"].String() == "id" {
						data["COLUMN_COMMENT"] = gvar.New("ID")
					}
					formtype := "text"
					gridwidth := 12
					width := 100
					show_ui := ""
					isorder := 0
					issearch := 0
					searchway := "="
					searchtype := "text"
					if data["COLUMN_NAME"].String() == "id" {
						isorder = 1
						width = 80
					}
					if strings.HasSuffix(data["COLUMN_NAME"].String(), "date") {
						formtype = "date"
						width = 160
					} else if strings.HasSuffix(data["COLUMN_NAME"].String(), "datetime") {
						formtype = "datetime"
						width = 160
					} else if strings.HasSuffix(data["COLUMN_NAME"].String(), "time") {
						formtype = "time"
						width = 120
					} else if strings.HasSuffix(data["COLUMN_NAME"].String(), "color") {
						formtype = "colorpicker"
						show_ui = "color"
						width = 80
					} else if strings.HasSuffix(data["COLUMN_NAME"].String(), "image") {
						gridwidth = 24
						formtype = "image"
						show_ui = "image"
						width = 80
					} else if strings.HasSuffix(data["COLUMN_NAME"].String(), "images") {
						formtype = "images"
						show_ui = "images"
						gridwidth = 24
					} else if strings.HasSuffix(data["COLUMN_NAME"].String(), "audio") {
						formtype = "audio"
						gridwidth = 24
					} else if strings.HasSuffix(data["COLUMN_NAME"].String(), "file") {
						formtype = "file"
						gridwidth = 24
					} else if strings.HasSuffix(data["COLUMN_NAME"].String(), "files") {
						formtype = "files"
						gridwidth = 24
					} else if (strings.HasSuffix(data["COLUMN_NAME"].String(), "gender") || strings.HasSuffix(data["COLUMN_NAME"].String(), "sex")) && data["DATA_TYPE"].String() == "tinyint" {
						formtype = "radio"
						show_ui = "gender"
						width = 110
					} else if data["DATA_TYPE"].String() == "int" {
						formtype = "number"
					} else if data["DATA_TYPE"].String() == "varchar" && data["CHARACTER_MAXIMUM_LENGTH"].Int64() <= 50 {
						width = 190
					} else if data["DATA_TYPE"].String() == "varchar" && data["CHARACTER_MAXIMUM_LENGTH"].Int64() > 50 && data["CHARACTER_MAXIMUM_LENGTH"].Int64() < 225 {
						width = 250
					} else if data["DATA_TYPE"].String() == "varchar" && data["CHARACTER_MAXIMUM_LENGTH"].Int64() >= 225 {
						formtype = "textarea"
						width = 280
						show_ui = "des"
						gridwidth = 24
					} else if data["DATA_TYPE"].String() == "text" || data["DATA_TYPE"].String() == "longtext" {
						formtype = "editor"
						show_ui = "des"
					} else if data["DATA_TYPE"].String() == "enum" {
						formtype = "select"
						show_ui = "tags"
						searchtype = "select"
					} else if data["DATA_TYPE"].String() == "tinyint" {
						formtype = "radio"
						show_ui = "tag"
					}
					//备注
					option_value := ""
					name_value := data["COLUMN_COMMENT"].String()
					if strings.Contains(name_value, ":") {
						name_arr := strings.Split(name_value, ":")
						name_value = name_arr[0]
						option_value = name_arr[1]
					}
					if fieldval, _ := gf.Model("common_generatecode_field").Where("generatecode_id", id).Where("field", data["COLUMN_NAME"]).Value("id"); fieldval != nil {
						haseids = append(haseids, fieldval)
						if option_value != "" {
							dielddata_list = append(dielddata_list, gf.Map{"id": fieldval, "option_value": option_value})
						}
					} else {
						if data["COLUMN_DEFAULT"] == nil {
							data["COLUMN_DEFAULT"] = gvar.New("")
						}
						if data["DATA_TYPE"].String() == "varchar" && (data["COLUMN_NAME"].String() == "name" || data["COLUMN_NAME"].String() == "title") {
							issearch = 1
						}
						if data["DATA_TYPE"].String() == "varchar" && data["COLUMN_NAME"].String() == "status" {
							issearch = 1
						}
						if data["DATA_TYPE"].String() == dbConf_arr["createdAt"] {
							issearch = 1
						}
						if data["DATA_TYPE"].String() == dbConf_arr["createdAt"] || data["COLUMN_NAME"].String() == dbConf_arr["updatedAt"] || data["COLUMN_NAME"].String() == dbConf_arr["deletedAt"] {
							searchway = "between"
							searchtype = "daterange"
						}
						maxid, _ := gf.Model("common_generatecode_field").Where("generatecode_id", id).OrderDesc("id").Value("id")
						dielddata_list = append(dielddata_list, gf.Map{"generatecode_id": id, "name": name_value, "option_value": option_value, "field": data["COLUMN_NAME"], "formtype": formtype, "gridwidth": gridwidth, "def_value": data["COLUMN_DEFAULT"], "isorder": isorder, "issearch": issearch, "searchway": searchway, "searchtype": searchtype, "field_weigh": maxid, "list_weigh": maxid, "search_weigh": maxid, "width": width, "show_ui": show_ui})
					}
				}
				if haseids != nil {
					gf.Model("common_generatecode_field").Where("generatecode_id", id).WhereNotIn("id", haseids).Delete()
				}
				if dielddata_list != nil {
					_, err := gf.Model("common_generatecode_field").Data(dielddata_list).Save()
					if err != nil {
						gf.Failed().SetMsg("新增字段失败").SetData(err).Regin(c)
						return
					}
				}
				field_list, _ := gf.Model("common_generatecode_field").Where("generatecode_id", id).
					Fields("id,isform,name,field,required,formtype,datatable,datatablename,dic_group_id,field_weigh,gridwidth").Order("field_weigh asc,id asc").Select()
				for _, fval := range field_list {
					if fval["isform"].Int() == 1 {
						fval["isform"] = gvar.New(true)
					} else {
						fval["isform"] = gvar.New(false)
					}
					if fval["required"].Int() == 1 {
						fval["required"] = gvar.New(true)
					} else {
						fval["required"] = gvar.New(false)
					}
					if strings.Contains(fval["name"].String(), ":") {
						fval["name"] = gvar.New(strings.Split(fval["name"].String(), ":")[0])
					}
				}
				list, _ := gf.Model("common_generatecode_field").Where("generatecode_id", id).
					Fields("id,islist,name,field,isorder,align,width,show_ui,list_weigh").Order("list_weigh asc,id asc").Select()
				for _, lval := range list {
					if lval["islist"].Int() == 1 {
						lval["islist"] = gvar.New(true)
					} else {
						lval["islist"] = gvar.New(false)
					}
					if lval["isorder"].Int() == 1 {
						lval["isorder"] = gvar.New(true)
					} else {
						lval["isorder"] = gvar.New(false)
					}
				}
				search_list, _ := gf.Model("common_generatecode_field").Where("generatecode_id", id).
					Fields("id,issearch,name,searchway,searchtype,search_weigh,searchwidth").Order("search_weigh asc,id asc").Select()
				for _, sval := range search_list {
					if sval["issearch"].Int() == 1 {
						sval["issearch"] = gvar.New(true)
					} else {
						sval["issearch"] = gvar.New(false)
					}
				}
				haseadmin := false
				vue_viewsfiles_path_admin := filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(appConf_arr["adminDirName"]))
				if _, err := os.Stat(vue_viewsfiles_path_admin); err == nil {
					haseadmin = true
				}
				gf.Success().SetMsg("获取生成表单信息成功！").SetData(gf.Map{"data": data, "haseadmin": haseadmin, "field_list": field_list, "list": list, "search_list": search_list}).Regin(c)
			}
		}
	}

}

// 获取菜单数据
func (api *Generatecode) GetMenuParent(c *gf.GinCtx) {
	codelocation := c.DefaultQuery("codelocation", "busDirName")
	tablename := "business_auth_rule"
	if codelocation == "adminDirName" {
		tablename = "admin_auth_rule"
	}
	menuList, err := gf.Model(tablename).WhereIn("type", []interface{}{0, 1}).Fields("id,pid,title,locale,routepath").Order("weigh asc").Select()
	if err != nil {
		gf.Failed().SetMsg("获取数据失败").SetData(err).Regin(c)
	} else {
		if menuList == nil {
			menuList = make(gf.OrmResult, 0)
		}
		for _, val := range menuList {
			if val["title"].String() == "" {
				val["title"] = val["locale"]
			}
		}
		menuList = gf.GetMenuChildrenArray(menuList, 0, "pid")
		ids, _ := gf.Model(tablename).WhereIn("type", []interface{}{0, 1}).Array("id")
		gf.Success().SetMsg("菜单父级数据！").SetData(gf.Map{"list": menuList, "ids": ids}).Regin(c)
	}
}

// 代码生成工具-获取前后端代码目录
func (api *Generatecode) GetGoVueDir(c *gf.GinCtx) {
	codelocation := c.DefaultQuery("codelocation", "busDirName")
	tablename := appConf_arr["busDirName"]
	if codelocation == "adminDirName" {
		tablename = appConf_arr["adminDirName"]
	}
	path, _ := os.Getwd()
	//后端目录
	_, go_app_arr, _ := GetAppFileNoAdmin(filepath.Join(path, "/app"))
	go_app_dir, _ := gf.GetAllFileArray(filepath.Join(path, "/app"))
	//前端目录
	vue_views, _ := gf.GetAllFileArray(filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(tablename), "/src/views"))
	//判断是否admin
	haseadmin := false
	vue_viewsfiles_path_admin := filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(appConf_arr["adminDirName"]))
	if _, err := os.Stat(vue_viewsfiles_path_admin); err == nil {
		haseadmin = true
	}
	gf.Success().SetMsg("获取前后端代码目录").SetData(gf.Map{"go_app_dir": go_app_dir["folders"], "go_file_dir": go_app_arr, "views_dir": vue_views["folders"], "haseadmin": haseadmin}).Regin(c)
}

// 代码生成工具-生成代码
func (api *Generatecode) MarkeTplCode(c *gf.GinCtx) {
	if appConf_arr["runEnv"] == "release" {
		gf.Failed().SetMsg("生产环境禁止操作，请在开发环境下操作").Regin(c)
		return
	}
	param, _ := gf.RequestParam(c)
	vue_dir := appConf_arr["busDirName"]
	auth_rule_talbename := "business_auth_rule"
	if param["codelocation"] == "adminDirName" {
		vue_dir = appConf_arr["adminDirName"]
		auth_rule_talbename = "admin_auth_rule"
	}
	//处理变量-后端变量
	goappdir := gf.String(param["goappdir"])     //go模块
	godir_name := gf.String(param["godir"])      //go文件类/空间名
	gofilename := gf.String(param["gofilename"]) //go文件名称
	api_path := godir_name                       //go路径目录
	//前端表里
	viewsdir := gf.String(param["viewsdir"]) //views下面路由
	vuedir := gf.String(param["vuedir"])     //新来功能目录
	packageName := vuedir                    //模块名称
	//拼接路径
	if goappdir != "" {
		api_path = goappdir + "/" + godir_name
	}
	//生成的go文件名
	filename := "index"
	if gofilename != "" {
		filename = gofilename
	}
	//生成菜单
	if param["rule_name"] != "" {
		user_id, _ := c.Get("userID")
		findrule_id, _ := gf.Model(auth_rule_talbename).Where("type", 1).Where("routepath", param["routepath"]).WhereOr("routename", param["routename"]).Value("id")
		if findrule_id == nil {
			save_arr := gf.Map{"createtime": time.Now().Unix(),
				"title": param["rule_name"], "type": 1, "uid": user_id,
				"icon": param["icon"], "routepath": param["routepath"], "routename": param["routename"],
				"pid": param["pid"], "component": param["component"],
			}
			getId, err := gf.Model(auth_rule_talbename).Data(save_arr).InsertAndGetId()
			if err != nil {
				gf.Failed().SetMsg("添加菜单失败").SetData(err).Regin(c)
			} else { //更新排序
				param["rule_id"] = getId
				gf.Model(auth_rule_talbename).Data(gf.Map{"weigh": getId}).Where("id", getId).Update()
			}
		} else {
			save_arr := gf.Map{"createtime": time.Now().Unix(),
				"title": param["rule_name"], "type": 1, "uid": user_id,
				"icon": param["icon"], "routepath": param["routepath"], "routename": param["routename"],
				"pid": param["pid"], "component": param["component"],
			}
			param["rule_id"] = findrule_id
			gf.Model(auth_rule_talbename).Data(save_arr).Where("id", findrule_id).Update()
		}
	} else {
		param["rule_id"] = 0
	}
	//处理后端
	go_file_path := ""
	if param["codetpl"] == "go" || param["codetpl"] == "goweb" {
		go_file_path = filepath.Join("app/", goappdir, godir_name)
		if _, err := os.Stat(go_file_path); err != nil {
			if !os.IsExist(err) {
				os.MkdirAll(go_file_path, os.ModePerm)
			}
		}
		//创建go文件并替换内容
		go MarkeGoCodeTpl(go_file_path, filename, godir_name)
		//3. 查看是否添加文件到控制器文件
		go CheckIsAddController(goappdir, api_path, false)
	}
	//处理前端
	vue_path := ""
	if param["codetpl"] == "web" || param["codetpl"] == "goweb" {
		vue_path = filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(vue_dir), "/src/views/", viewsdir, vuedir) //前端文件路径
		//1. 如果没有filepath文件目录就创建一个
		if _, err := os.Stat(vue_path); err != nil {
			if !os.IsExist(err) {
				os.MkdirAll(vue_path, os.ModePerm)
			}
		}
		//2. 复制前端模板到新创建文件夹下
		if param["codetpl"] == "goweb" { //有请求api
			CopyAllDir(filepath.Join("devsource/developer/codetpl/vue/vuegotpl"), vue_path)
			//修改api/index.ts文件
			router_main := godir_name
			if goappdir != "business" && goappdir != "admin" {
				router_main = "./" + router_main
			} else {
				router_main = "/" + router_main
			}
			ApitsReplayTpl(filepath.Join(vue_path, "api/index.ts"), router_main, filename)
		} else {
			CopyAllDir(filepath.Join("devsource/developer/codetpl/vue/vuetpl"), vue_path)
		}
		UpFieldIndexTpl(filepath.Join(vue_path, "index.vue"), packageName)
	}
	// 保存生成记录
	commentstr := ""
	if param["codetpl"] == "web" {
		commentstr = fmt.Sprintf("代码工具生成前端%v", packageName)
	} else if param["codetpl"] == "go" {
		commentstr = fmt.Sprintf("代码工具生成后端%v", filename)
	} else {
		commentstr = fmt.Sprintf("代码工具生成前端%v后端%v", packageName, filename)
	}
	_, err := gf.Model("common_generatecode").Data(gf.Map{
		"fromtype":       1,
		"is_install":     1,
		"comment":        commentstr,
		"codelocation":   param["codelocation"],
		"pid":            param["pid"],
		"rule_id":        param["rule_id"],
		"rule_name":      param["rule_name"],
		"routename":      param["routename"],                     //路由名称
		"component":      fmt.Sprintf("%v/%v", viewsdir, vuedir), //记录前端生成位置
		"godir":          go_file_path,                           //go后端生成位置
		"api_filename":   filename + ".go",                       //go文件名称
		"cate_tablename": goappdir,                               //用cate_tablename存go模块名
		"api_path":       api_path,                               //控制器path
		"tpl_type":       param["codetpl"],                       //模板类型

	}).InsertAndGetId()
	gf.Success().SetMsg("生成代码成功").SetData(param).SetExdata(err).Regin(c)
}

// 删除模板代码
func (api *Generatecode) DelTplCode(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	data, err := gf.Model("common_generatecode").Where("id", param["id"]).Fields("rule_id,api_path,api_filename,cate_tablename,tpl_type,component,codelocation").Find()
	if err != nil {
		gf.Failed().SetMsg("查找数据失败").SetData(err).Regin(c)
		return
	}
	vue_dir := appConf_arr["busDirName"]
	auth_rule_talbename := "business_auth_rule"
	if data["codelocation"].String() == "adminDirName" {
		vue_dir = appConf_arr["adminDirName"]
		auth_rule_talbename = "admin_auth_rule"
	}
	//删除后端
	if data["tpl_type"].String() == "go" || data["tpl_type"].String() == "goweb" {
		go_file_root_dir := filepath.Join("app/", data["api_path"].String())
		//删除代码
		filego_path := filepath.Join("app/", data["api_path"].String(), data["api_filename"].String())
		if _, err := os.Stat(filego_path); err == nil {
			//1.文件存在删除文件
			os.Remove(filego_path)
		}
		//删除控制器文件
		dir, _ := os.ReadDir(go_file_root_dir)
		if len(dir) == 0 {
			os.RemoveAll(go_file_root_dir)
			//3.移除路由
			CheckApiRemoveController(data["cate_tablename"].String(), data["api_path"].String())
		}
	}
	//删除前端代码
	if data["tpl_type"].String() == "web" || data["tpl_type"].String() == "goweb" {
		vue_path := filepath.Join(gf.String(appConf_arr["vueobjroot"]), gf.String(vue_dir), "/src/views/", data["component"].String()) //前端文件路径
		if _, err := os.Stat(vue_path); err == nil {
			os.RemoveAll(vue_path)
		}
	}
	if data["rule_id"].Int() > 0 {
		//删除菜单
		gf.Model(auth_rule_talbename).Where("id", data["rule_id"]).Delete()
	}
	res, err := gf.Model("common_generatecode").Where("id", param["id"]).Delete()
	if err != nil {
		gf.Failed().SetMsg("删除失败").SetData(err).Regin(c)
	} else {
		gf.Success().SetMsg("删除成功！").SetData(res).Regin(c)
	}
}
