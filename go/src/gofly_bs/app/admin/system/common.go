package system

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
