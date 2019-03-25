package index

import (
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
)

type UploadController struct {
	beego.Controller
}


func (this *UploadController) Get() {

}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}


// 文件上传
func (this *UploadController) Post() {
	f, h, err := this.GetFile("uploadname")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	err2 := this.SaveToFile("uploadname",getCurrentDirectory() + "/static/upload/" + string(rand.Intn(100)) + h.Filename );

	fmt.Println(err2)
	this.Ctx.WriteString(h.Filename)
	this.Ctx.WriteString(getCurrentDirectory())
}

