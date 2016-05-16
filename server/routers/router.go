package routers

import (
	"github.com/fishedee/web"
	"mymanager/controllers"
)

func init(){
	web.InitRoute("main",&controllers.MainController{})
}
