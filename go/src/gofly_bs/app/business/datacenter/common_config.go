package datacenter

import (
	"gofly/utils/gf"
	"gofly/utils/tools/gcfg"
	"gofly/utils/tools/gconv"
	"gofly/utils/tools/gctx"
	"os"
)

type Common_config struct{}

func init() {
	fpath := Common_config{}
	gf.Register(&fpath, fpath)
}

// 获取配置-动态配置config.yarn
func (api *Common_config) GetConfig(c *gf.GinCtx) {
	ctx := gctx.New()
	appConf, _ := gcfg.Instance().Get(ctx, "app")
	appConf_arr := gconv.Map(appConf)
	gf.Success().SetMsg("获取配置").SetData(gf.Map{"vueobjroot": appConf_arr["vueobjroot"], "mainurl": appConf_arr["mainurl"]}).Regin(c)
}

// 保存系统配置
func (api *Common_config) SaveConfig(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	path, _ := os.Getwd()
	upAppconf := gf.Map{
		"vueobjroot": param["vueobjroot"],
		"mainurl":    param["mainurl"],
	}
	cferr := gf.UpConfFieldData(path, upAppconf, "  ")
	if cferr != nil {
		gf.Failed().SetMsg("修改前端路径失败").SetData(cferr.Error()).Regin(c)
		return
	}
	gf.Success().SetMsg("保存系统配置成功").Regin(c)
}
