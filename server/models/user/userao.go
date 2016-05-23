package user

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	. "github.com/fishedee/language"
	"github.com/fishedee/web"
	"mymanager/models/condition"
	"strconv"
	"time"
)

type UserAoModel struct {
	web.Model
	UserDb UserDbModel
}

//用户登陆
func (this *UserAoModel) Login(name string, password string) {
	if name == "" || password == "" {
		Throw(1, "请输入账户或密码")
	}

	safePwd := this.SafePwd(password)

	users := this.UserDb.QueryUserByNameAndPassword(name, safePwd)

	if len(users) == 0 {
		Throw(1, "账户或密码错误")
	}

	sess, err := this.Session.SessionStart()
	if err != nil {
		panic(err)
	}

	sess.Set("UserId", users[0].UserId)
	defer sess.SessionRelease()
}

func (this *UserAoModel) Logout() {
	sess, err := this.Session.SessionStart()

	if err != nil {
		panic(err)
	}

	sess.Delete("UserId")
	defer sess.SessionRelease()
}

// 验证Session是否过期
func (this *UserAoModel) CheckMustLogin() {
	sess, err := this.Session.SessionStart()
	if err != nil {
		Throw(10001, "调用Session失败")
	}

	defer sess.SessionRelease()
	userid := sess.Get("UserId")

	clientIdString := fmt.Sprintf("%v", userid)
	clientIdInt, err := strconv.Atoi(clientIdString)
	if err != nil {
		clientIdInt = 0
	}

	if clientIdInt == 0 {
		Throw(10001, "用户没有登陆")
	}
}

// 判断当前角色是否是管理员
func (this *UserAoModel) IsAdmin() {
	uid := this.GetUserIdFromSession()
	uses := this.QueryUserByUserId(uid)
	if len(uses) == 0 || uses[0].UserId == 0 {
		Throw(1, "您还没有登陆")
	}

	if uses[0].Type != 1 {
		Throw(1, "您没有相关权限")
	}
}

//  根据条件查询用户信息
func (this *UserAoModel) QueryUserByCondition(data User, lim condition.Limit) Users {
	this.IsAdmin()
	datas := this.UserDb.QueryUserByCondition(data, lim)

	return datas
}

// 根据用户ID查询用户信息
func (this *UserAoModel) QueryUserByUserId(userid int) []User {
	datas := this.UserDb.QueryUserByUserId(userid)

	return datas
}

// 添加用户
func (this *UserAoModel) Add(name string, password string, t int) {
	this.IsAdmin()
	this.CheckMustLogin()
	if name == "" || password == "" {
		Throw(1, "请输入账户或密码")
	}

	users := this.UserDb.QueryUserByName(name)
	if len(users) != 0 {
		Throw(1, "存在重复的用户名")
	}

	safePwd := this.SafePwd(password)

	this.UserDb.Add(User{
		Name:       name,
		Password:   safePwd,
		Type:       t,
		CreateTime: time.Now(),
		ModifyTime: time.Now(),
	})
}

// 修改用户密码
func (this *UserAoModel) UpdatePassword(data User) {
	this.IsAdmin()

	if data.UserId == 0 {
		Throw(1, "不存在该用户")
	}

	if data.Password == "" {
		Throw(1, "密码不能为空")
	}

	oneuser := this.QueryUserByUserId(data.UserId)
	if len(oneuser) == 0 {
		Throw(1, "不存在该用户")
	}

	data.Password = this.SafePwd(data.Password)
	data.ModifyTime = time.Now()

	this.UserDb.Update(data)
}

// 修改用户权限
func (this *UserAoModel) UpdateRole(data User) {
	this.IsAdmin()
	if data.UserId == 0 {
		Throw(1, "不存在该用户")
	}

	oneuser := this.QueryUserByUserId(data.UserId)
	if len(oneuser) == 0 {
		Throw(1, "不存在该用户")
	}

	this.UserDb.Update(data)
}

// 修改自身密码
func (this *UserAoModel) UpdateSelfPwd(oldPwd string, newPwd string) {
	if oldPwd == "" || newPwd == "" {
		Throw(1, "请输入旧密码和新密码")
	}

	// 验证旧密码是否正确
	soldPwd := this.SafePwd(oldPwd)
	userid := this.GetUserIdFromSession()
	self := this.QueryUserByUserId(userid)
	if len(self) == 0 || self[0].UserId == 0 {
		Throw(1, "获取您的信息失败")
	}

	if soldPwd != self[0].Password {
		Throw(1, "您输入的旧密码不正确")
	}

	// 保存修改后的信息
	sfNewPwd := this.SafePwd(newPwd)
	self[0].Password = sfNewPwd
	self[0].ModifyTime = time.Now()

	this.UserDb.Update(self[0])
}

// 删除用户信息
func (this *UserAoModel) Delete(userid int) {
	this.IsAdmin()
	this.UserDb.Delete(userid)
}

// 字符串加密
func (this *UserAoModel) SafePwd(password string) string {
	sPwd := sha1.New()
	sPwd.Write([]byte(password))
	re := sPwd.Sum(nil)
	re2 := hex.EncodeToString(re)

	return re2
}

// 获取session中的用户ID
func (this *UserAoModel) GetUserIdFromSession() int {

	sess, err := this.Session.SessionStart()
	if err != nil {
		Throw(10001, "调用Session失败")
	}
	defer sess.SessionRelease()

	userid := sess.Get("UserId")

	clientIdString := fmt.Sprintf("%v", userid)
	clientIdInt, errs := strconv.Atoi(clientIdString)
	if errs != nil {
		panic(errs)
	}

	return clientIdInt
}

// 统计用户数量
//func (this *UserAoModel) Count() int {
//	total := this.UserDb.UserCount()
//	return total
//}
