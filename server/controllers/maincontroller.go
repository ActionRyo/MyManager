package controllers

import (
	"github.com/fishedee/web"
)

type MainController struct{
	web.Controller
}

func (this *MainController)Go_Raw()interface{}{
	return "Hello World"
}

func (this *MainController)AutoRender(data interface{},viewname string){
	this.Ctx.Write([]byte(data.(string)))
}
