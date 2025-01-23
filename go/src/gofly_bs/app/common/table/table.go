package table

import (
	"gofly/utils/gf"
	"gofly/utils/tools/gcfg"
	"gofly/utils/tools/gconv"
	"gofly/utils/tools/gctx"
)

// 数据表操作
type Index struct{}

func init() {
	fpath := Index{}
	gf.Register(&fpath, fpath)
}

var (
	ctx        = gctx.New()
	dbConf, _  = gcfg.Instance().Get(ctx, "database.default")
	dbConf_arr = gconv.Map(dbConf)
)

// 更新排序
func (api *Index) TableWeigh(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	res, err := gf.Model(param["tableanme"]).Where("pid", param["pid"]).Save(param["weighList"])
	if err != nil {
		gf.Failed().SetMsg("排序更新失败！").SetData(err).Regin(c)
	} else {
		gf.Success().SetMsg("排序更新成功！").SetData(res).Regin(c)
	}
}

// 获取锁数据表
func (api *Index) GetTables(c *gf.GinCtx) {
	tablelist, _ := gf.DB().Query(c, "select TABLE_NAME,TABLE_COMMENT from information_schema.tables where table_schema = '"+gconv.String(dbConf_arr["dbname"])+"'")
	var talbe_list []interface{}
	for _, Val := range tablelist {
		talbe_list = append(talbe_list, map[string]interface{}{"name": Val["TABLE_NAME"], "title": Val["TABLE_COMMENT"]})
	}
	gf.Success().SetMsg("获取锁数据表").SetData(talbe_list).Regin(c)
}
