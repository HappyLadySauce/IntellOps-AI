package gf

import (
	"encoding/json"
	"fmt"
	"gofly/utils/tools/gstr"
	"gofly/utils/tools/gvar"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// 判断元素是否存在数组中
func IsContain(items []interface{}, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// 判断元素是否存在数组中-泛型val
func IsContainVal(items []*gvar.Var, item *gvar.Var) bool {
	for _, eachItem := range items {
		if eachItem.String() == item.String() {
			return true
		}
	}
	return false
}

// 判断元素是否存在数组中-字符串类型
func IsContainStr(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// 获取ip函数
func GetIp(c *GinCtx) string {
	reqIP := c.Request.Header.Get("X-Forwarded-For")
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	return reqIP
}

// 获取本地ip
func LocalIP() string {
	ip := ""
	if addrs, err := net.InterfaceAddrs(); err == nil {
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && !ipnet.IP.IsMulticast() && !ipnet.IP.IsLinkLocalUnicast() && !ipnet.IP.IsLinkLocalMulticast() && ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
			}
		}
	}
	return ip
}

/*
*
  - 1.批量获取子节点id
  - @tablename 数据表名称
    @ids 要获取的id
*/
func GetAllChilIds(tablename string, ids []*gvar.Var) []interface{} {
	var allsubids []interface{}
	for _, id := range ids {
		sub_ids := GetAllChilId(tablename, id)
		allsubids = append(allsubids, sub_ids...)
	}
	return allsubids
}

// 1.2获取所有子级ID
func GetAllChilId(tablename string, id interface{}) []interface{} {
	var subids []interface{}
	sub_ids, _ := Model(tablename).Where("pid", id).Array("id")
	if len(sub_ids) > 0 {
		for _, sid := range sub_ids {
			subids = append(subids, sid)
			subids = append(subids, GetAllChilId(tablename, sid)...)
		}
	}
	return subids
}

// 合并数组-两个数组合并为一个数组
func MergeArr(a []*gvar.Var, b []interface{}) []interface{} {
	var arr []interface{}
	for _, i := range a {
		arr = append(arr, i)
	}
	for _, j := range b {
		arr = append(arr, j)
	}
	return arr
}

// 多维数组合并-权限
func ArrayMerge(data []*gvar.Var) []interface{} {
	var rule_ids_arr []interface{}
	for _, mainv := range data {
		ids_arr := strings.Split(mainv.String(), `,`)
		for _, intv := range ids_arr {
			rule_ids_arr = append(rule_ids_arr, intv)
		}
	}
	return rule_ids_arr
}

// 把字符串打散为数组
func Axplode(data string) []interface{} {
	var rule_ids_arr []interface{}
	ids_arr := strings.Split(data, `,`)
	for _, intv := range ids_arr {
		rule_ids_arr = append(rule_ids_arr, intv)
	}
	return rule_ids_arr
}

// 获取账号的数据权限
func GetDataAuthor(c *GinCtx) ([]interface{}, bool) {
	user_id := c.GetInt64("userID") //当前用户ID
	table_str := ""
	var acount_id []interface{} = Slice{user_id}
	if strings.HasPrefix(c.Request.URL.Path, "/admin/") {
		table_str = "admin"
	} else if strings.HasPrefix(c.Request.URL.Path, "/business/") {
		table_str = "business"
	}
	if table_str != "" {
		role_ids, _ := Model(table_str+"_auth_role_access").Where("uid", user_id).Array("role_id")
		data_access, _ := Model(table_str+"_auth_role").WhereIn("id", role_ids).Array("data_access")
		if IntInVarArray(1, data_access) { //数据权限0=自己1=自己及子权限，2=全部
			chri_role_ids := GetAllChilIds(table_str+"_auth_role", role_ids) //批量获取子节点id
			uid_ids, _ := Model(table_str+"_auth_role_access").WhereIn("role_id", chri_role_ids).Array("uid")
			for _, val := range uid_ids {
				acount_id = append(acount_id, val)
			}
			return acount_id, true //自己及子权限
		} else if IntInVarArray(0, data_access) {
			return acount_id, true //自己
		}
	}
	return acount_id, false //全部
}

// Int类型是否存在Var数组中
func IntInVarArray(target int, arr []*gvar.Var) bool {
	for _, element := range arr {
		if target == Int(element) {
			return true
		}
	}
	return false
}

// Int类型是否存在interface数组中
func IntInInterfaceArray(target int, arr []interface{}) bool {
	for _, element := range arr {
		if target == Int(element) {
			return true
		}
	}
	return false
}

// 转JSON编码为字符串
func JSONToString(data interface{}) string {
	if str, err := json.Marshal(data); err != nil {
		return ""
	} else {
		return string(str)
	}
}

// 字符串转JSON编码
func StringToJSON(val interface{}) interface{} {
	str := val.(string)
	if strings.HasPrefix(str, "{") && strings.HasSuffix(str, "}") {
		var parameter interface{}
		_ = json.Unmarshal([]byte(str), &parameter)
		return parameter
	} else {
		var parameter []interface{}
		_ = json.Unmarshal([]byte(str), &parameter)
		return parameter
	}
}

// tool-获取树状数组
func GetTreeArray(list OrmResult, pid int64, itemprefix string) OrmResult {
	childs := ToolFar(list, pid) //获取pid下的所有数据
	var chridnum OrmResult
	if childs != nil {
		var number int = 1
		var total int = len(childs)
		for _, v := range childs {
			j := ""
			k := ""
			if number == total {
				j += "└"
				k = ""
				if itemprefix != "" {
					k = "&nbsp;"
				}

			} else {
				j += "├"
				k = ""
				if itemprefix != "" {
					k = "│"
				}
			}
			spacer := ""
			if itemprefix != "" {
				spacer = itemprefix + j
			}
			v["spacer"] = gvar.New(spacer)
			v["children"] = gvar.New(GetTreeArray(list, v["id"].Int64(), itemprefix+k+"&nbsp;"))
			chridnum = append(chridnum, v)
			number++
		}
	}
	return chridnum
}

// 将getTreeArray的结果返回为二维数组
func GetTreeToList(list []Map, field string) []Map {
	var midleArr []Map
	for _, v := range list {
		var children []Map
		if childrendata, ok := v["children"]; ok && childrendata != nil {
			switch childrendata.(type) {
			case []interface{}:
				for _, cv := range childrendata.([]interface{}) {
					children = append(children, cv.(Map))
				}
			case []Map:
				children = childrendata.([]Map)
			}
		} else {
			children = nil
		}
		delete(v, "children")
		v[field+"_txt"] = fmt.Sprintf("%v %v", v["spacer"], v[field+""])
		if _, ok := v["id"]; ok {
			midleArr = append(midleArr, v)
		}
		if len(children) > 0 {
			newarr := GetTreeToList(children, field)
			midleArr = ArrayMerge_x(midleArr, newarr)
		}
	}
	return midleArr
}

// 数组拼接
func ArrayMerge_x(ss ...[]Map) []Map {
	n := 0
	for _, v := range ss {
		n += len(v)
	}
	s := make([]Map, 0, n)
	for _, v := range ss {
		s = append(s, v...)
	}
	return s
}

// 获取菜单树形-打包代码菜单
func GetRuleTreeArrayByPack(list OrmResult, pid int64) OrmResult {
	childs := ToolFar(list, pid) //获取pid下的所有数据
	var chridnum OrmResult
	if childs != nil {
		for _, v := range childs {
			newdata := GetRuleTreeArrayByPack(list, v["id"].Int64())
			if newdata != nil {
				v["children"] = gvar.New(GetRuleTreeArrayByPack(list, v["id"].Int64()))
			}
			chridnum = append(chridnum, v)
		}
	}
	return chridnum
}

// base_tool-获取pid下所有数组
func ToolFar(data OrmResult, pid int64) OrmResult {
	var mapString OrmResult
	for _, v := range data {
		if v["pid"].Int64() == pid {
			mapString = append(mapString, v)
		}
	}
	return mapString
}

// 获取子菜单包含的父级ID-返回全部ID
func GetRulesID(tablename string, field string, menus interface{}) interface{} {
	menus_rang := menus.([]interface{})
	var fnemuid []interface{}
	for _, v := range menus_rang {
		fid := getParentID(tablename, field, v)
		if fid != nil {
			fnemuid = MergeArr_interface(fnemuid, fid)
		}
	}
	r_nemu := MergeArr_interface(menus_rang, fnemuid)
	uni_fnemuid := UniqueArr(r_nemu) //去重
	return uni_fnemuid
}

// 获取所有父级ID
func getParentID(tablename string, field string, id interface{}) []interface{} {
	var pids []interface{}
	pid, _ := Model(tablename).Where("id", id).Value(field)
	if pid != nil {
		a_pid := pid.Int64()
		var zr_pid int64 = 0
		if a_pid != zr_pid {
			pids = append(pids, a_pid)
			getParentID(tablename, field, pid)
		}
	}
	return pids
}

// 去重
func UniqueArr(datas []interface{}) []interface{} {
	d := make([]interface{}, 0)
	tempMap := make(map[int]bool, len(datas))
	for _, v := range datas { // 以值作为键名
		keyv := Int(v)
		if tempMap[keyv] == false {
			tempMap[keyv] = true
			d = append(d, v)
		}
	}
	return d
}

// 合并数组-interface
func MergeArr_interface(a, b []interface{}) []interface{} {
	var arr []interface{}
	for _, i := range a {
		arr = append(arr, i)
	}
	for _, j := range b {
		arr = append(arr, j)
	}
	return arr
}

// 将带有逗号的数组中字符串差分合并为数组
func ArraymoreMerge(data []*gvar.Var) []interface{} {
	var rule_ids_arr []interface{}
	for _, mainv := range data {
		ids_arr := strings.Split(mainv.String(), `,`)
		for _, intv := range ids_arr {
			rule_ids_arr = append(rule_ids_arr, intv)
		}
	}
	return rule_ids_arr
}

// 获取树结构数据
func GetTreeData(pdata OrmResult, parent_id int64, pid_file string) OrmResult {
	var returnList OrmResult
	for _, v := range pdata {
		if v[pid_file].Int64() == parent_id {
			children := GetTreeData(pdata, v["id"].Int64(), pid_file)
			if children != nil {
				v["children"] = gvar.New(children)
			}
			returnList = append(returnList, v)
		}
	}
	if returnList == nil {
		returnList = make(OrmResult, 0)
	}
	return returnList
}

// 获取后台菜单子树结构
func GetMenuChildrenArray(pdata OrmResult, parent_id int64, pid_file string) OrmResult {
	var returnList OrmResult
	for _, v := range pdata {
		if v[pid_file].Int64() == parent_id {
			children := GetMenuChildrenArray(pdata, v["id"].Int64(), pid_file)
			if children != nil {
				v["children"] = gvar.New(children)
			}
			returnList = append(returnList, v)
		}
	}
	return returnList
}

// 删除本地附件
func Del_file(file_list []*gvar.Var) {
	path, _ := os.Getwd()
	for _, val := range file_list {
		deldir := filepath.Join(path, val.String())
		os.Remove(deldir)
	}
}

// 删除单文件本地附件
func DelOneFile(file_path string) {
	path, _ := os.Getwd()
	deldir := filepath.Join(path, file_path)
	os.Remove(deldir)
}

// 判断某个数据表是否存在指定字段
// tablename=表名 field=字段
func DbHaseField(tablename, fields string) bool {
	//获取数据库名
	dielddata, _ := DB().Query(ctx, "select COLUMN_NAME from information_schema.columns where TABLE_SCHEMA='"+String(dbConf_arr["dbname"])+"' AND TABLE_NAME='"+tablename+"'")
	var tablefields []interface{}
	for _, val := range dielddata {
		var valjson map[string]interface{}
		mdata, _ := json.Marshal(val)
		json.Unmarshal(mdata, &valjson)
		tablefields = append(tablefields, valjson["COLUMN_NAME"].(string))
	}
	return IsContain(tablefields, fields)
}

// 获取请求参数id-用于数据保存或更新
func GetEditId(idstr interface{}) (f_id float64) {
	if idstr != nil {
		f_id = Float64(idstr)
	} else {
		f_id = 0
	}
	return
}

// 日期时间转时间戳
// timetype时间格式类型  datetime=日期时间 datesecond=日期时间秒date=日期
func StringTimestamp(timeLayout, timetype string) int64 {
	timetpl := "2006-01-02 15:04:05"
	if timetype == "date" {
		timetpl = "2006-01-02"
	} else if timetype == "datetime" {
		timetpl = "2006-01-02 15:04"
	}
	times, _ := time.ParseInLocation(timetpl, timeLayout, time.Local)
	timeUnix := times.Unix()
	return timeUnix
}

// 时间戳格式化为日期字符串
// timetype时间格式类型 date=日期 datetime=日期时间 datesecond=日期时间秒
func TimestampString(timedata interface{}, timetype string) string {
	timetpl := "2006-01-02 15:04:05"
	if timetype == "date" {
		timetpl = "2006-01-02"
	} else if timetype == "datetime" {
		timetpl = "2006-01-02 15:04"
	}
	return time.Unix(timedata.(int64), 0).Format(timetpl)
}

// 获取数据表下的字段值
func GetTalbeFieldVal(tablename, field, id interface{}) *gvar.Var {
	data, _ := Model(tablename).Where("id", id).Value(field)
	return gvar.New(data)
}

// 获取字典数据下的字段值
func GetDicFieldVal(group_id, val interface{}) *gvar.Var {
	if String(val) == "" {
		return VarNew(nil)
	}
	tablename, _ := Model("common_dictionary_group").Where("id", group_id).Value("tablename")
	data, _ := Model(tablename.String()).Where("group_id", group_id).Where("keyvalue", val).Fields("keyname,tagcolor").Find()
	return gvar.New(data)
}

// 判断字符串是否包含
func StrContains(str, filed string) bool {
	return strings.Contains(str, filed)
}

// 把字符串打散为数组
func SplitAndStr(str, step string) []string {
	return strings.Split(str, step)
}

// 把数组转字符串,号分隔
func ArrayToStr(data interface{}, step string) string {
	if data != nil && data != "" {
		data_arr := data.([]interface{})
		var str_arr = make([]string, len(data_arr))
		for k, v := range data_arr {
			str_arr[k] = fmt.Sprintf("%v", v)
		}
		return strings.Join(str_arr, step)
	} else {
		return ""
	}
}

// 判断字符串是否在一个数组中
func StrInArray(target string, str_array []string) bool {
	for _, element := range str_array {
		if target == element {
			return true
		}
	}
	return false
}

// 获取分类下全部子id
func CateAllChilId(tablename string, cid interface{}) []interface{} {
	cids := GetAllChilId(tablename, cid)
	return append(cids, cid)
}

// 判断请求路由是否是该模块
func IsModelPath(path, model string) bool {
	if strings.HasPrefix(path, "/"+model+"/") {
		return true
	} else {
		return false
	}
}

// 隐藏手机号等敏感信息用*替换展示
func HideStrInfo(strtype, val string) string {
	if val == "" {
		return ""
	}
	switch strtype {
	case "email":
		var arr = strings.Split(val, "@")
		var star = ""
		if len(arr[0]) <= 3 {
			star = "*"
			arr[0] = gstr.SubStr(arr[0], 0, len(arr[0])) + star
		} else {
			star = "***"
			arr[0] = gstr.SubStr(arr[0], 0, 1) + star + gstr.SubStr(arr[0], len(arr[0])-1, 1)
		}
		return arr[0] + "@" + arr[1]
	case "mobile":
		if len(val) <= 10 {
			return val
		}
		return val[:3] + "****" + val[len(val)-4:]
	}
	return ""
}

// 创建低代码接口/更新
// 参数:api_id接口数据、 title接口名称、tablename操作数据表、fields获取字段、istoken是否需要登录
func CreateAndUpdateApi(param Map) (api_id interface{}, err error) {
	if _, ok := param["api_id"]; !ok {
		api_id, err = Model("common_api").Data(param).InsertAndGetId()
	} else {
		apidata, _ := Model("common_api").Where("id", param["api_id"]).Value("id")
		if apidata == nil {
			api_id, err = Model("common_api").Data(param).InsertAndGetId()
		} else {
			api_id, err = Model("common_api").Data(param).Where("id", param["id"]).Update()
		}
	}
	return
}

// 删除低代码api
// 参数:api_id接口数据id
func DelApi(api_id interface{}) (err error) {
	_, err = Model("common_api").Where("id", api_id).Delete()
	return
}
