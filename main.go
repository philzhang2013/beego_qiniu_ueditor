package main

import (
	_ "beego_qiniu_ueditor/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"os"
)

func main() {

	_, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		fmt.Printf("NewConfig failed. path[%s], error:%s\n\n", "conf/app.conf", err)
		os.Exit(1)
	}

	beego.Run()

}
