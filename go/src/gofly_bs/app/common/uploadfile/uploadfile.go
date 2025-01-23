package uploadfile

import (
	"gofly/utils/gf"
	"os"
	"path/filepath"
	"strings"
)

// 这用来显示图片、附件等资源
type Index struct{ NoNeedLogin []string }

func init() {
	fpath := Index{NoNeedLogin: []string{"*"}}
	gf.Register(&fpath, fpath)
}

// 1.显示uploads目录下的图片
func (api *Index) Getfile(c *gf.GinCtx) {
	fileName := c.Query("url")
	fileurl := strings.Split(fileName, "?")
	if len(fileurl) == 0 || !strings.Contains(fileurl[0], "resource/uploads/") || strings.Contains(fileurl[0], "../") {
		gf.Failed().SetMsg("文件不存在或者禁止访问").Regin(c)
		return
	}
	path, _ := os.Getwd() //获取当前路径
	mfileurl := strings.Split(fileName, "resource/uploads/")
	scfileurl := strings.Split(mfileurl[1], "/")
	if len(scfileurl) == 2 {
		filetype := strings.Split(scfileurl[1], ".")
		if filetype[1] == "zip" || filetype[1] == "rar" {
			_, erByStatFile := os.Stat(filepath.Join(path, fileName))
			if erByStatFile != nil {
				gf.Failed().SetMsg("文件打开失败").SetData(erByStatFile.Error()).Regin(c)
				return
			}
			c.Header("Content-Type", "application/octet-stream")
			c.Header("Content-Disposition", "attachment; filename="+scfileurl[1])
			c.Header("Content-Disposition", "inline;filename="+scfileurl[1])
			c.Header("Content-Transfer-Encoding", "binary")
			c.File(filepath.Join(path, fileName))
		} else {
			c.File(filepath.Join(path, fileName))
		}
	} else if len(scfileurl) == 3 {
		fname := scfileurl[len(scfileurl)-1]
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", "attachment; filename="+fname)
		c.Header("Content-Disposition", "inline;filename="+fname)
		c.Header("Content-Transfer-Encoding", "binary")
		c.File(filepath.Join(path, fileName))
	} else {
		gf.Failed().SetMsg("请用合法方式请求文件").Regin(c)
	}
}

// 2.显示uploads目录下的图片已base64格式返回
func (api *Index) GetImagebase(c *gf.GinCtx) {
	imageName := c.Query("url")
	fileurl := strings.Split(imageName, "?")
	if len(fileurl) > 0 && strings.Contains(fileurl[0], "resource/uploads/") && !strings.Contains(fileurl[0], "../") {
		imageName = fileurl[0]
		path, _ := os.Getwd() //获取当前路径
		mfileurl := strings.Split(imageName, "resource/uploads/")
		scfileurl := strings.Split(mfileurl[1], "/")
		if len(scfileurl) == 2 || len(scfileurl) == 3 {
			file, _ := os.ReadFile(filepath.Join(path, imageName))
			c.Writer.WriteString(string(file))
		} else {
			gf.Failed().SetMsg("请用合法方式请求文件").Regin(c)
		}
	} else {
		gf.Failed().SetMsg("文件禁止访问").Regin(c)
	}
}
