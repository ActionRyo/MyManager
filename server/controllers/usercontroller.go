package controllers

import (
	. "mymanager/models/condition"
	. "mymanager/models/user"
)

type UserController struct {
	BaseController
	UserAo UserAoModel
}

//type Passwords struct {
//	OldPassword string
//	NewPassword string
//}

// 添加用户
func (this *UserController) Add_Json() interface{} {
	var oneuser User

	this.CheckPost(&oneuser)

	this.UserAo.Add(oneuser.Name, oneuser.Password, oneuser.Type)

	return nil
}

// 查询用户信息
func (this *UserController) Search_Json() interface{} {
	var lim Limit
	this.CheckGet(&lim)

	var oneuser User
	this.CheckGet(&oneuser)

	users := this.UserAo.QueryUserByCondition(oneuser, lim)

	return users
}

// 根据用户ID查询用户信息
func (this *UserController) Get_Json() interface{} {
	var oneuser User
	this.CheckGet(&oneuser)

	users := this.UserAo.QueryUserByUserId(oneuser.UserId)

	return users[0]
}

// 修改用户密码
func (this *UserController) ModPassword_Json() /*interface{} */ {
	var oneuser User
	this.CheckPost(&oneuser)

	// 验证Session
	// to do

	this.UserAo.UpdatePassword(oneuser)

	//return nil
}

// 修改用户角色
func (this *UserController) ModType_Json() /*interface{} */ {
	var oneuser User
	this.CheckPost(&oneuser)

	this.UserAo.UpdateRole(oneuser)
}

// 修改自身密码
func (this *UserController) ModMyPassword_Json() /*interface{} */ {
	var input struct {
		OldPassword string
		NewPassword string
	}

	this.CheckPost(&input)

	//fmt.Println("loginUser", loginUser)
	//fmt.Printf("#+v", input)

	this.UserAo.UpdateSelfPwd(input.OldPassword, input.NewPassword)
}

// 查询用户信息
func (this *UserController) Del_Json() interface{} {
	var input struct {
		UserId int
	}
	this.CheckPost(&input)

	this.UserAo.Delete(input.UserId)

	return nil
}
