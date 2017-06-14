package routers

import (
	"beego_qiniu_ueditor/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.IndexController{})
	beego.Router("/add", &controllers.AddController{})
	beego.Router("/controller", &controllers.UEController{})

}
