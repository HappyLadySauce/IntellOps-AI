package datacenter

import (
	"fmt"
	"gofly/utils/gf"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Configuration struct{}

func init() {
	fpath := Configuration{}
	gf.Register(&fpath, fpath)
}

// 获取邮箱
func (api *Configuration) GetEmail(c *gf.GinCtx) {
	businessID := c.GetInt64("businessID") //当前商户ID
	data, _ := gf.Model("common_email").Where("business_id", businessID).Find()
	gf.Success().SetMsg("获取邮箱").SetData(data).Regin(c)
}

// 保存邮箱
func (api *Configuration) SaveEmail(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	businessID := c.GetInt64("businessID") //当前商户ID
	GetID, _ := gf.Model("common_email").Where("business_id", businessID).Value("id")
	if GetID == nil {
		param["business_id"] = businessID
		param["data_from"] = "business"
		addId, err := gf.Model("common_email").Data(param).InsertAndGetId()
		if err != nil {
			gf.Failed().SetMsg("添加失败").SetData(err).Regin(c)
		} else {
			gf.Success().SetMsg("添加成功！").SetData(addId).Regin(c)
		}
	} else {
		res, err := gf.Model("common_email").Data(param).Where("id", GetID).Update()
		if err != nil {
			gf.Failed().SetMsg("更新失败").SetData(err).Regin(c)
		} else {
			gf.Success().SetMsg("更新成功！").SetData(res).Regin(c)
		}
	}
}

// 获取安装的代码仓配置
func (api *Configuration) GetCodestoreConfig(c *gf.GinCtx) {
	path, _ := os.Getwd()
	configPath := filepath.Join(path, "/resource/config")
	go_app_dir, _ := gf.GetAllFileArray(configPath)
	var list []interface{} = make([]interface{}, 0)
	for _, val := range gf.Strings(go_app_dir["files"]) {
		configstr := gf.ReaderFileByline(filepath.Join(path, "/resource/config", val))
		fileName := strings.Split(val, ".")
		install_cofig, _ := gf.GetYmlConfigData(filepath.Join(path, "/resource/config"), fileName[0])
		data_item := install_cofig.(map[string]interface{})
		if data_item["conftype"] == "configuration" {
			var new_data []map[string]interface{}
			for k, v := range data_item["data"].(map[string]interface{}) {
				keyname := k
				for _, cfstr := range configstr {
					cf_str := gf.String(cfstr)
					cf_str_arr := strings.Split(cf_str, "#")
					if strings.Contains(cf_str, k) && len(cf_str_arr) == 2 {
						keyname = cf_str_arr[1]
					}
				}
				new_data = append(new_data, gf.Map{"keyname": keyname, "keyfield": k, "keyvalue": v})
			}
			data_item["name"] = fileName[0]
			data_item["data"] = new_data
			list = append(list, install_cofig)
		}
	}
	gf.Success().SetMsg("获取安装的代码仓配置列表").SetData(list).Regin(c)
}

// 修改安装的代码仓配置
func (api *Configuration) SaveCodeStoreConfig(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	path, _ := os.Getwd()
	configPath := filepath.Join(path, "/resource/config", gf.String(param["name"])+".yaml")
	var upAppconf gf.Map = make(map[string]interface{}, 0)
	for _, val := range param["data"].([]interface{}) {
		item := val.(map[string]interface{})
		var val_data interface{}
		switch item["keyvalue"].(type) {
		case string: // 处理 string 类型
			val_str := gf.String(item["keyvalue"])
			if val_str == "" {
				val_str = "\"\""
			} else {
				val_str = strconv.Quote(val_str)
			}
			val_data = val_str
		default:
			val_data = gf.String(item["keyvalue"])
		}
		upAppconf[gf.String(item["keyfield"])] = fmt.Sprintf("%s  #%v", val_data, item["keyname"])
	}
	gf.UpConfigFild(configPath, upAppconf, "  ")
	gf.Success().SetMsg("修改安装的代码仓配置").SetData(param).SetExdata(configPath).Regin(c)
}

// 更新配置使用状态
func (api *Configuration) UpConfigStatus(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	path, _ := os.Getwd()
	//如果改为true，则把其他通用标识改为false
	if gf.Bool(param["status"]) {
		path, _ := os.Getwd()
		configPath := filepath.Join(path, "/resource/config")
		go_app_dir, _ := gf.GetAllFileArray(configPath)
		for _, val := range gf.Strings(go_app_dir["files"]) {
			fileName := strings.Split(val, ".")
			install_cofig, _ := gf.GetYmlConfigData(filepath.Join(path, "/resource/config"), fileName[0])
			data_item := install_cofig.(map[string]interface{})
			if fileName[0] != gf.String(param["name"]) && gf.String(data_item["pluginident"]) == gf.String(param["pluginident"]) && gf.Bool(data_item["status"]) {
				configPath := filepath.Join(path, "/resource/config", val)
				gf.UpConfigFild(configPath, gf.Map{"status": false}, "")
			}
		}
	}
	configPath := filepath.Join(path, "/resource/config", gf.String(param["name"])+".yaml")
	gf.UpConfigFild(configPath, gf.Map{"status": param["status"]}, "")
	gf.Success().SetMsg("更新配置使用状态成功").SetData(param).SetExdata(configPath).Regin(c)
}
