package user

import (
	"gofly/utils/gf"
)

// 获取权限菜单
func GetMenuArray(pdata gf.OrmResult, parent_id int64, roles []interface{}) []map[string]interface{} {
	var returnList []map[string]interface{}
	var one int64 = 1
	for _, v := range pdata {
		if v["pid"].Int64() == parent_id {
			mid_item := map[string]interface{}{
				"path":      v["routepath"],
				"name":      v["routename"],
				"component": v["component"],
			}
			children := GetMenuArray(pdata, v["id"].Int64(), roles)
			if children != nil {
				mid_item["children"] = children
			}
			//1.标题
			// var Menu_title interface{}
			// if v["locale"] != nil && v["locale"].String() != "" {
			// 	Menu_title = v["locale"]
			// } else {
			// 	Menu_title = v["title"]
			// }
			meta := map[string]interface{}{
				"locale": v["locale"],
				"title":  v["title"],
				"id":     v["id"],
			}
			//2.重定向
			if v["redirect"] != nil && v["redirect"].String() != "" {
				mid_item["redirect"] = v["redirect"]
			}
			//3.隐藏子菜单
			if v["hidechildreninmenu"] != nil && v["hidechildreninmenu"].Int64() == one {
				meta["hideChildrenInMenu"] = true
			}
			//3.图标
			if v["icon"] != nil && v["icon"].String() != "" {
				meta["icon"] = v["icon"]
			}
			//4.缓存
			if v["keepalive"] != nil && v["keepalive"].Int64() == one {
				meta["ignoreCache"] = false
			} else {
				meta["ignoreCache"] = true
			}
			//5.隐藏菜单
			if v["hideinmenu"] != nil && v["hideinmenu"].Int64() == one {
				meta["hideInMenu"] = true
			}
			//6.在标签隐藏
			if v["noaffix"] != nil && v["noaffix"].Int64() == one {
				meta["noAffix"] = true
			}
			//7.详情页在本业打开-用于配置详情页时左侧激活的菜单路径
			if v["activemenu"] != nil && v["activemenu"].Int64() == one {
				meta["activeMenu"] = true
			}
			//8.是否需要登录鉴权
			if v["requiresauth"] != nil && v["requiresauth"].Int64() == one {
				meta["requiresAuth"] = true
			}
			//9.是否需要登录鉴权
			if v["isext"] != nil && v["isext"].Int64() == one {
				meta["isExt"] = true
			}
			//10.是否需要登录鉴权
			if v["onlypage"] != nil && v["onlypage"].Int64() == one {
				meta["onlypage"] = true
			}
			//11.按钮权限
			if len(roles) == 0 { //超级权限
				permission, _ := gf.Model("admin_auth_rule").Where("status", 0).Where("type", 2).Where("pid", v["id"]).Array("permission")
				if permission != nil && len(permission) > 0 {
					meta["btnroles"] = permission
				} else {
					meta["btnroles"] = [1]string{"*"}
				}
			} else { //选择路由
				permission, _ := gf.Model("admin_auth_rule").Where("status", 0).Where("type", 2).Where("pid", v["id"]).WhereIn("id", roles).Array("permission")
				if permission != nil && len(permission) > 0 {
					meta["btnroles"] = permission
				} else {
					hasepermission, _ := gf.Model("admin_auth_rule").Where("status", 0).Where("type", 2).Where("pid", v["id"]).Array("permission")
					if hasepermission == nil {
						meta["btnroles"] = make([]interface{}, 0)
					} else {
						meta["btnroles"] = [1]string{"*"}
					}
				}
			}
			//赋值
			mid_item["meta"] = meta
			returnList = append(returnList, mid_item)
		}
	}
	return returnList
}
