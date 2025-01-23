package system

import (
	"gofly/utils/gf"
	"gofly/utils/tools/gconv"
	"gofly/utils/tools/gmap"
)

// 部门管理
type Dept struct{ NoNeedAuths []string }

func init() {
	fpath := Dept{NoNeedAuths: []string{"getParent"}}
	gf.Register(&fpath, fpath)
}

// 获取部门列表
func (api *Dept) GetList(c *gf.GinCtx) {
	//搜索添条件
	param, _ := gf.RequestParam(c)
	whereMap := gmap.New()
	if name, ok := param["name"]; ok && name != "" {
		whereMap.Set("name like ?", "%"+gconv.String(name)+"%")
	}
	if status, ok := param["status"]; ok && status != "" {
		whereMap.Set("status", status)
	}
	if createtime, ok := param["createtime"]; ok && createtime != "" {
		datetime_arr := gf.SplitAndStr(gf.String(createtime), ",")
		whereMap.Set("createtime between ? and ?", gf.Slice{datetime_arr[0] + " 00:00", datetime_arr[1] + " 23:59"})
	}
	list, _ := gf.Model("admin_auth_dept").Where(whereMap).Order("weigh asc").Select()
	if len(list) > 0 {
		max_role_id, _ := gf.Model("admin_auth_dept").Where(whereMap).Order("id asc").Value("pid")
		list = gf.GetTreeArray(list, gf.Int64(max_role_id), "")
	}
	gf.Success().SetMsg("获取部门列表").SetData(list).Regin(c)
}

// 获取部门列表-表单
func (api *Dept) GetParent(c *gf.GinCtx) {
	list, _ := gf.Model("admin_auth_dept").Fields("id,pid,name").Order("weigh asc").Select()
	list = gf.GetMenuChildrenArray(list, 0, "pid")
	gf.Success().SetMsg("获取部门列表").SetData(list).Regin(c)
}

// 保存
func (api *Dept) Save(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	var f_id = gf.GetEditId(param["id"])
	if f_id == 0 {
		param["account_id"] = c.GetInt64("userID") //当前用户ID
		addId, err := gf.Model("admin_auth_dept").Data(param).InsertAndGetId()
		if err != nil {
			gf.Failed().SetMsg("添加失败").SetData(err).Regin(c)
		} else {
			if addId != 0 {
				gf.Model("admin_auth_dept").Data(map[string]interface{}{"weigh": addId}).Where("id", addId).Update()
			}
			gf.Success().SetMsg("添加成功！").SetData(addId).Regin(c)
		}
	} else {
		res, err := gf.Model("admin_auth_dept").Data(param).Where("id", f_id).Update()
		if err != nil {
			gf.Failed().SetMsg("更新失败").SetData(err).Regin(c)
		} else {
			gf.Success().SetMsg("更新成功！").SetData(res).Regin(c)
		}
	}
}

// 更新状态
func (api *Dept) UpStatus(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	res, err := gf.Model("admin_auth_dept").Where("id", param["id"]).Data(map[string]interface{}{"status": param["status"]}).Update()
	if err != nil {
		gf.Failed().SetMsg("更新失败！").SetData(err).Regin(c)
	} else {
		msg := "更新成功！"
		if res == nil {
			msg = "暂无数据更新"
		}
		gf.Success().SetMsg(msg).Regin(c)
	}
}

// 删除
func (api *Dept) Del(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	res, err := gf.Model("admin_auth_dept").WhereIn("id", param["ids"]).Delete()
	if err != nil {
		gf.Failed().SetMsg("删除失败").SetData(err).Regin(c)
	} else {
		gf.Success().SetMsg("删除成功！").SetData(res).Regin(c)
	}
}
