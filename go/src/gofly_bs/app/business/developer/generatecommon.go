package developer

import (
	"bufio"
	"fmt"
	"gofly/utils/gf"
	"gofly/utils/tools/garray"
	"gofly/utils/tools/gstr"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

/**
* 代码安装工具
 */
/********************************后端*****************************************/
// tablename, tablenamecate, fields string  parameter=codedata前端参数
func MarkeGoCode(file_path, filename, packageName string, parameter map[string]interface{}, tableformlist, tablefieldlist, search_field gf.OrmResult) {
	ReplayToJSON := ""       //数组字符串转json对象
	ReplayFieldVal := ""     //添加关联表对应值显示
	ReplayFieldVal_mid := "" //添加关联表对应值显示过度
	ReplayArrayToStr := ""   //保存是数组转字符串
	ReplayURL := ""          //添加附件地址域名
	ReplaySearch := ""       //搜索查找
	//查找
	for _, searchjson := range search_field {
		field_str := searchjson["field"].String()
		searchtype_str := searchjson["searchtype"].String()
		searchway_str := searchjson["searchway"].String() //查找方式 = > like
		if ReplaySearch != "" {
			ReplaySearch += "\n"
		}
		//判断字段是否为保留关键字，如果是go预留关键字则关键字+_s
		fieldName := field_str
		bls := garray.NewStrArray()
		bls.SetArray(gf.SliceStr{"import", "package", "chan", "const", "func", "interface", "map", "struct", "type", "var", "break", "case", "continue", "default", "defer", "else", "fallthrough", "for", "go", "goto", "if", "range", "return", "select", "switch"}) //25个保留关键字
		if bls.Contains("field_str") {
			fieldName = field_str + "_s"
		}
		if searchway_str == "=" && searchtype_str == "belongto" {
			if field_str != "cid" {
				ReplaySearch += fmt.Sprintf("\tif %v, ok := param[\"%v\"]; ok && gf.Int(%v) != 0 {\n\t\t%vs := gf.CateAllChilId(\"%v\", %v)\n\t\twhereMap.Set(\"%v In(?)\", %vs)\n\t}", fieldName, field_str, fieldName, field_str, searchjson["datatable"], field_str, field_str, fieldName)
			}
		} else if searchway_str == "=" {
			ReplaySearch += fmt.Sprintf("\tif %v, ok := param[\"%v\"]; ok && %v != \"\" {\n\t\twhereMap.Set(\"%v\", %v)\n\t}", fieldName, field_str, fieldName, field_str, fieldName)
		} else if searchway_str == "!=" {
			ReplaySearch += fmt.Sprintf("\tif %v, ok := param[\"%v\"]; ok && %v != \"\" {\n\t\twhereMap.Set(\"%v != ?\", %v)\n\t}", fieldName, field_str, fieldName, field_str, fieldName)
		} else if searchway_str == ">" {
			ReplaySearch += fmt.Sprintf("\tif %v, ok := param[\"%v\"]; ok && %v != \"\" {\n\t\twhereMap.Set(\"%v > ?\", %v)\n\t}", fieldName, field_str, fieldName, field_str, fieldName)
		} else if searchway_str == ">=" {
			ReplaySearch += fmt.Sprintf("\tif %v, ok := param[\"%v\"]; ok && %v != \"\" {\n\t\twhereMap.Set(\"%v >= ?\", %v)\n\t}", fieldName, field_str, fieldName, field_str, fieldName)
		} else if searchway_str == "<" {
			ReplaySearch += fmt.Sprintf("\tif %v, ok := param[\"%v\"]; ok && %v != \"\" {\n\t\twhereMap.Set(\"%v < ?\", %v)\n\t}", fieldName, field_str, fieldName, field_str, fieldName)
		} else if searchway_str == "<=" {
			ReplaySearch += fmt.Sprintf("\tif %v, ok := param[\"%v\"]; ok && %v != \"\" {\n\t\twhereMap.Set(\"%v <= ?\", %v)\n\t}", fieldName, field_str, fieldName, field_str, fieldName)
		} else if searchway_str == "like" {
			ReplaySearch += fmt.Sprintf("\tif %v, ok := param[\"%v\"]; ok && %v != \"\" {\n\t\twhereMap.Set(\"%v like ?\", \"%%\"+gf.String(%v)+\"%%\")\n\t}", fieldName, field_str, fieldName, field_str, fieldName)
		} else if searchway_str == "in" {
			ReplaySearch += fmt.Sprintf("\tif %v, ok := param[\"%v\"]; ok && %v != \"\" {\n\t\twhereMap.Set(\"%v IN(?)\", gf.SplitAndStr(gf.String(%v), \",\"))\n\t}", fieldName, field_str, fieldName, field_str, fieldName)
		} else if searchway_str == "set" {
			ReplaySearch += fmt.Sprintf("\tif %v, ok := param[\"%v\"]; ok && %v != \"\" {\n\t\twhereMap.Set(\"find_in_set(?, %v)\", %v)\n\t}", fieldName, field_str, fieldName, field_str, fieldName)
		} else if searchway_str == "between" && searchtype_str == "daterange" {
			ReplaySearch += fmt.Sprintf("\tif %v, ok := param[\"%v\"]; ok && %v != \"\" {\n\t\tdatetime_arr := gf.SplitAndStr(gf.String(%v), \",\")\n\t\twhereMap.Set(\"%v between ? and ?\", gf.Slice{datetime_arr[0] + \" 00:00\", datetime_arr[1] + \" 23:59\"})\n\t}", fieldName, field_str, fieldName, field_str, fieldName)
		}
	}
	if ReplaySearch == "\n" {
		ReplaySearch = ""
	}
	//表单
	for _, webjson := range tableformlist {
		value_str := webjson["field"].String()
		type_str := webjson["formtype"].String()
		if type_str == "checkbox" {
			if ReplayArrayToStr != "" {
				ReplayArrayToStr += "\n"
			}
			if ReplayToJSON != "" {
				ReplayToJSON += "\n"
			}
			//保存时
			ReplayArrayToStr += fmt.Sprintf("\tparam[\"%v\"] = gf.ArrayToStr(param[\"%v\"], \",\")", value_str, value_str)
			//获取详情
			ReplayToJSON += fmt.Sprintf("\t\t\tif data != nil && data[\"%v\"].String() != \"\" {\n\t\t\t\tdata[\"%v\"] = gf.VarNew(gf.SplitAndStr(data[\"%v\"].String(), \",\"))\n\t\t\t}", value_str, value_str, value_str)
		}
	}
	//列表
	for _, webjson := range tablefieldlist {
		fvalue_str := webjson["field"].String()
		ftype_str := webjson["formtype"].String()
		if ftype_str == "belongto" {
			ReplayFieldVal_mid += fmt.Sprintf("\n\t\t\tval[\"%vName\"] = gf.GetTalbeFieldVal(\"%v\", \"%v\", val[\"%v\"])", fvalue_str, webjson["datatable"], webjson["datatablename"], fvalue_str)
		} else if ftype_str == "belongDic" {
			ReplayFieldVal_mid += fmt.Sprintf("\n\t\t\tval[\"%v\"] = gf.GetDicFieldVal(\"%v\", val[\"%v\"])", fvalue_str, webjson["dic_group_id"], fvalue_str)
		} else if ftype_str == "checkbox" {
			ReplayFieldVal_mid += fmt.Sprintf("\n\t\t\tif val[\"%v\"].String() != \"\" {\n\t\t\t\tval[\"%v\"] = gf.VarNew(gf.SplitAndStr(val[\"%v\"].String(), \",\"))\n\t\t\t}", fvalue_str, fvalue_str, fvalue_str)
		}
	}
	//出来列表查询
	if ReplayFieldVal_mid != "" {
		if ReplayURL != "" {
			ReplayFieldVal = ReplayURL + "\n\t\tfor _, val := range list {" + ReplayFieldVal_mid + "\n\t\t}"
		} else {
			ReplayFieldVal = "\n\t\tfor _, val := range list {" + ReplayFieldVal_mid + "\n\t\t}"
		}
	}
	//变量参数
	tablename := gf.String(parameter["tablename"])
	whereBusinessId := ""     //替换business_id查找条件
	whereMapBusinessId := ""  //替换business_id查找条件
	paramSaveBusinessId := "" //保存添加business_id
	if parameter["codelocation"] == "busDirName" && gf.DbHaseField(tablename, "business_id") {
		whereBusinessId = ".Where(\"business_id\", c.GetInt64(\"businessID\"))"
		whereMapBusinessId = "\twhereMap.Set(\"business_id\", c.GetInt64(\"businessID\"))"
		paramSaveBusinessId = "\t\tparam[\"business_id\"] = c.GetInt64(\"businessID\") //当前用户商户ID"
	}
	//分类数据表
	tablenamecate := gf.String(parameter["cate_tablename"])
	fields := gf.String(parameter["fields"])
	// 创建go文件
	filePath := filepath.Join(file_path, filename+".go")
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			os.Create(filePath)
		}
	}
	//复制go文件模板到新创建文件
	copyfile := "list"
	if parameter["tpl_type"] != "" {
		copyfile = gf.String(parameter["tpl_type"])
		if strings.Contains(parameter["tpl_type"].(string), "cate") {
			filename_cate := filename + "cate"
			filePath_cate := filepath.Join(file_path, filename_cate+".go")
			MarkeBelongCate(filePath_cate, filename_cate, packageName, tablenamecate, fields, whereBusinessId, whereMapBusinessId, paramSaveBusinessId)
		}
	}
	err := CopyFileContents(filepath.Join("devsource/developer/codetpl/go/", copyfile+".gos"), filePath)
	if err != nil {
		panic(err)
	}
	//打开新键go文件内容-并替换
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
		if strings.Contains(string(a), "Replace") {
			datestr := strings.ReplaceAll(string(a), "Replace", gstr.UcFirst(filename))
			result += datestr + "\n"
		} else if strings.Contains(string(a), "packageName") {
			datestr := strings.ReplaceAll(string(a), "packageName", packageName)
			result += datestr + "\n"
		} else if strings.Contains(string(a), "//{ReplaySearch}") { //搜索条件
			datestr := strings.ReplaceAll(string(a), "//{ReplaySearch}", ReplaySearch)
			result += datestr + "\n"
		} else if strings.Contains(string(a), "//{ReplayToJSON}") { //数组字符串转json对象
			datestr := strings.ReplaceAll(string(a), "//{ReplayToJSON}", ReplayToJSON)
			result += datestr + "\n"
		} else if strings.Contains(string(a), "//{ReplayArrayToStr}") { //保存是数组转字符串,号分隔
			datestr := strings.ReplaceAll(string(a), "//{ReplayArrayToStr}", ReplayArrayToStr)
			result += datestr + "\n"
		} else if strings.Contains(string(a), "//{ReplayFieldVal}") { //添加关联表对应值显示
			datestr := strings.ReplaceAll(string(a), "//{ReplayFieldVal}", ReplayFieldVal)
			result += datestr + "\n"
		} else if strings.Contains(string(a), "{tablename}") {
			datestr := strings.ReplaceAll(string(a), "{tablename}", tablename)
			if strings.Contains(string(a), "{ReplayWhereBusinessId}") {
				datestr = strings.ReplaceAll(datestr, "{ReplayWhereBusinessId}", whereBusinessId)
			}
			result += datestr + "\n"
		} else if strings.Contains(string(a), "{tablenamecate}") {
			datestr := strings.ReplaceAll(string(a), "{tablenamecate}", tablenamecate)
			result += datestr + "\n"
		} else if strings.Contains(string(a), "//{ReplayWhereBusinessId}") { //替换business_id查找条件map
			datestr := strings.ReplaceAll(string(a), "//{ReplayWhereBusinessId}", whereMapBusinessId)
			result += datestr + "\n"
		} else if strings.Contains(string(a), "//{ReplaySaveBusinessId}") { //保存添加business_id
			datestr := strings.ReplaceAll(string(a), "//{ReplaySaveBusinessId}", paramSaveBusinessId)
			result += datestr + "\n"
		} else if strings.Contains(string(a), "{fields}") {
			datestr := strings.ReplaceAll(string(a), "{fields}", fields)
			result += datestr + "\n"
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

// 创建数据关联的分类
func MarkeBelongCate(filePath, filename, packageName, tablename, fields, whereBusinessId, whereMapBusinessId, paramSaveBusinessId string) {
	err := CopyFileContents(filepath.Join("devsource/developer/codetpl/go/contentcate.gos"), filePath)
	if err != nil {
		panic(err)
	}
	//打开新键go文件内容-并替换
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
		if strings.Contains(string(a), "Replace") {
			datestr := strings.ReplaceAll(string(a), "Replace", gstr.UcFirst(filename))
			result += datestr + "\n"
		} else if strings.Contains(string(a), "packageName") {
			datestr := strings.ReplaceAll(string(a), "packageName", packageName)
			result += datestr + "\n"
		} else if strings.Contains(string(a), "{tablename}") {
			datestr := strings.ReplaceAll(string(a), "{tablename}", tablename)
			if strings.Contains(string(a), "{ReplayWhereBusinessId}") {
				datestr = strings.ReplaceAll(datestr, "{ReplayWhereBusinessId}", whereBusinessId)
			}
			result += datestr + "\n"
		} else if strings.Contains(string(a), "//{ReplayWhereBusinessId}") { //替换business_id查找条件map
			datestr := strings.ReplaceAll(string(a), "//{ReplayWhereBusinessId}", whereMapBusinessId)
			result += datestr + "\n"
		} else if strings.Contains(string(a), "//{ReplaySaveBusinessId}") { //保存添加business_id
			datestr := strings.ReplaceAll(string(a), "//{ReplaySaveBusinessId}", paramSaveBusinessId)
			result += datestr + "\n"
		} else if strings.Contains(string(a), "{fields}") {
			datestr := strings.ReplaceAll(string(a), "{fields}", fields)
			result += datestr + "\n"
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

// 生成代码模板
func MarkeGoCodeTpl(file_path, filename, packageName string) {
	// 创建go文件
	filePath := filepath.Join(file_path, filename+".go")
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			os.Create(filePath)
		}
	}
	//复制go文件模板到新创建文件
	err := CopyFileContents(filepath.Join("devsource/developer/codetpl/go/gotpl.gos"), filePath)
	if err != nil {
		panic(err)
	}
	//打开新键go文件内容-并替换
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
		if strings.Contains(string(a), "Replace") {
			datestr := strings.ReplaceAll(string(a), "Replace", gstr.UcFirst(filename))
			result += datestr + "\n"
		} else if strings.Contains(string(a), "packageName") {
			datestr := strings.ReplaceAll(string(a), "packageName", packageName)
			result += datestr + "\n"
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

/**************************前端处理**********************************/
// 1修改api.ts
//packageName=包名，filename文件名
func ApitsReplay(filePath, packageName, filename string) {
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
		if strings.Contains(string(a), "modname/filename") {
			datestr := ""
			if filename == "index" {
				datestr = strings.ReplaceAll(string(a), "modname/filename", packageName)
			} else {
				datestr = strings.ReplaceAll(string(a), "modname/filename", fmt.Sprintf("%s/%s", packageName, filename))
			}
			result += datestr + "\n"
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

// 生成模板处理api.ts
func ApitsReplayTpl(filePath, packageName, filename string) {
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
		if strings.Contains(string(a), "/modname/filename") {
			datestr := ""
			if filename == "index" {
				datestr = strings.ReplaceAll(string(a), "/modname/filename", packageName)
			} else {
				datestr = strings.ReplaceAll(string(a), "/modname/filename", fmt.Sprintf("%s/%s", packageName, filename))
			}
			result += datestr + "\n"
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

// 1.1、修改data.ts字段
// file_path文件路径，tablefieldlist 字段列表
func UpFieldData(file_path string, tablefieldlist gf.OrmResult) {
	f, err := os.Open(file_path)
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
		if strings.Contains(string(a), "{},") {
			relaystr := ""
			for _, webjson := range tablefieldlist {
				width := ""
				if webjson["width"].Int() > 0 {
					width = fmt.Sprintf("       width: %v,\n", webjson["width"])
				}
				slotName := webjson["show_ui"].String()
				field_str := webjson["field"].String()
				type_str := webjson["formtype"].String()
				if type_str == "belongto" {
					relaystr += fmt.Sprintf("     {\n       title:  \"%v\",\n       dataIndex: \"%vName\",\n       align:\"%v\",\n%v     },\n", webjson["name"], field_str, webjson["align"], width)
				} else if !webjson["show_ui"].IsEmpty() {
					if slotName == "dotstatus" || slotName == "tag" {
						relaystr += fmt.Sprintf("     {\n       title:  \"%v\",\n       dataIndex: \"%v\",\n       slotName:  \"%v\",\n       align:\"%v\",\n%v     },\n", webjson["name"], field_str, field_str, webjson["align"], width)
					} else {
						relaystr += fmt.Sprintf("     {\n       title:  \"%v\",\n       dataIndex: \"%v\",\n       slotName:  \"%v\",\n       align:\"%v\",\n%v     },\n", webjson["name"], field_str, slotName, webjson["align"], width)
					}
				} else {
					relaystr += fmt.Sprintf("     {\n       title:  \"%v\",\n       dataIndex: \"%v\",\n       align:\"%v\",\n%v     },\n", webjson["name"], field_str, webjson["align"], width)
				}
			}
			datestr := strings.ReplaceAll(string(a), "{},", relaystr)
			result += datestr
		} else {
			result += string(a) + "\n"
		}
	}
	fw, err := os.OpenFile(file_path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //os.O_TRUNC清空文件重新写入，否则原文件内容可能残留
	w := bufio.NewWriter(fw)
	w.WriteString(result)
	if err != nil {
		panic(err)
	}
	w.Flush()
}

// 2.1、修改AddForm.vue字段
// file_path文件路径，tablefieldname 字段
func UpFieldAddForm(file_path string, fields interface{}, tablefieldlist gf.OrmResult) {
	f, err := os.Open(file_path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	var result = ""
	//处理数据
	FieldData := ""   //数据字段初始
	relayhtml := ""   //HTML模板
	replaceFile := "" //替换附件字段
	EditList := ""    //替换富文本编辑器
	EditCount := 0    //替换富文本编辑器个数
	Edithtml := ""    //编辑器HTML模板
	for _, webjson := range tablefieldlist {
		option_value := webjson["option_value"].String()
		value_str := webjson["field"].String()
		label_str := webjson["name"].String()
		type_str := webjson["formtype"].String()
		isrequired := webjson["required"].Int()
		if strings.Contains(value_str, "file") {
			replaceFile = value_str
		}
		if value_str != "id" {
			var defval interface{} = "\"\""
			if type_str == "number" {
				defval = "null"
			} else if type_str == "checkbox" {
				defval = "[]"
			}
			if !webjson["def_value"].IsEmpty() {
				_, err := strconv.ParseFloat(webjson["def_value"].String(), 64)
				if err == nil {
					defval = webjson["def_value"]
				} else {
					defval = fmt.Sprintf("\"%v\"", webjson["def_value"])
				}
			}
			FieldData += fmt.Sprintf("            %v:%v,\n", value_str, defval)
		}
		//处理html模版
		if value_str != "id" && value_str != "createtime" && value_str != "updatetime" {
			if type_str == "switch" { //单选框
				rule_str := ""
				if isrequired == 1 {
					rule_str = fmt.Sprintf(":rules=\"%v\"", "[{required:true,message:'请选择"+label_str+"'}]")
				}
				//组装选项
				option_list := ""
				var yes_val interface{} = true
				var no_val interface{} = false
				if option_value != "" {
					option_arr := strings.Split(option_value, `,`)
					if len(option_arr) == 2 {
						one_arr := strings.Split(option_arr[0], `=`)
						yes_val = one_arr[0]
						option_list += "\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t<<template #checked>" + one_arr[1] + "</template>"
						tow_arr := strings.Split(option_arr[1], `=`)
						no_val = tow_arr[0]
						option_list += "\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t<<template #checked>" + tow_arr[1] + "</template>"
					}
				}
				relayhtml += fmt.Sprintf("\t\t\t\t\t\t\t\t\t\t<a-col :span=\"%v\">\n\t\t\t\t\t\t\t\t\t\t\t<a-form-item field=\"%v\" label=\"%v\" %v >\n\t\t\t\t\t\t\t\t\t\t\t\t\t<a-switch v-model=\"formData.%v\" checked-value=\"%v\" unchecked-value=\"%v\">%v\n\t\t\t\t\t\t\t\t\t\t\t\t\t</a-switch>\n\t\t\t\t\t\t\t\t\t\t\t</a-form-item>\n\t\t\t\t\t\t\t\t\t\t</a-col>\n", webjson["gridwidth"], value_str, label_str, rule_str, value_str, option_list, yes_val, no_val)
			} else if type_str == "radio" { //单选框
				rule_str := ""
				if isrequired == 1 {
					rule_str = fmt.Sprintf(":rules=\"%v\"", "[{required:true,message:'请选择"+label_str+"'}]")
				}
				//组装选项
				option_list := ""
				if option_value != "" {
					option_arr := strings.Split(option_value, `,`)
					for _, option_str := range option_arr {
						val_arr := strings.Split(option_str, `=`)
						_, err := strconv.ParseFloat(val_arr[0], 64)
						if err == nil {
							option_list += "\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t<a-radio :value=\"" + val_arr[0] + "\" >" + val_arr[1] + "</a-radio>"
						} else {
							option_list += "\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t<a-radio value=\"" + val_arr[0] + "\" >" + val_arr[1] + "</a-radio>"
						}
					}
				}
				relayhtml += fmt.Sprintf("\t\t\t\t\t\t\t\t\t\t<a-col :span=\"%v\">\n\t\t\t\t\t\t\t\t\t\t\t<a-form-item field=\"%v\" label=\"%v\" %v >\n\t\t\t\t\t\t\t\t\t\t\t\t\t<a-radio-group v-model=\"formData.%v\" >%v\n\t\t\t\t\t\t\t\t\t\t\t\t\t</a-radio-group>\n\t\t\t\t\t\t\t\t\t\t\t</a-form-item>\n\t\t\t\t\t\t\t\t\t\t</a-col>\n", webjson["gridwidth"], value_str, label_str, rule_str, value_str, option_list)
			} else if type_str == "select" { //下拉框
				rule_str := ""
				if isrequired == 1 {
					rule_str = fmt.Sprintf(":rules=\"%v\"", "[{required:true,message:'请选择"+label_str+"'}]")
				}
				//组装选项
				option_list := ""
				if option_value != "" {
					option_arr := strings.Split(option_value, `,`)
					for _, option_str := range option_arr {
						val_arr := strings.Split(option_str, `=`)
						_, err := strconv.ParseFloat(val_arr[0], 64)
						if err == nil {
							option_list += "\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t<a-option :value=\"" + val_arr[0] + "\" >" + val_arr[1] + "</a-option>"
						} else {
							option_list += "\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t<a-option value=\"" + val_arr[0] + "\" >" + val_arr[1] + "</a-option>"
						}
					}
				}
				relayhtml += fmt.Sprintf("\t\t\t\t\t\t\t\t\t\t<a-col :span=\"%v\">\n\t\t\t\t\t\t\t\t\t\t\t<a-form-item field=\"%v\" label=\"%v\" %v >\n\t\t\t\t\t\t\t\t\t\t\t\t\t<a-select  v-model=\"formData.%v\" >%v\n\t\t\t\t\t\t\t\t\t\t\t\t\t</a-select>\n\t\t\t\t\t\t\t\t\t\t\t</a-form-item>\n\t\t\t\t\t\t\t\t\t\t</a-col>\n", webjson["gridwidth"], value_str, label_str, rule_str, value_str, option_list)
			} else if type_str == "checkbox" { //复选框
				rule_str := ""
				if isrequired == 1 {
					rule_str = fmt.Sprintf(":rules=\"%v\"", "[{required:true,message:'请选择"+label_str+"'}]")
				}
				//组装选项
				option_list := ""
				if option_value != "" {
					option_arr := strings.Split(option_value, `,`)
					for _, option_str := range option_arr {
						val_arr := strings.Split(option_str, `=`)
						_, err := strconv.ParseFloat(val_arr[0], 64)
						if err == nil {
							option_list += "\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t<a-checkbox :value=\"" + val_arr[0] + "\" >" + val_arr[1] + "</a-checkbox>"
						} else {
							option_list += "\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t<a-checkbox value=\"" + val_arr[0] + "\" >" + val_arr[1] + "</a-checkbox>"
						}
					}
				}
				relayhtml += fmt.Sprintf("\t\t\t\t\t\t\t\t\t\t<a-col :span=\"%v\">\n\t\t\t\t\t\t\t\t\t\t\t<a-form-item field=\"%v\" label=\"%v\" %v >\n\t\t\t\t\t\t\t\t\t\t\t\t\t<a-checkbox-group v-model=\"formData.%v\" >%v\n\t\t\t\t\t\t\t\t\t\t\t\t\t</a-checkbox-group>\n\t\t\t\t\t\t\t\t\t\t\t</a-form-item>\n\t\t\t\t\t\t\t\t\t\t</a-col>\n", webjson["gridwidth"], value_str, label_str, rule_str, value_str, option_list)
			} else if type_str == "belongto" { //关联字段
				rule_str := ""
				if isrequired == 1 {
					rule_str = fmt.Sprintf(":rules=\"%v\"", "[{required:true,message:'请选择"+label_str+"'}]")
				}
				relayhtml += fmt.Sprintf("\t\t\t\t\t\t\t\t\t\t<a-col :span=\"%v\">\n\t\t\t\t\t\t\t\t\t\t\t<a-form-item field=\"%v\" label=\"%v\" %v >\n\t\t\t\t\t\t\t\t\t\t\t\t\t<FormBelongTable v-model=\"formData.%v\" placeholder=\"请选%v\" tablename=\"%v\" showfield=\"%v\"/>\n\t\t\t\t\t\t\t\t\t\t\t</a-form-item>\n\t\t\t\t\t\t\t\t\t\t</a-col>\n", webjson["gridwidth"], value_str, label_str, rule_str, value_str, label_str, webjson["datatable"], webjson["datatablename"])
			} else if type_str == "belongDic" { //关联字典
				rule_str := ""
				if isrequired == 1 {
					rule_str = fmt.Sprintf(":rules=\"%v\"", "[{required:true,message:'请选择"+label_str+"'}]")
				}
				relayhtml += fmt.Sprintf("\t\t\t\t\t\t\t\t\t\t<a-col :span=\"%v\">\n\t\t\t\t\t\t\t\t\t\t\t<a-form-item field=\"%v\" label=\"%v\" %v >\n\t\t\t\t\t\t\t\t\t\t\t\t\t<FormDicSelect v-model=\"formData.%v\" placeholder=\"请选%v\" dicgroupid=\"%v\"/>\n\t\t\t\t\t\t\t\t\t\t\t</a-form-item>\n\t\t\t\t\t\t\t\t\t\t</a-col>\n", webjson["gridwidth"], value_str, label_str, rule_str, value_str, label_str, webjson["dic_group_id"])
			} else if type_str == "image" {
				rule_str := ""
				if isrequired == 1 {
					rule_str = fmt.Sprintf(":rules=\"%v\"", "[{required:true,message:'请选择"+label_str+"'}]")
				}
				relayhtml += fmt.Sprintf("\t\t\t\t\t\t\t\t\t\t<a-col :span=\"%v\">\n\t\t\t\t\t\t\t\t\t\t\t<a-form-item field=\"%v\" label=\"%v\" %v >\n\t\t\t\t\t\t\t\t\t\t\t\t\t<FormImageBox v-model=\"formData.%v\" placeholder=\"请选%v\"/>\n\t\t\t\t\t\t\t\t\t\t\t</a-form-item>\n\t\t\t\t\t\t\t\t\t\t</a-col>\n", webjson["gridwidth"], value_str, label_str, rule_str, value_str, label_str)
			} else if type_str == "images" {
				rule_str := ""
				if isrequired == 1 {
					rule_str = fmt.Sprintf(":rules=\"%v\"", "[{required:true,message:'请选择"+label_str+"'}]")
				}
				relayhtml += fmt.Sprintf("\t\t\t\t\t\t\t\t\t\t<a-col :span=\"%v\">\n\t\t\t\t\t\t\t\t\t\t\t<a-form-item field=\"%v\" label=\"%v\" %v >\n\t\t\t\t\t\t\t\t\t\t\t\t\t<FormImagesBox v-model=\"formData.%v\" placeholder=\"请选%v\"/>\n\t\t\t\t\t\t\t\t\t\t\t</a-form-item>\n\t\t\t\t\t\t\t\t\t\t</a-col>\n", webjson["gridwidth"], value_str, label_str, rule_str, value_str, label_str)
			} else if type_str == "audio" {
				rule_str := ""
				if isrequired == 1 {
					rule_str = fmt.Sprintf(":rules=\"%v\"", "[{required:true,message:'请选择"+label_str+"'}]")
				}
				relayhtml += fmt.Sprintf("\t\t\t\t\t\t\t\t\t\t<a-col :span=\"%v\">\n\t\t\t\t\t\t\t\t\t\t\t<a-form-item field=\"%v\" label=\"%v\" %v >\n\t\t\t\t\t\t\t\t\t\t\t\t\t<FormAudioBox v-model=\"formData.%v\" placeholder=\"请选%v\"/>\n\t\t\t\t\t\t\t\t\t\t\t</a-form-item>\n\t\t\t\t\t\t\t\t\t\t</a-col>\n", webjson["gridwidth"], value_str, label_str, rule_str, value_str, label_str)
			} else if type_str == "file" {
				rule_str := ""
				if isrequired == 1 {
					rule_str = fmt.Sprintf(":rules=\"%v\"", "[{required:true,message:'请选择"+label_str+"'}]")
				}
				relayhtml += fmt.Sprintf("\t\t\t\t\t\t\t\t\t\t<a-col :span=\"%v\">\n\t\t\t\t\t\t\t\t\t\t\t<a-form-item field=\"%v\" label=\"%v\" %v >\n\t\t\t\t\t\t\t\t\t\t\t\t\t<FormFileBox v-model=\"formData.%v\" placeholder=\"请选%v\"/>\n\t\t\t\t\t\t\t\t\t\t\t</a-form-item>\n\t\t\t\t\t\t\t\t\t\t</a-col>\n", webjson["gridwidth"], value_str, label_str, rule_str, value_str, label_str)
			} else if type_str == "number" { //数字本文
				rule_str := ""
				if isrequired == 1 {
					rule_str = fmt.Sprintf(":rules=\"%v\"", "[{required:true,message:'请填写"+label_str+"'}]")
				}
				relayhtml += fmt.Sprintf("\t\t\t\t\t\t\t\t\t\t<a-col :span=\"%v\">\n\t\t\t\t\t\t\t\t\t\t\t<a-form-item field=\"%v\" label=\"%v\" %v >\n\t\t\t\t\t\t\t\t\t\t\t\t\t<a-input-number v-model=\"formData.%v\" placeholder=\"输入%v\" />\n\t\t\t\t\t\t\t\t\t\t\t</a-form-item>\n\t\t\t\t\t\t\t\t\t\t</a-col>\n", webjson["gridwidth"], value_str, label_str, rule_str, value_str, label_str)
			} else if type_str == "time" { //时间选择器
				rule_str := ""
				if isrequired == 1 {
					rule_str = fmt.Sprintf(":rules=\"%v\"", "[{required:true,message:'请选择"+label_str+"'}]")
				}
				relayhtml += fmt.Sprintf("\t\t\t\t\t\t\t\t\t\t<a-col :span=\"%v\">\n\t\t\t\t\t\t\t\t\t\t\t<a-form-item field=\"%v\" label=\"%v\" %v >\n\t\t\t\t\t\t\t\t\t\t\t\t\t<a-time-picker  format=\"HH:mm\" v-model=\"formData.%v\" placeholder=\"请选%v\" />\n\t\t\t\t\t\t\t\t\t\t\t</a-form-item>\n\t\t\t\t\t\t\t\t\t\t</a-col>\n", webjson["gridwidth"], value_str, label_str, rule_str, value_str, label_str)
			} else if type_str == "date" { //日期控件
				rule_str := ""
				if isrequired == 1 {
					rule_str = fmt.Sprintf(":rules=\"%v\"", "[{required:true,message:'请选择"+label_str+"'}]")
				}
				relayhtml += fmt.Sprintf("\t\t\t\t\t\t\t\t\t\t<a-col :span=\"%v\">\n\t\t\t\t\t\t\t\t\t\t\t<a-form-item field=\"%v\" label=\"%v\" %v >\n\t\t\t\t\t\t\t\t\t\t\t\t\t<a-date-picker v-model=\"formData.%v\" placeholder=\"请选%v\" />\n\t\t\t\t\t\t\t\t\t\t\t</a-form-item>\n\t\t\t\t\t\t\t\t\t\t</a-col>\n", webjson["gridwidth"], value_str, label_str, rule_str, value_str, label_str)
			} else if type_str == "datetime" { //日期时间控件
				rule_str := ""
				if isrequired == 1 {
					rule_str = fmt.Sprintf(":rules=\"%v\"", "[{required:true,message:'请选择"+label_str+"'}]")
				}
				relayhtml += fmt.Sprintf("\t\t\t\t\t\t\t\t\t\t<a-col :span=\"%v\">\n\t\t\t\t\t\t\t\t\t\t\t<a-form-item field=\"%v\" label=\"%v\" %v >\n\t\t\t\t\t\t\t\t\t\t\t\t\t<a-date-picker show-time  format=\"YYYY-MM-DD HH:mm\" v-model=\"formData.%v\" placeholder=\"请选%v\" />\n\t\t\t\t\t\t\t\t\t\t\t</a-form-item>\n\t\t\t\t\t\t\t\t\t\t</a-col>\n", webjson["gridwidth"], value_str, label_str, rule_str, value_str, label_str)
			} else if type_str == "colorpicker" { //颜色选择器
				rule_str := ""
				if isrequired == 1 {
					rule_str = fmt.Sprintf(":rules=\"%v\"", "[{required:true,message:'请选择"+label_str+"'}]")
				}
				relayhtml += fmt.Sprintf("\t\t\t\t\t\t\t\t\t\t<a-col :span=\"%v\">\n\t\t\t\t\t\t\t\t\t\t\t<a-form-item field=\"%v\" label=\"%v\" %v >\n\t\t\t\t\t\t\t\t\t\t\t\t\t<a-color-picker showPreset v-model=\"formData.%v\"/>\n\t\t\t\t\t\t\t\t\t\t\t</a-form-item>\n\t\t\t\t\t\t\t\t\t\t</a-col>\n", webjson["gridwidth"], value_str, label_str, rule_str, value_str)
			} else if type_str == "text" { //文本输入框
				rule_str := ""
				if isrequired == 1 {
					rule_str = fmt.Sprintf(":rules=\"%v\"", "[{required:true,message:'请填写"+label_str+"'}]")
				}
				relayhtml += fmt.Sprintf("\t\t\t\t\t\t\t\t\t\t<a-col :span=\"%v\">\n\t\t\t\t\t\t\t\t\t\t\t<a-form-item field=\"%v\" label=\"%v\" %v >\n\t\t\t\t\t\t\t\t\t\t\t\t\t<a-input v-model=\"formData.%v\" placeholder=\"输入%v\" />\n\t\t\t\t\t\t\t\t\t\t\t</a-form-item>\n\t\t\t\t\t\t\t\t\t\t</a-col>\n", webjson["gridwidth"], value_str, label_str, rule_str, value_str, label_str)
			} else if type_str == "textarea" {
				rule_str := ""
				if isrequired == 1 {
					rule_str = fmt.Sprintf(":rules=\"%v\"", "[{required:true,message:'请填写"+label_str+"'}]")
				}
				relayhtml += fmt.Sprintf("\t\t\t\t\t\t\t\t\t\t<a-col :span=\"%v\">\n\t\t\t\t\t\t\t\t\t\t\t<a-form-item field=\"%v\" label=\"%v\" %v >\n\t\t\t\t\t\t\t\t\t\t\t\t\t<a-textarea v-model=\"formData.%v\" placeholder=\"输入%v\" :auto-size=\"{minRows:3,maxRows:5}\"/>\n\t\t\t\t\t\t\t\t\t\t\t</a-form-item>\n\t\t\t\t\t\t\t\t\t\t</a-col>\n", webjson["gridwidth"], value_str, label_str, rule_str, value_str, label_str)
			} else if type_str == "editor" {
				EditCount++
				EditList += fmt.Sprintf("\n\t\t\t\t\t{id:%v,name:\"%v\"},", EditCount+1, label_str)
				//编辑器
				Edithtml += fmt.Sprintf("\n\t\t\t\t\t\t\t<FormEditorBox v-model=\"formData.%v\" placeholder=\"请编辑%v\" :winHeights=\"windHeight\" :activeKey=\"activeKey\" :subnum=\"%v\" />", value_str, label_str, EditCount+1)
			}
		}
	}
	for {
		a, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		if strings.Contains(string(a), "isEditor=ref(false)") && EditCount != 0 {
			datestr := strings.ReplaceAll(string(a), "isEditor=ref(false)", "isEditor=ref(true)")
			result += datestr + "\n"
		} else if strings.Contains(string(a), "replaceField:null") {
			datestr := strings.ReplaceAll(string(a), "replaceField:null", FieldData)
			result += datestr + "\n"
		} else if strings.Contains(string(a), "{replacEdit:null}") { //添加富文本列表
			datestr := strings.ReplaceAll(string(a), "{replacEdit:null}", EditList)
			result += datestr + "\n"
		} else if strings.Contains(string(a), "<!--replaceEditHtml-->  ") { //添加富文本编辑器
			datestr := strings.ReplaceAll(string(a), "<!--replaceEditHtml-->  ", Edithtml)
			result += datestr + "\n"
		} else if strings.Contains(string(a), "['replaceFile']") && replaceFile != "" {
			datestr := strings.ReplaceAll(string(a), "['replaceFile']", "."+replaceFile)
			result += datestr + "\n"
		} else if strings.Contains(string(a), "<!--replaceTpl-->") {
			datestr := strings.ReplaceAll(string(a), "<!--replaceTpl-->", relayhtml)
			result += datestr
		} else {
			result += string(a) + "\n"
		}
	}
	fw, err := os.OpenFile(file_path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //os.O_TRUNC清空文件重新写入，否则原文件内容可能残留
	w := bufio.NewWriter(fw)
	w.WriteString(result)
	if err != nil {
		panic(err)
	}
	w.Flush()
}

// 2.2、修改index.vue页面内容
// file_path文件路径，tablefieldname 字段
func UpFieldIndex(file_path string, tablefieldlist, listfield gf.OrmResult, tpl_type, filename string) {
	f, err := os.Open(file_path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	var result = ""
	//处理搜索数据
	FieldData := "" //数据字段初始
	relayhtml := "" //HTML模板
	nohasecid := true
	for _, webjson := range tablefieldlist {
		option_value := webjson["option_value"].String()
		field_str := webjson["field"].String()
		label_str := webjson["name"].String()
		type_str := webjson["searchtype"].String()
		if field_str == "cid" {
			nohasecid = false
		}
		var defval any = "\"\""
		if type_str == "selectmore" || type_str == "daterange" {
			defval = "[]"
		} else if field_str == "cid" {
			defval = 0
		}
		FieldData += fmt.Sprintf("\t\t\t%v:%v,\n", field_str, defval)
		//处理html模版
		if type_str == "text" { //本文
			relayhtml += fmt.Sprintf("\t\t\t\t\t\t\t\t<a-input style=\"width: %vpx;\" v-model=\"formModel.%v\" placeholder=\"输入%v\" allow-clear />\n", webjson["searchwidth"], field_str, label_str)
		} else if type_str == "number" { //数字本文
			relayhtml += fmt.Sprintf("\t\t\t\t\t\t\t\t<a-input-number style=\"width: %vpx;\" v-model=\"formModel.%v\" placeholder=\"输入%v\" allow-clear />\n", webjson["searchwidth"], field_str, label_str)
		} else if type_str == "dic" { //字典数据
			relayhtml += fmt.Sprintf("\t\t\t\t\t\t\t\t<FormDicSelect style=\"width: %vpx;\" v-model=\"formModel.%v\" placeholder=\"输入%v\" dicgroupid=\"%v\"/>\n", webjson["searchwidth"], field_str, label_str, webjson["dic_group_id"])
		} else if type_str == "belongto" { //关联表数据
			relayhtml += fmt.Sprintf("\t\t\t\t\t\t\t\t<FormBelongTable from=\"search\" style=\"width: %vpx;\" v-model=\"formModel.%v\" placeholder=\"输入%v\" tablename=\"%v\" showfield=\"%v\"/>\n", webjson["searchwidth"], field_str, label_str, webjson["datatable"], webjson["datatablename"])
		} else if type_str == "select" { //下拉框
			//组装选项
			option_list := ""
			if option_value != "" {
				option_arr := strings.Split(option_value, `,`)
				for _, option_str := range option_arr {
					val_arr := strings.Split(option_str, `=`)
					_, err := strconv.ParseFloat(val_arr[0], 64)
					if err == nil {
						option_list += "\n\t\t\t\t\t\t\t\t\t<a-option :value=\"" + val_arr[0] + "\" >" + val_arr[1] + "</a-option>"
					} else {
						option_list += "\n\t\t\t\t\t\t\t\t\t<a-option value=\"" + val_arr[0] + "\" >" + val_arr[1] + "</a-option>"
					}
				}
			}
			relayhtml += fmt.Sprintf("\t\t\t\t\t\t\t\t<a-select style=\"width: %vpx;\" v-model=\"formModel.%v\" placeholder=\"选择%v\" allow-clear>%v\n\t\t\t\t\t\t\t\t</a-select>\n", webjson["searchwidth"], field_str, label_str, option_list)
		} else if type_str == "selectmore" { //多选下拉框
			//组装选项
			option_list := ""
			if option_value != "" {
				option_arr := strings.Split(option_value, `,`)
				for _, option_str := range option_arr {
					val_arr := strings.Split(option_str, `=`)
					_, err := strconv.ParseFloat(val_arr[0], 64)
					if err == nil {
						option_list += "\n\t\t\t\t\t\t\t\t\t<a-option :value=\"" + val_arr[0] + "\" >" + val_arr[1] + "</a-option>"
					} else {
						option_list += "\n\t\t\t\t\t\t\t\t\t<a-option value=\"" + val_arr[0] + "\" >" + val_arr[1] + "</a-option>"
					}
				}
			}
			relayhtml += fmt.Sprintf("\t\t\t\t\t\t\t\t<a-select style=\"width: %vpx;\" v-model=\"formModel.%v\" placeholder=\"选择%v\" allow-clear multiple>%v\n\t\t\t\t\t\t\t\t</a-select>\n", webjson["searchwidth"], field_str, label_str, option_list)
		} else if type_str == "date" { //日期选择器
			relayhtml += fmt.Sprintf("\t\t\t\t\t\t\t\t<a-date-picker style=\"width: %vpx;\" v-model=\"formModel.%v\" placeholder=\"选择%v\" allow-clear/>\n", webjson["searchwidth"], field_str, label_str)
		} else if type_str == "daterange" { //日期区间选择器
			relayhtml += fmt.Sprintf("\t\t\t\t\t\t\t\t<span>%v：<a-range-picker style=\"width: %vpx;\" v-model=\"formModel.%v\" allow-clear :shortcuts=\"shortcuts\" shortcuts-position=\"left\" @change=\"fetchData\"/></span>\n", label_str, webjson["searchwidth"], field_str)
		}
	}
	//处理table列表显示组件
	tableUIHtml := "" //HTML模板
	var tableUIHtml_Hase []string
	stableh := garray.NewStrArray()
	for _, listdata := range listfield {
		slotName := listdata["show_ui"].String()
		field_str := listdata["field"].String()
		option_value := listdata["option_value"].String()
		dict := "[]"
		if option_value != "" {
			option_arr := strings.Split(option_value, `,`)
			option_list := ""
			option_color := gf.SliceStr{"primary", "success", "warning", "danger", "info"}
			for index, option_str := range option_arr {
				val_arr := strings.Split(option_str, `=`)
				colorval := "info"
				if index < 5 {
					colorval = option_color[index]
				}
				//判断值类型
				var defval interface{} = "''"
				_, err := strconv.ParseFloat(val_arr[0], 64)
				if err == nil {
					defval = val_arr[0]
				} else {
					defval = fmt.Sprintf("'%v'", val_arr[0])
				}
				if index == 0 {
					option_list += fmt.Sprintf("{value:%v,label:'%v',color:'%v'}", defval, val_arr[1], colorval)
				} else {
					option_list += fmt.Sprintf(",{value:%v,label:'%v',color:'%v'}", defval, val_arr[1], colorval)
				}
			}
			dict = "[" + option_list + "]"
		}

		if slotName != "" {
			stableh.SetArray(tableUIHtml_Hase)
			if slotName == "des" {
				if !stableh.Contains("des") {
					tableUIHtml_Hase = append(tableUIHtml_Hase, "des")
					tableUIHtml += fmt.Sprintf("\t\t\t\t\t\t<template #%v=\"{record,column}\">\n\t\t\t\t\t\t\t<a-typography-paragraph :ellipsis=\"{rows: 2,showTooltip:true}\">{{ record[column.dataIndex] }}</a-typography-paragraph>\n\t\t\t\t\t\t</template>\n", slotName)
				}
			} else if slotName == "cellcopy" {
				if !stableh.Contains("cellcopy") {
					tableUIHtml_Hase = append(tableUIHtml_Hase, "cellcopy")
					tableUIHtml += fmt.Sprintf("\t\t\t\t\t\t<template #%v=\"{record,column}\">\n\t\t\t\t\t\t\t<cell-copy :data=\"record[column.dataIndex]\"/>\n\t\t\t\t\t\t</template>\n", slotName)
				}
			} else if slotName == "dic" { //字典数据显示
				if !stableh.Contains("dic") {
					tableUIHtml_Hase = append(tableUIHtml_Hase, "dic")
					tableUIHtml += fmt.Sprintf("\t\t\t\t\t\t<template #%v=\"{record,column}\">\n\t\t\t\t\t\t\t <cell-tag-plus v-if=\"record[column.dataIndex]\" :color=\"record[column.dataIndex].tagcolor\" :value=\"record[column.dataIndex].keyname\"/>\n\t\t\t\t\t\t</template>\n", slotName)
				}
			} else if slotName == "image" {
				if !stableh.Contains("image") {
					tableUIHtml_Hase = append(tableUIHtml_Hase, "image")
					tableUIHtml += fmt.Sprintf("\t\t\t\t\t\t<template #%v=\"{record,column}\">\n\t\t\t\t\t\t\t<cell-image height=\"32\" width=\"32\" :src=\"record[column.dataIndex]\"/>\n\t\t\t\t\t\t</template>\n", slotName)
				}
			} else if slotName == "images" {
				if !stableh.Contains("images") {
					tableUIHtml_Hase = append(tableUIHtml_Hase, "images")
					tableUIHtml += fmt.Sprintf("\t\t\t\t\t\t<template #%v=\"{record,column}\">\n\t\t\t\t\t\t\t<cell-images :data=\"record[column.dataIndex]\"/>\n\t\t\t\t\t\t</template>\n", slotName)
				}
			} else if slotName == "tag" {
				tableUIHtml += fmt.Sprintf("\t\t\t\t\t\t<template #%v=\"{record,column}\">\n\t\t\t\t\t\t\t<cell-tag :value=\"record[column.dataIndex]\" :dict=\"%v\"/>\n\t\t\t\t\t\t</template>\n", field_str, dict)
			} else if slotName == "tags" {
				if !stableh.Contains("tags") {
					tableUIHtml_Hase = append(tableUIHtml_Hase, "tags")
					tableUIHtml += fmt.Sprintf("\t\t\t\t\t\t<template #%v=\"{record,column}\">\n\t\t\t\t\t\t\t<cell-tags :data=\"record[column.dataIndex]\"/>\n\t\t\t\t\t\t</template>\n", slotName)
				}
			} else if slotName == "cellstatus" {
				if !stableh.Contains("cellstatus") {
					tableUIHtml_Hase = append(tableUIHtml_Hase, "cellstatus")
					tableUIHtml += fmt.Sprintf("\t\t\t\t\t\t<template #%v=\"{record,column}\">\n\t\t\t\t\t\t\t<cell-status :status=\"record[column.dataIndex]\"/>\n\t\t\t\t\t\t</template>\n", slotName)
				}
			} else if slotName == "switchstatus" {
				if !stableh.Contains("switchstatus") {
					tableUIHtml_Hase = append(tableUIHtml_Hase, "switchstatus")
					tableUIHtml += fmt.Sprintf("\t\t\t\t\t\t<template #%v=\"{record,column}\">\n\t\t\t\t\t\t\t<a-switch v-model=\"record[column.dataIndex]\" :checked-value=\"0\" :unchecked-value=\"1\" @change=\"handleStatus(record,column.dataIndex)\"></a-switch>\n\t\t\t\t\t\t</template>\n", slotName)
				}
			} else if slotName == "dotstatus" {
				tableUIHtml += fmt.Sprintf("\t\t\t\t\t\t<template #%v=\"{record,column}\">\n\t\t\t\t\t\t\t<dot-status :value=\"record[column.dataIndex]\" :dict=\"%v\"/>\n\t\t\t\t\t\t</template>\n", field_str, dict)
			} else if slotName == "link" {
				if !stableh.Contains("link") {
					tableUIHtml_Hase = append(tableUIHtml_Hase, "link")
					tableUIHtml += fmt.Sprintf("\t\t\t\t\t\t<template #%v=\"{record,column}\">\n\t\t\t\t\t\t\t<a-link :href=\"record[column.dataIndex]\" class=\"nowraphind\" target=\"_blank\" :title=\"record[column.dataIndex]\">{{ record[column.dataIndex] }}</a-link>\n\t\t\t\t\t\t</template>\n", slotName)
				}
			} else if slotName == "gender" {
				if !stableh.Contains("gender") {
					tableUIHtml_Hase = append(tableUIHtml_Hase, "gender")
					tableUIHtml += fmt.Sprintf("\t\t\t\t\t\t<template #%v=\"{record,column}\">\n\t\t\t\t\t\t\t<cell-gender :gender=\"record[column.dataIndex]\"/>\n\t\t\t\t\t\t</template>\n", slotName)
				}
			} else if slotName == "avatar" {
				if !stableh.Contains("avatar") {
					tableUIHtml_Hase = append(tableUIHtml_Hase, "avatar")
					tableUIHtml += fmt.Sprintf("\t\t\t\t\t\t<template #%v=\"{record,column}\">\n\t\t\t\t\t\t\t<cell-node-avatar :avatar=\"record.avatar?record.avatar:record.image\" :size=\"20\" :name=\"record[column.dataIndex]\"/>\n\t\t\t\t\t\t</template>\n", slotName)
				}
			} else if slotName == "nodeavatar" {
				if !stableh.Contains("nodeavatar") {
					tableUIHtml_Hase = append(tableUIHtml_Hase, "nodeavatar")
					tableUIHtml += fmt.Sprintf("\t\t\t\t\t\t<template #%v=\"{record,column}\">\n\t\t\t\t\t\t\t<cell-node-avatar :avatar=\"record.avatar?record.avatar:record.image\" :size=\"20\" :name=\"record[column.dataIndex]\"/>\n\t\t\t\t\t\t</template>\n", slotName)
				}
			} else if slotName == "date" {
				if !stableh.Contains("date") {
					tableUIHtml_Hase = append(tableUIHtml_Hase, "date")
					tableUIHtml += fmt.Sprintf("\t\t\t\t\t\t<template #%v=\"{record,column}\">\n\t\t\t\t\t\t\t{{dayjs(record[column.dataIndex]).format(\"YYYY-MM-DD\")}}\n\t\t\t\t\t\t</template>\n", slotName)
				}
			} else if slotName == "datetime" {
				if !stableh.Contains("datetime") {
					tableUIHtml_Hase = append(tableUIHtml_Hase, "datetime")
					tableUIHtml += fmt.Sprintf("\t\t\t\t\t\t<template #%v=\"{record,column}\">\n\t\t\t\t\t\t\t{{dayjs(record[column.dataIndex]).format(\"YYYY-MM-DD HH:mm\")}}\n\t\t\t\t\t\t</template>\n", slotName)
				}
			} else if slotName == "color" {
				if !stableh.Contains("color") {
					tableUIHtml_Hase = append(tableUIHtml_Hase, "color")
					tableUIHtml += fmt.Sprintf("\t\t\t\t\t\t<template #%v=\"{record,column}\">\n\t\t\t\t\t\t\t<a-color-picker :defaultValue=\"record[column.dataIndex]\" :trigger-props=\"{popupVisible: false}\"/>\n\t\t\t\t\t\t</template>\n", slotName)
				}
			}
		}
	}
	//是否存在cid
	if tpl_type == "sitecatelist" && nohasecid {
		FieldData += "\t\t\tcid: 0,\n"
	}
	for {
		a, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		if strings.Contains(string(a), "//{SearchField}") {
			datestr := strings.ReplaceAll(string(a), "//{SearchField}", FieldData)
			result += datestr + "\n"
		} else if strings.Contains(string(a), "<!--SearchHtml-->") {
			datestr := strings.ReplaceAll(string(a), "<!--SearchHtml-->", relayhtml)
			result += datestr
		} else if strings.Contains(string(a), "<!--tableUIHtml-->") {
			datestr := strings.ReplaceAll(string(a), "<!--tableUIHtml-->", tableUIHtml)
			result += datestr
		} else if strings.Contains(string(a), "{vuetplname}") { //替换index.vue页面name对应类名-页签缓存使用
			datestr := strings.ReplaceAll(string(a), "{vuetplname}", filename)
			result += datestr + "\n"
		} else {
			result += string(a) + "\n"
		}
	}
	fw, err := os.OpenFile(file_path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //os.O_TRUNC清空文件重新写入，否则原文件内容可能残留
	w := bufio.NewWriter(fw)
	w.WriteString(result)
	if err != nil {
		panic(err)
	}
	w.Flush()
}

// 生成代码模板
// file_path文件路径，tablefieldname 字段
func UpFieldIndexTpl(file_path string, packageName string) {
	f, err := os.Open(file_path)
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
		if strings.Contains(string(a), "{vuetplname}") {
			datestr := strings.ReplaceAll(string(a), "{vuetplname}", packageName)
			result += datestr + "\n"
		} else {
			result += string(a) + "\n"
		}
	}
	fw, err := os.OpenFile(file_path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //os.O_TRUNC清空文件重新写入，否则原文件内容可能残留
	w := bufio.NewWriter(fw)
	w.WriteString(result)
	if err != nil {
		panic(err)
	}
	w.Flush()
	defer fw.Close()
}

/********************************************卸载前后端********************************/
// 卸载/删除文件
func UnInstallCodeFile(data gf.OrmRecord) {
	//1.删除后端代码
	//go文件目录
	file_path_go_root := filepath.Join("app/", data["api_path"].String())
	//go文件
	filego_path := filepath.Join(file_path_go_root, data["api_filename"].String())
	if _, err := os.Stat(filego_path); err == nil {
		//1.文件存在删除文件
		os.Remove(filego_path)
		if strings.Contains(data["tpl_type"].String(), "cate") {
			filename_arr := strings.Split(data["api_filename"].String(), `.`)
			filecatego_path := filepath.Join(file_path_go_root, filename_arr[0]+"cate.go")
			os.Remove(filecatego_path)
		}
		//2.删除文件夹
		dir, _ := os.ReadDir(file_path_go_root)
		if len(dir) == 0 {
			os.RemoveAll(file_path_go_root)
			//3.移除路由
			packgename_arr := strings.Split(data["api_path"].String(), `/`)
			modelname := "business" //模块名称
			if len(packgename_arr) > 0 {
				modelname = packgename_arr[0]
			}
			CheckApiRemoveController(modelname, data["api_path"].String())
		}
	}
	//2.2 删除前端下代码
	component_arr := strings.Split(data["component"].String(), `/`)
	if data["component"] != nil {
		componentpah_arr := strings.Split(data["component"].String(), (component_arr[len(component_arr)-1]))
		code_main_dir := appConf_arr["busDirName"]
		if data["codelocation"].String() == "adminDirName" {
			code_main_dir = appConf_arr["adminDirName"]
		}
		vue_path := filepath.Join(gf.String(appConf_arr["vueobjroot"]), gf.String(code_main_dir), "/src/views/", componentpah_arr[0]) //前端文件路径
		if _, err := os.Stat(vue_path); err == nil {
			os.RemoveAll(vue_path)
			//2.3.模块目录文件夹
			vue_model_path := filepath.Join(gf.String(appConf_arr["vueobjroot"]), gf.String(code_main_dir), "/src/views/", component_arr[0])
			dirs, _ := os.ReadDir(vue_model_path)
			if len(dirs) == 0 {
				os.RemoveAll(vue_model_path)
			}
		}
	}
}
