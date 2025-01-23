package common

/**
* 系统消息
 */
import (
	"gofly/utils/gf"
)

// 用于自动注册路由
type Message struct{ NoNeedAuths []string }

func init() {
	fpath := Message{NoNeedAuths: []string{"*"}}
	gf.Register(&fpath, fpath)
}

// 获取消息列表
func (api *Message) GetList(c *gf.GinCtx) {
	userID := c.GetInt64("userID")
	usertype := 1 //用户类型
	list, err := gf.Model("common_message").Fields("id,type,title,path,content,isread,createtime").
		WhereIn("usertype", []interface{}{0, usertype}).Where("touid", userID).Order("id desc").Select()
	if err != nil {
		gf.Failed().SetMsg("加载数据失败").Regin(c)
	} else {
		var totalCount int
		totalCount, _ = gf.Model("common_message").WhereIn("usertype", []interface{}{0, usertype}).Where("touid", userID).Count()
		gf.Success().SetMsg("获取全部列表").SetData(gf.Map{
			"total": totalCount,
			"items": list,
		}).Regin(c)
	}

}

// 设置为已读
func (api *Message) Read(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	res2, err := gf.Model("common_message").WhereIn("id", param["ids"]).Data(gf.Map{"isread": 1}).Update()
	if err != nil {
		gf.Failed().SetMsg("更新失败！").Regin(c)
	} else {
		msg := "更新成功！"
		if res2 == nil {
			msg = "暂无数据更新"
		}
		gf.Success().SetMsg(msg).SetData(res2).Regin(c)
	}
}
