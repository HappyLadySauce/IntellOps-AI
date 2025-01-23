package developer

import (
	"encoding/json"
	"gofly/utils/gf"
	"gofly/utils/tools/gcfg"
	"gofly/utils/tools/gconv"
	"gofly/utils/tools/gctx"
	"os"
	"path/filepath"
)

// 用于自动注册路由
type Codestore struct{ NoNeedAuths []string }

func init() {
	fpath := Codestore{NoNeedAuths: []string{"*"}}
	gf.Register(&fpath, fpath)
}

var (
	ctx         = gctx.New()
	appConf, _  = gcfg.Instance().Get(ctx, "app")
	appConf_arr = gconv.Map(appConf)
)

// 获取代码商城分类
func (api *Codestore) GetCodeCate(c *gf.GinCtx) {
	appConf, _ = gcfg.Instance().Get(ctx, "app")
	appConf_arr = gconv.Map(appConf)
	baseurl := c.DefaultQuery("baseurl", "")
	if baseurl == "" {
		gf.Success().SetMsg("代码仓地址不存在！").SetData(make([]interface{}, 0)).Regin(c)
	} else {
		ref := gf.Get(baseurl + "/goflycode/cate/getCate")
		var parameter gf.Map
		if err := json.Unmarshal([]byte(ref), &parameter); err == nil {
			if gconv.Int(parameter["code"]) == 0 {
				path, _ := os.Getwd()
				downdir := filepath.Join(path, "/devsource/codemarket/release")
				gf.Success().SetMsg("获取代码商城分类").SetData(gf.Map{
					"catedata":            parameter["data"],
					"companyPrivateHouse": appConf_arr["companyPrivateHouse"],
					"codepack":            downdir,
					"version":             appConf_arr["version"],
				}).Regin(c)
			} else {
				gf.Failed().SetMsg("请求GoFLy社区获取代码商店分类失败").SetData(parameter).Regin(c)
			}
		}
	}
}

// 获取代码商城
func (api *Codestore) CodeList(c *gf.GinCtx) {
	baseurl := c.DefaultQuery("baseurl", "")
	param, _ := gf.PostParam(c)
	if baseurl == "" {
		gf.Success().SetMsg("代码仓地址不存在！").SetData(make(gf.Slice, 0)).Regin(c)
	} else {
		ref, resErro := gf.HttpGet(baseurl+"/goflycode/content/getCode", param)
		if resErro != nil {
			gf.Failed().SetMsg("请求GoFLy社区获取代码商店失败").SetData(resErro).Regin(c)
		} else {
			if gconv.Int64(ref["code"]) == 0 {
				data := ref["data"].(map[string]interface{})
				list := data["items"].([]interface{})
				path, _ := os.Getwd()
				for _, val := range list {
					item := val.(map[string]interface{})
					installconfig := filepath.Join(path, "/resource/codeinstall", gconv.String(item["name"]))
					if _, err := os.Stat(installconfig); os.IsNotExist(err) { //不存在
						item["is_install"] = false
					} else {
						install_cofig, _ := GetInstallConfig(installconfig)
						item["is_install"] = install_cofig.App.Isinstall
					}
				}
				gf.Success().SetMsg("获取代码商城成功！").SetData(data).Regin(c)
			} else {
				gf.Failed().SetMsg("请求代码商店数据失败").SetData(ref).Regin(c)
			}
		}

	}
}

// 更新私有代码仓地址
func (api *Codestore) UpPrivateHouse(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	path, err := os.Getwd() //获取当前路径
	if err != nil {
		gf.Failed().SetMsg("项目路径获取失败").Regin(c)
		return
	}
	companyPrivateHouse := ""
	if param["companyPrivateHouse"] != nil {
		companyPrivateHouse = gconv.String(param["companyPrivateHouse"])
	}
	//修改应用配置
	upAppconf := gf.Map{
		"companyPrivateHouse": companyPrivateHouse,
	}
	cferr := gf.UpConfFieldData(path, upAppconf, "  ")
	if cferr != nil {
		gf.Failed().SetMsg("修改数据库配置失败").Regin(c)
	} else {
		gf.Success().SetMsg("更新代码仓成功").Regin(c)
	}
}

/***账号登录***/
// 1登录
func (api *Codestore) Login(c *gf.GinCtx) {
	baseurl := c.DefaultQuery("baseurl", "")
	if baseurl == "" {
		gf.Failed().SetMsg("代码仓地址不存在！").Regin(c)
	} else {
		param, _ := gf.RequestParam(c)
		ref := gf.Post(baseurl+"/goflycode/user/login", param, "application/json")
		var parameter map[string]interface{}
		if err := json.Unmarshal([]byte(ref), &parameter); err == nil {
			if gconv.Int(parameter["code"]) == 0 {
				gf.Success().SetMsg("登录").SetData(parameter["data"]).Regin(c)
			} else {
				gf.Failed().SetMsg(gconv.String(parameter["message"])).Regin(c)
			}
		}
	}
}

// 2.免密登录
func (api *Codestore) FreeLogin(c *gf.GinCtx) {
	baseurl := c.DefaultQuery("baseurl", "")
	if baseurl == "" {
		gf.Failed().SetMsg("代码仓地址不存在！").Regin(c)
	} else {
		param, _ := gf.RequestParam(c)
		ref := gf.Post(baseurl+"/goflycode/user/freeLogin", param, "application/json")
		var parameter map[string]interface{}
		if err := json.Unmarshal([]byte(ref), &parameter); err == nil {
			if gconv.Int(parameter["code"]) == 0 {
				gf.Success().SetMsg("登录").SetData(parameter["data"]).Regin(c)
			} else {
				gf.Failed().SetMsg("登录失败").Regin(c)
			}
		}
	}
}

// 3.注册账号
func (api *Codestore) RegisterUser(c *gf.GinCtx) {
	baseurl := c.DefaultQuery("baseurl", "")
	if baseurl == "" {
		gf.Failed().SetMsg("代码仓地址不存在！").Regin(c)
	} else {
		ref := gf.Get(baseurl + "/goflycode/user/registerUser")
		var parameter map[string]interface{}
		if err := json.Unmarshal([]byte(ref), &parameter); err == nil {
			if gconv.Int(parameter["code"]) == 0 {
				gf.Success().SetMsg("注册").SetData(parameter["data"]).Regin(c)
			} else {
				gf.Failed().SetMsg("注册失败").Regin(c)
			}
		}
	}
}

// 4.获取验证码
func (api *Codestore) LoginCode(c *gf.GinCtx) {
	baseurl := c.DefaultQuery("baseurl", "")
	if baseurl == "" {
		gf.Failed().SetMsg("代码仓地址不存在！").Regin(c)
	} else {
		param, _ := gf.RequestParam(c)
		ref := gf.Post(baseurl+"/goflycode/user/loginCode", param, "application/json")
		var parameter map[string]interface{}
		if err := json.Unmarshal([]byte(ref), &parameter); err == nil {
			if gconv.Int(parameter["code"]) == 0 {
				gf.Success().SetMsg("获取验证码").SetData(parameter["data"]).Regin(c)
			} else {
				gf.Failed().SetMsg("获取验证码失败").Regin(c)
			}
		}
	}
}

// 5.把代码推到服务器
func (api *Codestore) UpPackToService(c *gf.GinCtx) {
	baseurl := c.DefaultQuery("baseurl", "")
	if baseurl == "" {
		gf.Failed().SetMsg("代码仓地址不存在！").Regin(c)
	} else {
		param, _ := gf.RequestParam(c)
		pushdata := gf.Map{
			"cid":          param["cid"],
			"code_token":   param["code_token"],
			"title":        param["title"],
			"name":         param["name"],
			"des":          param["des"],
			"price":        param["price"],
			"version":      param["version"],
			"install_file": param["fileurl"],
		}
		ref := gf.Post(baseurl+"/goflycode/content/save", pushdata, "application/json")
		var parameter map[string]interface{}
		if err := json.Unmarshal([]byte(ref), &parameter); err == nil {
			if gconv.Int(parameter["code"]) == 0 {
				gf.Success().SetMsg("代码推到服务器成功").SetData(parameter["data"]).Regin(c)
			} else {
				gf.Failed().SetMsg("代码推到服务器失败").Regin(c)
			}
		}
	}
}

// 6.发布需求
func (api *Codestore) Requirement(c *gf.GinCtx) {
	baseurl := c.DefaultQuery("baseurl", "")
	if baseurl == "" {
		gf.Failed().SetMsg("代码仓地址不存在！").Regin(c)
	} else {
		param, _ := gf.RequestParam(c)
		pushdata := gf.Map{
			"code_token": param["code_token"],
			"title":      param["title"],
			"name":       param["name"],
			"des":        param["des"],
			"price":      param["price"],
			"customer":   param["customer"],
			"mobile":     param["mobile"],
			"wx":         param["wx"],
			"type":       1,
			"cid":        11,
			"status":     2,
			"content":    param["content"],
		}
		ref := gf.Post(baseurl+"/goflycode/content/save", pushdata, "application/json")
		var parameter map[string]interface{}
		if err := json.Unmarshal([]byte(ref), &parameter); err == nil {
			if gconv.Int(parameter["code"]) == 0 {
				gf.Success().SetMsg("发布需求到服务器成功").SetData(parameter["data"]).Regin(c)
			} else {
				gf.Failed().SetMsg("发布需求到服务器失败").Regin(c)
			}
		}
	}
}

// 7.检查更新代码版本
func (api *Codestore) AsyncVersion(c *gf.GinCtx) {
	baseurl := c.DefaultQuery("baseurl", "")
	if baseurl == "" {
		gf.Failed().SetMsg("代码仓地址不存在！").Regin(c)
	} else {
		param, _ := gf.RequestParam(c)
		ref := gf.Post(baseurl+"/goflycode/version/asyncVersion", gf.Map{"code_token": param["code_token"], "version": appConf_arr["version"], "from": "check"}, "application/json")
		var parameter map[string]interface{}
		if err := json.Unmarshal([]byte(ref), &parameter); err == nil {
			if gconv.Int(parameter["code"]) == 0 {
				gf.Success().SetMsg("检查更新代码版本成功").SetData(parameter["data"]).Regin(c)
			} else {
				gf.Failed().SetMsg(gconv.String(parameter["message"])).SetData(parameter).Regin(c)
			}
		} else {
			gf.Failed().SetMsg("请求gofly社区服务失败！").SetData(err).Regin(c)
		}
	}
}

// 8.更新基座代码
func (api *Codestore) UpBaseCode(c *gf.GinCtx) {
	if appConf_arr["runEnv"] == "release" {
		gf.Failed().SetMsg("生产环境禁止操作，请在开发环境下操作").Regin(c)
		return
	}
	baseurl := c.DefaultQuery("baseurl", "")
	if baseurl == "" {
		gf.Failed().SetMsg("代码仓地址不存在！").Regin(c)
	} else {
		param, _ := gf.RequestParam(c)
		ref := gf.Post(baseurl+"/goflycode/version/asyncVersion", gf.Map{"code_token": param["code_token"], "version": appConf_arr["version"], "from": "up"}, "application/json")
		var parameter map[string]interface{}
		if err := json.Unmarshal([]byte(ref), &parameter); err == nil {
			if gconv.Int(parameter["code"]) == 0 {
				codedata := parameter["data"].(map[string]interface{})
				updirpath := "/devsource/developer/upnewversion"
				path, _ := os.Getwd()
				upnewversion := filepath.Join(path, updirpath)
				defer os.RemoveAll(upnewversion) //删除更新源代码包文件
				if _, err := os.Stat(upnewversion); os.IsNotExist(err) {
					os.MkdirAll(upnewversion, os.ModePerm) //创建更新文件容器
				}
				downdir := filepath.Join(path, updirpath, "goflyenterprise.zip")
				downstatus, downdir_zippath := DownFileToDir(gconv.String(codedata["dowurl"]), downdir)
				if downstatus {
					//解压
					dezipdir := filepath.Join(path, updirpath)
					err := Unzip(downdir_zippath, dezipdir)
					if err == nil {
						os.Remove(downdir_zippath)
						//更新配置
						cf, _ := gcfg.New()
						data := cf.MustGet(ctx, "app")
						appConf_arr = gconv.Map(data)
						//开始更新
						CopyAllDir(filepath.Join(path, updirpath, "/goflyenterprise/go"), filepath.Join(path)) //go
						if appConf_arr["vueobjroot"] != "" {
							CopyAllDir(filepath.Join(path, updirpath, "/goflyenterprise/vue/business"), filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(appConf_arr["busDirName"]))) //vue-busseness
							//判断admin端是否存在
							admin_path := filepath.Join(gconv.String(appConf_arr["vueobjroot"]), gconv.String(appConf_arr["adminDirName"]))
							if _, err := os.Stat(admin_path); !os.IsNotExist(err) { //存在
								CopyAllDir(filepath.Join(path, updirpath, "/goflyenterprise/vue/admin"), admin_path) //vue-admin
							}
						}
						//更新版本号
						upAppconf := gf.Map{
							"version": codedata["version"],
						}
						gf.UpConfFieldData(path, upAppconf, "  ")
						gf.Success().SetMsg("更新成功").SetData(true).SetExdata(err).Regin(c)
					} else {
						gf.Failed().SetMsg("解压代码包失败").SetData(err).Regin(c)
					}
				} else {
					gf.Failed().SetMsg("下载代码失败").SetData(parameter).Regin(c)
				}
			} else {
				gf.Failed().SetMsg("检查更新代码版本失败").SetData(parameter).Regin(c)
			}
		}
	}
}

// 检测包名是否可用
func (api *Codestore) CheckPackName(c *gf.GinCtx) {
	baseurl := c.DefaultQuery("baseurl", "")
	if baseurl == "" {
		gf.Success().SetMsg("代码仓地址不存在！").SetData(make([]interface{}, 0)).Regin(c)
	} else {
		param, _ := gf.PostParam(c)
		ref := gf.Post(baseurl+"/goflycode/ident/checkPackName", gf.Map{"type": param["type"], "name": param["name"]}, "application/json")
		var parameter gf.Map
		if err := json.Unmarshal([]byte(ref), &parameter); err == nil {
			if gconv.Int(parameter["code"]) == 0 {
				gf.Success().SetMsg("检测包名结果").SetData(parameter["data"]).Regin(c)
			} else {
				gf.Failed().SetMsg("请求GoFLy社区获取检查数据失败").Regin(c)
			}
		}
	}
}

// 提交标识占用
func (api *Codestore) SavePackName(c *gf.GinCtx) {
	baseurl := c.DefaultQuery("baseurl", "")
	if baseurl == "" {
		gf.Success().SetMsg("代码仓地址不存在！").SetData(make([]interface{}, 0)).Regin(c)
	} else {
		param, _ := gf.PostParam(c)
		ref := gf.Post(baseurl+"/goflycode/ident/savePackName", gf.Map{"type": param["type"], "name": param["name"]}, "application/json")
		var parameter gf.Map
		if err := json.Unmarshal([]byte(ref), &parameter); err == nil {
			if gconv.Int(parameter["code"]) == 0 {
				gf.Success().SetMsg("提交标识占用").SetData(parameter["data"]).Regin(c)
			} else {
				gf.Failed().SetMsg("请求GoFLy社区提交标识占用失败").SetData(parameter).Regin(c)
			}
		}
	}
}

// 检测插件是否已经支付
func (api *Codestore) CheckIsPay(c *gf.GinCtx) {
	baseurl := c.DefaultQuery("baseurl", "")
	if baseurl == "" {
		gf.Success().SetMsg("代码仓地址不存在！").SetData(make([]interface{}, 0)).Regin(c)
	} else {
		param, _ := gf.PostParam(c)
		ref := gf.Post(baseurl+"/goflycode/order/checkIsPay", param, "application/json")
		var parameter gf.Map
		if err := json.Unmarshal([]byte(ref), &parameter); err == nil {
			if gconv.Int(parameter["code"]) == 0 {
				gf.Success().SetMsg("检测插件是否已经支付成功").SetData(parameter["data"]).SetExdata(parameter).Regin(c)
			} else {
				gf.Failed().SetMsg("检测插件是否已经支付失败").Regin(c)
			}
		}
	}
}

// 提交插件订单
func (api *Codestore) AddOrder(c *gf.GinCtx) {
	baseurl := c.DefaultQuery("baseurl", "")
	if baseurl == "" {
		gf.Success().SetMsg("gofly请求地址不存在！").SetData(make([]interface{}, 0)).Regin(c)
	} else {
		param, _ := gf.PostParam(c)
		ref := gf.Post(baseurl+"/goflycode/order/addOrder", param, "application/json")
		var parameter gf.Map
		if err := json.Unmarshal([]byte(ref), &parameter); err == nil {
			if gconv.Int(parameter["code"]) == 0 {
				gf.Success().SetMsg("提交插件订单成功").SetData(parameter["data"]).Regin(c)
			} else {
				gf.Failed().SetMsg("提交插件订单失败").SetData(parameter).Regin(c)
			}
		}
	}
}
