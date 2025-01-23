package common

/**
* 对外api接口
 */
import (
	"encoding/json"
	"gofly/utils/gf"
	"gofly/utils/tools/gconv"

	"github.com/gin-gonic/gin"
)

type Api struct{}

func init() {
	fpath := Api{}
	gf.Register(&fpath, fpath)
}

// 生成公式对应的 SVG 代码
// 这里可以自己使用mathjax 生成 SVG 代码（https://www.mathjax.org)
func (api *Api) Latex(c *gin.Context) {
	params, _ := gf.RequestParam(c)
	ref := gf.Post("https://drawing.aomao.com/api/latex", params, "application/json")
	var parameter gf.Map
	if err := json.Unmarshal([]byte(ref), &parameter); err == nil {
		if gconv.Bool(parameter["success"]) {
			gf.Success().SetMsg("生成公式对应的SVG").SetData(parameter["svg"]).SetExdata(params).Regin(c)
		} else {
			gf.Failed().SetMsg("生成公式对应的SVG失败").Regin(c)
		}
	}
}
