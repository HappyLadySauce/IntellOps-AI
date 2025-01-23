package dashboard

import (
	"gofly/utils/gf"
)

/**
* 使用说明：
* 首页统计是根据业务需求数据来统计的，框架无法预知你的项目实际需求，我们只能内置一些方法仅供参考，
* 实际项目开发完成后，根据项目需求自己编写统计数据接口
* business_youtablebane 是你的项目实际数据表(泛指)，不是实际测存在表，切记！自己根据需求开发出对应接口
 */
type Workplace struct{ NoNeedAuths []string }

func init() {
	fpath := Workplace{NoNeedAuths: []string{"*"}}
	gf.Register(&fpath, fpath)
}

// 1获取快捷操作
func (api *Workplace) GetQuick(c *gf.GinCtx) {
	businessID := c.GetInt64("businessID") //当前商户ID
	list, err := gf.Model("business_home_quickop").Where("business_id", businessID).WhereOr("is_common", 1).Fields("id,uid,path_url,name,icon,type,is_common,weigh").Order("weigh asc,id asc").Select()
	if err != nil {
		gf.Failed().SetMsg("获取快捷操作失败").SetData(err).Regin(c)
	} else {
		gf.Success().SetMsg("获取快捷操作数据").SetData(list).Regin(c)
	}
}

// 3保存快捷操作
func (api *Workplace) SaveQuick(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	var f_id = gf.GetEditId(param["id"])
	if f_id == 0 {
		param["uid"] = c.GetInt64("userID")             //当前用户
		param["business_id"] = c.GetInt64("businessID") //当前商户
		addId, err := gf.Model("business_home_quickop").Data(param).InsertAndGetId()
		if err != nil {
			gf.Failed().SetMsg("添加失败").SetData(err).Regin(c)
		} else {
			if addId != 0 {
				gf.Model("business_home_quickop").Data(map[string]interface{}{"weigh": addId}).Where("id", addId).Update()
			}
			gf.Success().SetMsg("添加成功！").SetData(addId).Regin(c)
		}
	} else {
		res, err := gf.Model("business_home_quickop").Data(param).Where("id", f_id).Update()
		if err != nil {
			gf.Failed().SetMsg("更新失败").SetData(err).Regin(c)
		} else {
			gf.Success().SetMsg("更新成功！").SetData(res).Regin(c)
		}
	}
}

// 3删除快捷操作
func (api *Workplace) DelQuick(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	res2, err := gf.Model("business_home_quickop").Where("id", param["id"]).Delete()
	if err != nil {
		gf.Failed().SetMsg("删除失败").SetData(err).Regin(c)
	} else {
		gf.Success().SetMsg("删除成功！").SetData(res2).Regin(c)
	}
}
