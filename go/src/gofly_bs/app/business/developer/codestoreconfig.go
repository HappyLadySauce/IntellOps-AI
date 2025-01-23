package developer

import (
	"github.com/spf13/viper"
)

// 获取配置
type App struct {
	Version      string `yaml:"version"`
	Title        string `yaml:"title"`
	Des          string `yaml:"des"`
	Name         string `yaml:"name"`
	Isinstall    bool   `yaml:"isinstall"`
	Installcover bool   `yaml:"installcover"`
}

// 前端
type Installvue struct {
	//admin
	Viewsfilesadmin       string `yaml:"viewsfilesadmin "`
	Viewsfoldersadmin     string `yaml:"viewsfoldersadmin "`
	Componentfilesadmin   string `yaml:"componentfilesadmin "`
	Componentfoldersadmin string `yaml:"componentfoldersadmin "`
	//business
	Viewsfilesbusiness       string `yaml:"viewsfilesbusiness"`
	Viewsfoldersbusiness     string `yaml:"viewsfoldersbusiness"`
	Componentfilesbusiness   string `yaml:"componentfilesbusiness"`
	Componentfoldersbusiness string `yaml:"componentfoldersbusiness"`
}
type Installgo struct {
	Appfolders      string `yaml:"appfolders"`
	Utilsfiles      string `yaml:"utilsfiles"`
	NoVerifyAPIRoot string `yaml:"noVerifyAPIRoot"`
}
type Sqldb struct {
	Packtables      string `yaml:"packtables"`
	Businessmenuids string `yaml:"businessmenuids"`
	Adminmenuids    string `yaml:"adminmenuids"`
}
type Config struct {
	App        App        `yaml:"app"`
	Installvue Installvue `yaml:"installvue"`
	Installgo  Installgo  `yaml:"installgo"`
	Sqldb      Sqldb      `yaml:"sqldb"`
}

// 读取Yaml配置文件，并转换成Config对象  struct结构
func GetInstallConfig(path string) (*Config, error) {
	var config *Config
	vip := viper.New()
	vip.AddConfigPath(path)     //设置读取的文件路径
	vip.SetConfigName("config") //设置读取的文件名
	vip.SetConfigType("yaml")   //设置文件的类型
	//尝试进行配置读取
	if err := vip.ReadInConfig(); err != nil {
		return nil, err
	}
	verr := vip.Unmarshal(&config)
	if verr != nil {
		return nil, verr
	}
	return config, nil
}
