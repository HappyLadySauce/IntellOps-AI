package developer

import (
	"bufio"
	"errors"
	"fmt"
	"gofly/utils/tools/gcfg"
	"gofly/utils/tools/gconv"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var (
	dbConf, _  = gcfg.Instance().Get(ctx, "database.default")
	dbConf_arr = gconv.Map(dbConf)
)

// 1.1 检查该类是否添加到控制器，参数：modelname控制器模块名、path添加模块、haseMoleCtr模块是否存在控制器
func CheckIsAddController(modelname, path string, haseMoleCtr bool) error {
	filePath := filepath.Join("app/", modelname, "/controller.go")
	//1判断文件没有则添加
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			if modelname == "" {
				return errors.New("app下的controller.go文件不存在")
			}
			//模块控制器没可以自动创建
			os.Create(filePath)
			//复制文件
			err := CopyFileContents("devsource/developer/codetpl/go/controller.gos", filePath)
			if err != nil {
				return err
			}
		}
	}
	con_path := "gofly/app/" + path
	//打开controller.go控制文件
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	var result = ""
	ishase := false
	for {
		a, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		//对非根目录控制器文件包名进行替换成改模块名字
		if strings.Contains(string(a), "package controller") && modelname != "" {
			datestr := strings.ReplaceAll(string(a), "package controller", "package "+modelname)
			result += datestr + "\n"
		} else {
			result += string(a) + "\n"
		}
		//判断控制器内容是否存在要引入的模块
		if strings.Contains(string(a), con_path) {
			ishase = true
		}
	}
	if !ishase {
		if modelname == "" && haseMoleCtr { //app根控制器器模块下存在控制器-根控制器
			//1.引入模块控制器
			addstr := "	\"" + con_path + "\"\n"
			addstr += ")"
			result = strings.Replace(result, ")", addstr, 1)
			//2添加路由钩子
			addrouterstr := fmt.Sprintf("	%v.RouterHandler(c, \"%v\")\n", path, path)
			addrouterstr += "}"
			result = strings.Replace(result, "}", addrouterstr, 1)
		} else { //处理模块控制器
			addstr := "	_ \"" + con_path + "\"\n"
			addstr += ")"
			result = strings.Replace(result, ")", addstr, 1)
		}
	}

	fw, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //os.O_TRUNC清空文件重新写入，否则原文件内容可能残留
	w := bufio.NewWriter(fw)
	w.WriteString(result)
	if err != nil {
		return err
	}
	w.Flush()
	fw.Close()
	return nil
}

// 1.2 存在控制器则移除
func CheckApiRemoveController(modelname, path string) {
	filePath := filepath.Join("app/", modelname, "/controller.go")
	if _, err := os.Stat(filePath); os.IsNotExist(err) { //不存在
		return
	}
	con_path := "gofly/app/" + path
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	var result = ""
	for {
		a, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		if strings.Contains(string(a), con_path) || strings.Contains(string(a), fmt.Sprintf("%v.RouterHandler(c, \"%v\")", path, path)) { //存在路由则移除
			continue
		} else {
			result += string(a) + "\n"
		}
	}
	fw, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //os.O_TRUNC清空文件重新写入，否则原文件内容可能残留
	w := bufio.NewWriter(fw)
	w.WriteString(result)
	if err != nil {
		panic(err)
	}
	w.Flush()
	fw.Close()
}

// 单个文件复制
// 将 src 的文件内容拷贝到了 dst 里面
func CopyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}

// 整个文件复制
// 复制整个文件夹下文件到另个文件夹-targetPath复制的目标文件目录，destPath复制到新的目录
func CopyAllDir(targetPath string, destPath string) error {
	err := filepath.Walk(targetPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		destPath := filepath.Join(destPath, path[len(targetPath):])
		//如果是个文件夹则创建这个文件夹
		if info.IsDir() {
			return os.MkdirAll(destPath, info.Mode())
		}
		//如果是文件则生成这个文件
		return CopyFile(path, destPath)
	})
	return err
}

// 复制单个文件
func CopyFile(srcFile, destFile string) error {
	src, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer src.Close()
	//创建复制的文件
	dest, err := os.Create(destFile)
	if err != nil {
		return err
	}
	defer dest.Close()
	//复制内容到文件
	_, err = io.Copy(dest, src)
	if err != nil {
		return err
	}
	//让复制的文件将内容存到硬盘而不是缓存
	err = dest.Sync()
	if err != nil {
		return err
	}
	return nil
}
