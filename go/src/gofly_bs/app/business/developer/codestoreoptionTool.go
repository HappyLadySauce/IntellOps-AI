package developer

import (
	"archive/zip"
	"bufio"
	"fmt"
	"gofly/utils/gf"
	"gofly/utils/tools/gconv"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

// 获取文文件夹下的文件及文件goApp
func GetAllFileApp(pathname string) (string, []string, error) {
	rd, err := os.ReadDir(pathname)
	var folders = make([]string, 0)
	if err != nil {
		return "", folders, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir := pathname + "/" + fi.Name()
			sec_rd, _ := os.ReadDir(fullDir)
			if len(sec_rd) > 0 {
				hase_dir := false
				for _, sec_fi := range sec_rd {
					if sec_fi.IsDir() {
						folders = append(folders, fi.Name()+"/"+sec_fi.Name())
						hase_dir = true
					}
				}
				if hase_dir == false {
					folders = append(folders, fi.Name())
				}
			} else {
				folders = append(folders, fi.Name())
			}
		}
	}
	return strings.Join(folders, ","), folders, nil
}

// 获取go文件目录-后端打包用的
func GetAppFileNoAdmin(pathname string) (string, []string, error) {
	rd, err := os.ReadDir(pathname)
	var folders = make([]string, 0)
	if err != nil {
		return "", folders, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir := pathname + "/" + fi.Name()
			sec_rd, _ := os.ReadDir(fullDir)
			if len(sec_rd) > 0 {
				folders = append(folders, fi.Name())
				for _, sec_fi := range sec_rd {
					if sec_fi.IsDir() {
						folders = append(folders, fi.Name()+"/"+sec_fi.Name())
					}
				}
			} else {
				folders = append(folders, fi.Name())
			}
		}
	}
	return strings.Join(folders, ","), folders, nil
}

// 获取文文件夹下的文件及文件
func GetAllFile(pathname string) (map[string]interface{}, error) {
	rd, err := os.ReadDir(pathname)
	if err != nil {
		return map[string]interface{}{"folders": "", "files": ""}, err
	}
	var folders = make([]string, 0)
	var files = make([]string, 0)
	for _, fi := range rd {
		if fi.IsDir() {
			folders = append(folders, fi.Name())
		} else {
			files = append(files, fi.Name())
		}
	}
	var folders_str = ""
	if len(folders) > 0 {
		folders_str = strings.Join(folders, ",")
	}
	var files_str = ""
	if len(files) > 0 {
		files_str = strings.Join(files, ",")
	}
	return map[string]interface{}{"folders": folders_str, "files": files_str}, nil
}

// 获取文文件夹下是否存在目录
func GetDirHasefolder(pathname string) (bool, error) {
	var folders bool = false
	rd, err := os.ReadDir(pathname)
	if err != nil {
		return folders, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			folders = true
		}
	}
	return folders, nil
}

// 更新配置文件
func UpConfFieldData(path string, parameter map[string]interface{}) error {
	file_path := filepath.Join(path, "config.yml")
	f, err := os.Open(file_path)
	if err != nil {
		return err
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	var result = ""
	var is_hose = false
	for {
		is_hose = false
		a, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		for keys, Val := range parameter {
			if strings.Contains(string(a), keys) {
				is_hose = true
				datestr := strings.ReplaceAll(string(a), string(a), fmt.Sprintf("     %v: %v\n", keys, Val))
				result += datestr
			}
		}
		if !is_hose {
			result += string(a) + "\n"
		}
	}
	fw, err := os.OpenFile(file_path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //os.O_TRUNC清空文件重新写入，否则原文件内容可能残留
	w := bufio.NewWriter(fw)
	w.WriteString(result)
	if err != nil {
		panic(err)
	}
	w.Flush()
	fw.Close()
	return nil
}

// 导入菜单数据、并返回插入数据id
func Insertmenu(userID int64, data interface{}, pid interface{}, tablename string) []string {
	var menuids = make([]string, 0)
	for _, menuitem := range data.([]interface{}) {
		menuitem_obj := menuitem.(map[string]interface{})
		menuitem_obj["pid"] = pid
		menuitem_obj["uid"] = userID
		delete(menuitem_obj, "id")
		parent_id, _ := gf.Model(tablename).Where("pid", 0).Where("routename", menuitem_obj["routename"]).Value("id")
		if parent_id == nil {
			if _, ok := menuitem_obj["children"]; ok {
				subdata := menuitem_obj["children"]
				delete(menuitem_obj, "children")
				nemuid, _ := gf.Model(tablename).Data(menuitem_obj).InsertAndGetId()
				gf.Model(tablename).Where("id", nemuid).Data(map[string]interface{}{"weigh": nemuid}).Update()
				menuids = append(menuids, strconv.FormatInt(nemuid, 10))
				if subdata != nil {
					m_menuids := Insertmenu(userID, subdata, nemuid, tablename)
					menuids = append(menuids, m_menuids...)
				}
			} else {
				nemuid, _ := gf.Model(tablename).Data(menuitem_obj).InsertAndGetId()
				gf.Model(tablename).Where("id", nemuid).Data(map[string]interface{}{"weigh": nemuid}).Update()
				menuids = append(menuids, strconv.FormatInt(nemuid, 10))
			}
		} else {
			subdata := menuitem_obj["children"]
			if subdata != nil {
				m_menuids := Insertmenu(userID, subdata, parent_id, tablename)
				menuids = append(menuids, m_menuids...)
			}
		}
	}
	return menuids
}

// 下载文件到指定目录 url下载地址 downdir 下载到位置及文件名称
func DownFileToDir(url, downdir string) (bool, string) {
	// 获取网络文件的数据流
	resp, err := http.Get(url)
	if err != nil {
		// 处理错误
		return false, ""
	}
	defer resp.Body.Close()
	// 读取数据流到内存中
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, ""
	}
	// 将文件写入到指定目录
	err = os.WriteFile(downdir, data, 0644)
	if err != nil {
		return false, ""
	}
	return true, downdir
}

// Unzip decompresses a zip file to specified directory.
// Note that the destination directory don't need to specify the trailing path separator.
// If the destination directory doesn't exist, it will be created automatically.
func Unzip(zipath, dir string) error {
	// Open zip file.
	reader, err := zip.OpenReader(zipath)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, file := range reader.File {
		if err := unzipFile(file, dir); err != nil {
			return err
		}
	}
	return nil
}

func unzipFile(file *zip.File, dir string) error {
	// Prevent path traversal vulnerability.
	// Such as if the file name is "../../../path/to/file.txt" which will be cleaned to "path/to/file.txt".
	name := strings.TrimPrefix(filepath.Join(string(filepath.Separator), file.Name), string(filepath.Separator))
	filePath := path.Join(dir, name)
	// Create the directory of file.
	if file.FileInfo().IsDir() {
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
			return err
		}
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}
	// Open the file.
	r, err := file.Open()
	if err != nil {
		return err
	}
	defer r.Close()
	// Create the file.
	w, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer w.Close()
	// Save the decompressed file content.
	_, err = io.Copy(w, r)
	return err
}

// 添加模块验证和删除验证
func AddOrRemoveValidity(vApi, dotype, path string, appConf_arr map[string]interface{}) {
	noVerifyAPIRoot := gconv.String(appConf_arr["noVerifyAPIRoot"])
	var upAppconf gf.Map = make(map[string]interface{}, 0)
	vApi_arr := strings.Split(vApi, ",")
	if dotype == "install" { //安装
		api_arr := strings.Split(noVerifyAPIRoot, ",")
		var ContainsApiName []string
		for _, val := range vApi_arr {
			if !gf.StrInArray(val, api_arr) && val != "" {
				ContainsApiName = append(ContainsApiName, val)
			}
		}
		if len(ContainsApiName) > 0 {
			//添加api配置
			upAppconf["noVerifyAPIRoot"] = fmt.Sprintf("%v,%v", noVerifyAPIRoot, strings.Join(ContainsApiName, ","))
		}
	} else { //卸载
		if vApi != "" {
			//删除api配置
			api_arr := strings.Split(noVerifyAPIRoot, ",")
			var new_api_config []string
			for _, val := range api_arr {
				if !gf.StrInArray(val, vApi_arr) && val != "" {
					new_api_config = append(new_api_config, val)
				}
			}
			upAppconf["noVerifyAPIRoot"] = strings.Join(new_api_config, ",")
		}
	}
	gf.UpConfFieldData(path, upAppconf, "  ")
}
