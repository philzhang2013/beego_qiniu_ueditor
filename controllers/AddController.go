package controllers

import (
	"github.com/astaxie/beego"
)

type AddController struct {
	beego.Controller
}

func (controller *AddController) Get() {

	controller.TplName = "add.tpl"

}

func (controller *AddController) Post() {

	//发布

}
