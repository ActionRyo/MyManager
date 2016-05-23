package controllers

import (
	"mymanager/models/user"
)

type MainController struct {
	BaseController
	UserAo user.UserAoModel
}

// 验证用户是否登陆
func (this *MainController) IsLogin_Json() interface{} {

	this.UserAo.CheckMustLogin()

	return nil
}

// 用户登陆
func (this *MainController) Checkin_Json() interface{} {
	var input struct {
		Name     string
		Password string
	}

	this.CheckPost(&input)

	this.UserAo.Login(input.Name, input.Password)

	return nil
}

//用户注销登陆
func (this *MainController) Checkout_Json() interface{} {

	this.UserAo.Logout()

	return nil
}
