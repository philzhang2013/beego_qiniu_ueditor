package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (controller *IndexController) Get() {

	fmt.Println("hahaha***")
	controller.Redirect("/add", 302)

}
