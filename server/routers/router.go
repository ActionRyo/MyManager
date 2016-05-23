package routers

import (
	"github.com/fishedee/web"
	"mymanager/controllers"
)

func init() {
	web.InitRoute("login", &controllers.MainController{})
	web.InitRoute("user", &controllers.UserController{})
	web.InitRoute("card", &controllers.CardController{})
	web.InitRoute("category", &controllers.CategoryController{})
	web.InitRoute("account", &controllers.AccountController{})
}
