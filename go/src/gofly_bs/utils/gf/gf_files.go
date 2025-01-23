package gf

import (
	"bufio"
	"errors"
	"fmt"
	"gofly/utils/tools/gcfg"
	"gofly/utils/tools/gconv"
	"gofly/utils/tools/gctx"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/viper"
)

// 获取附件访问地址
func GetMainURL() string {
	path, _ := os.Getwd()
	configPath := filepath.Join(path, "/resource/config")
	go_app_dir, _ := GetAllFileArray(configPath)
	hase_thri := ""
	for _, val := range Strings(go_app_dir["files"]) {
		fileName := strings.Split(val, ".")
		install_cofig, _ := GetYmlConfigData(filepath.Join(path, "/resource/config"), fileName[0])
		data_item := install_cofig.(map[string]interface{})
		if data_item["pluginident"] == "upload" && Bool(data_item["status"]) {
			data := data_item["data"].(map[string]interface{})
			if domain, ok := data["domain"]; ok {
				hase_thri = String(domain)
				if !strings.HasSuffix(hase_thri, "/") {
					hase_thri = hase_thri + "/"
				}
			}
		}
	}
	if hase_thri == "" {
		ctx := gctx.New()
		appConf, _ := gcfg.Instance().Get(ctx, "app")
		appConf_arr := gconv.Map(appConf)
		return String(appConf_arr["mainurl"])
	} else {
		return hase_thri
	}
}

// 获取本地附件访问地址
func GetMainURLLocal() string {
	ctx := gctx.New()
	appConf, _ := gcfg.Instance().Get(ctx, "app")
	appConf_arr := gconv.Map(appConf)
	return String(appConf_arr["mainurl"])
}

// 补全附件访问路径（系统通用）
func GetFullPath(url string) (text string) {
	var domainRegex = regexp.MustCompile(`^(http(s)?:\/\/)\w+[^\s]+(\.[^\s]+){1,}$`)
	if url == "" || domainRegex.MatchString(url) {
		text = url
	} else if strings.HasPrefix(url, "resource/uploads/") { //自己服务器本地存附件
		text = GetMainURLLocal() + url
	} else { //其他附附件更加系统配置上传方式选择
		text = GetMainURL() + url
	}
	return
}

// 获取附件访问地址 localPath是默认本地上传路径(不填则默认地址)
func GetUploadURL(localPath, valkey string) string {
	path, _ := os.Getwd()
	configPath := filepath.Join(path, "/resource/config")
	go_app_dir, _ := GetAllFileArray(configPath)
	thirdPath := ""
	for _, val := range Strings(go_app_dir["files"]) {
		fileName := strings.Split(val, ".")
		install_cofig, _ := GetYmlConfigData(filepath.Join(path, "/resource/config"), fileName[0])
		data_item := install_cofig.(map[string]interface{})
		if data_item["pluginident"] == "upload" && Bool(data_item["status"]) {
			data := data_item["data"].(map[string]interface{})
			if uploadpath, ok := data[valkey]; ok {
				thirdPath = String(uploadpath)
			}
		}
	}
	if thirdPath == "" {
		if localPath == "" {
			localPath = "/common/upload/upFile"
		}
		return localPath
	} else {
		return thirdPath
	}
}

// 覆盖写入文件
// filePath文件路径
func WriteToFile(filePath string, content string) error {
	if _, err := os.Stat(filePath); err != nil {
		if !os.IsExist(err) {
			pathstr_arr := strings.Split(filePath, `/`)
			path_dirs := strings.Split(filePath, (pathstr_arr[len(pathstr_arr)-1]))
			os.MkdirAll(path_dirs[0], os.ModePerm)
			os.Create(filePath)
		}
	}
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("file create failed. err: " + err.Error())
	} else {
		n, _ := f.Seek(0, os.SEEK_END)
		_, err = f.WriteAt([]byte(content), n)
		defer f.Close()
	}
	return err
}

// 逐行读取文件
// filePath文件路径
func ReaderFileByline(filePath string) []interface{} {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	var list []interface{}
	for {
		line, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		}
		list = append(list, string(line))
	}
	return list
}

// 一次性读取全部文件
// filePath文件路径
func ReaderFileBystring(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err.Error())
	}
	return string(bytes)
}

// DownPic 远程下载图片
func DownPic(src, dest string) (string, error) {
	re, err := http.Get(src)
	if err != nil {
		return "", err
	}
	defer re.Body.Close()
	fix := "png"
	if idx := strings.LastIndex(src, "."); idx != -1 {
		fix = strings.ToLower(src[idx+1:])
		if strings.Contains(fix, "?") {
			fix_arr := strings.Split(fix, "?")
			fix = fix_arr[0]
		}
	}
	if fix == "" {
		return "", errors.New(fmt.Sprintf("unknow pic type, pic path: %s", src))
	}
	thumbF, err := os.OpenFile(dest+"."+fix, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return "", err
	}
	defer thumbF.Close()
	if fix == "jpeg" || fix == "jpg" {
		img, err := jpeg.Decode(re.Body)
		if err != nil {
			return "", err
		}
		if err = jpeg.Encode(thumbF, img, &jpeg.Options{Quality: 40}); err != nil {
			return "", err
		}
	} else if fix == "png" {
		img, err := png.Decode(re.Body)
		if err != nil {
			return "", err
		}
		if err = png.Encode(thumbF, img); err != nil {
			return "", err
		}
	} else if fix == "gif" {
		img, err := gif.Decode(re.Body)
		if err != nil {
			return "", err
		}
		if err = gif.Encode(thumbF, img, nil); err != nil {
			return "", err
		}
	} else {
		return "", errors.New("不支持的格式")
	}
	return "." + fix, nil
}

// 获取文文件夹下的文件及文件返回数组
func GetAllFileArray(pathname string) (map[string]interface{}, error) {
	var folders = make([]string, 0)
	var files = make([]string, 0)
	rd, err := os.ReadDir(pathname)
	if err != nil {
		return map[string]interface{}{"folders": folders, "files": files}, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			folders = append(folders, fi.Name())
		} else {
			files = append(files, fi.Name())
		}
	}
	return map[string]interface{}{"folders": folders, "files": files}, nil
}

// 读取Yaml配置文件， struct结构
func GetYmlConfigData(path, name string) (interface{}, error) {
	var config interface{}
	vip := viper.New()
	vip.AddConfigPath(path)   //设置读取的文件路径
	vip.SetConfigName(name)   //设置读取的文件名
	vip.SetConfigType("yaml") //设置文件的类型
	//尝试进行配置读取
	if err := vip.ReadInConfig(); err != nil {
		fmt.Println("尝试进行配置读取:", err)
		return nil, err
	}
	verr := vip.Unmarshal(&config)
	if verr != nil {
		return nil, verr
	}
	return config, nil
}
