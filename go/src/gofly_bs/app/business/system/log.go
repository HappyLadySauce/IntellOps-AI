package system

import (
	"gofly/utils/gf"
	"gofly/utils/tools/gmap"
	"gofly/utils/tools/gtime"
	"net"
)

// 登录和操作日志
type Log struct {
	NoNeedAuths []string
	NoNeedLogin []string
}

func init() {
	fpath := Log{NoNeedLogin: []string{"getList"}}
	gf.Register(&fpath, fpath)
}

// 获取登录日志列表
func (api *Log) GetLogin(c *gf.GinCtx) {
	pageNo := gf.Int(c.DefaultQuery("page", "1"))
	pageSize := gf.Int(c.DefaultQuery("pageSize", "10"))
	//搜索添条件
	param, _ := gf.RequestParam(c)
	whereMap := gmap.New()
	if gf.DbHaseField("common_sys_login_log", "business_id") {
		whereMap.Set("business_id", c.GetInt64("businessID"))
	}
	if user, ok := param["user"]; ok && user != "" {
		userids, _ := gf.Model("business_account").Where("name like ?", "%"+gf.String(user)+"%").Array("id")
		whereMap.Set("uid IN(?)", userids)
	}
	if ip, ok := param["ip"]; ok && ip != "" {
		address := net.ParseIP(gf.String(ip))
		if address == nil {
			whereMap.Set("address like ?", "%"+gf.String(ip)+"%")
		} else {
			whereMap.Set("ip", ip)
		}
	}
	if status, ok := param["status"]; ok && status != "" {
		whereMap.Set("status", status)
	}
	if createtime, ok := param["createtime"]; ok && createtime != "" {
		datetime_arr := gf.SplitAndStr(gf.String(createtime), ",")
		whereMap.Set("createtime between ? and ?", gf.Slice{datetime_arr[0] + " 00:00", datetime_arr[1] + " 23:59"})
	}
	MDB := gf.Model("common_sys_login_log").Where("type", "business").Where(whereMap)
	totalCount, _ := MDB.Clone().Count()
	list, err := MDB.Page(pageNo, pageSize).Order("id desc").Select()
	if err != nil {
		gf.Failed().SetMsg(err.Error()).Regin(c)
	} else {
		for _, val := range list {
			userdata, _ := gf.Model("business_account").Where("id", val["uid"]).Fields("username,name,nickname,avatar").Find()
			val["user"] = gf.VarNew(userdata)
		}
		gf.Success().SetMsg("获取登录日志").SetData(gf.Map{
			"page":     pageNo,
			"pageSize": pageSize,
			"total":    totalCount,
			"items":    list}).Regin(c)
	}
}

// 删除上个月登录日志
func (api *Log) DelLastLogin(c *gf.GinCtx) {
	res, err := gf.Model("common_sys_login_log").Where("type !=?", "admin").Where("createtime <", gtime.Now().AddDate(0, -1, 0).Format("Y-m-d H:i:s")).Delete()
	if err != nil {
		gf.Failed().SetMsg("删除失败").SetData(err).Regin(c)
	} else {
		gf.Success().SetMsg("删除成功！").SetData(res).Regin(c)
	}
}

// 获取操作日志列表
func (api *Log) GetOperation(c *gf.GinCtx) {
	pageNo := gf.Int(c.DefaultQuery("page", "1"))
	pageSize := gf.Int(c.DefaultQuery("pageSize", "10"))
	//搜索添条件
	param, _ := gf.RequestParam(c)
	whereMap := gmap.New()
	if gf.DbHaseField("common_sys_login_log", "business_id") {
		whereMap.Set("business_id", c.GetInt64("businessID"))
	}
	if user, ok := param["user"]; ok && user != "" {
		userids, _ := gf.Model("business_account").Where("name like ?", "%"+gf.String(user)+"%").Array("id")
		whereMap.Set("uid IN(?)", userids)
	}
	if ip, ok := param["ip"]; ok && ip != "" {
		address := net.ParseIP(gf.String(ip))
		if address == nil {
			whereMap.Set("address like ?", "%"+gf.String(ip)+"%")
		} else {
			whereMap.Set("ip", ip)
		}
	}
	if status, ok := param["status"]; ok && status != "" {
		if gf.Int64(status) == 200 {
			whereMap.Set("status", status)
		} else {
			whereMap.Set("status !=", 200)
		}
	}
	if createtime, ok := param["createtime"]; ok && createtime != "" {
		datetime_arr := gf.SplitAndStr(gf.String(createtime), ",")
		whereMap.Set("createtime between ? and ?", gf.Slice{datetime_arr[0] + " 00:00", datetime_arr[1] + " 23:59"})
	}
	MDB := gf.Model("common_sys_operation_log").Where("type !=?", "admin").Where(whereMap)
	totalCount, _ := MDB.Clone().Count()
	list, err := MDB.Page(pageNo, pageSize).Fields("id,uid,request_method,url,ip,address,des,latency,status,createtime").Order("id desc").Select()
	if err != nil {
		gf.Failed().SetMsg(err.Error()).Regin(c)
	} else {
		for _, val := range list {
			userdata, _ := gf.Model("business_account").Where("id", val["uid"]).Fields("username,name,nickname,avatar").Find()
			val["user"] = gf.VarNew(userdata)
		}
		gf.Success().SetMsg("获取登录日志").SetData(gf.Map{
			"page":     pageNo,
			"pageSize": pageSize,
			"total":    totalCount,
			"items":    list}).Regin(c)
	}
}

// 删除上个月操作日志
func (api *Log) DelLastOperation(c *gf.GinCtx) {
	res, err := gf.Model("common_sys_operation_log").Where("type !=?", "admin").Where("createtime <", gtime.Now().AddDate(0, -1, 0).Format("Y-m-d H:i:s")).Delete()
	if err != nil {
		gf.Failed().SetMsg("删除失败").SetData(err).Regin(c)
	} else {
		gf.Success().SetMsg("删除成功！").SetData(res).Regin(c)
	}
}

// 获取操作日志内容
func (api *Log) GetOperationDetail(c *gf.GinCtx) {
	id := c.DefaultQuery("id", "")
	if id == "" {
		gf.Failed().SetMsg("请传参数id").Regin(c)
	} else {
		data, err := gf.Model("common_sys_operation_log").Where("id", id).Find()
		if err != nil {
			gf.Failed().SetMsg("获取内容失败").SetData(err).Regin(c)
		} else {
			if data != nil {
				data["username"], _ = gf.Model("business_account").Where("id", data["uid"]).Value("name")
			}
			gf.Success().SetMsg("获取内容成功！").SetData(data).Regin(c)
		}
	}
}
