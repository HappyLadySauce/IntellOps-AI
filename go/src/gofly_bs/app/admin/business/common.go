package business

import (
	"strings"

	"gofly/utils/gf"
	"gofly/utils/tools/gvar"
)

// 获取菜单子树结构
func GetMenuChildrenArray(pdata gf.OrmResult, parent_id int64) gf.OrmResult {
	var returnList gf.OrmResult
	for _, v := range pdata {
		if v["pid"].Int64() == parent_id {
			children := GetMenuChildrenArray(pdata, v["id"].Int64())
			if children != nil {
				v["children"] = gvar.New(children)
			}
			returnList = append(returnList, v)
		}
	}
	return returnList
}

// tool-获取树状数组
func GetTreeArray(num gf.OrmResult, pid int64, itemprefix string) gf.OrmResult {
	childs := ToolFar(num, pid) //获取pid下的所有数据
	var chridnum gf.OrmResult
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
			v["childlist"] = gvar.New(GetTreeArray(num, v["id"].Int64(), itemprefix+k+"&nbsp;"))
			chridnum = append(chridnum, v)
			number++
		}
	}
	return chridnum
}

// base_tool-获取pid下所有数组
func ToolFar(data gf.OrmResult, pid int64) gf.OrmResult {
	var mapString gf.OrmResult
	for _, v := range data {
		if v["pid"].Int64() == pid {
			mapString = append(mapString, v)
		}
	}
	return mapString
}

// 数组拼接
func ArrayMerge(ss ...gf.OrmResult) gf.OrmResult {
	n := 0
	for _, v := range ss {
		n += len(v)
	}
	s := make(gf.OrmResult, 0, n)
	for _, v := range ss {
		s = append(s, v...)
	}
	return s
}

// 判断元素是否存在数组中
func IsContain(items []interface{}, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// 多维数组合并
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

// 把字符串打散为数组
func Axplode(data interface{}) []interface{} {
	var rule_ids_arr []interface{}
	ids_arr := strings.Split(data.(string), `,`)
	for _, intv := range ids_arr {
		rule_ids_arr = append(rule_ids_arr, intv)
	}
	return rule_ids_arr
}
