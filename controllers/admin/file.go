package admin

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

type FileController struct {
	//baseController
	beego.Controller
}

func (this *FileController) Upload() {
	beego.Debug("here upload")
	_, header, err := this.GetFile("imgFile")
	out := make(map[string]interface{})
	out["error"] = 0
	out["url"] = ""
	if err != nil {
		beego.Error(err)
		out["error"] = 1
	} else {
		savepath := "./static/upload/" + time.Now().Format("20060102")
		if err := os.MkdirAll(savepath, os.ModePerm); err != nil {
			out["error"] = 2
		} else {
			ext := strings.ToLower(header.Filename[strings.LastIndex(header.Filename, "."):])
			filename := fmt.Sprintf("%s/%d%s", savepath, time.Now().UnixNano(), ext)
			if err := this.SaveToFile("imgFile", filename); err != nil {
				out["error"] = 3
			} else {
				out["url"] = "/dl/" + filename[9:]
			}
		}
	}

	beego.Debug(out)

	this.Data["json"] = out
	this.ServeJson()
}

func (this *FileController) Del() {
	//todo
}
