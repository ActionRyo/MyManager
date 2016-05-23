package user

import (
	"github.com/fishedee/web"
	"mymanager/models/condition"
)

type UserDbModel struct {
	web.Model
}

// 根据账户名和密码查询用户信息
func (this *UserDbModel) QueryUserByNameAndPassword(name string, pwd string) []User {
	data := []User{}

	var err error
	err = this.DB.Where("name=?", name).Where("password=?", pwd).Find(&data)

	if err != nil {
		panic(err)
	}

	return data
}

// 根据用户ID查询用户信息
func (this *UserDbModel) QueryUserByUserId(userid int) []User {
	data := []User{}

	err := this.DB.Where("userid=?", userid).Find(&data)
	if err != nil {
		panic(err)
	}

	return data
}

// 根据用户名获取用户信息
func (this *UserDbModel) QueryUserByName(name string) []User {
	data := []User{}

	err := this.DB.Where("name=?", name).Find(&data)
	if err != nil {
		panic(err)
	}

	return data
}

// 获取用户数量、用户信息
func (this *UserDbModel) QueryUserByCondition(data User, lim condition.Limit) Users {

	datas := []User{}

	db := this.DB.NewSession()

	if data.Name != "" {
		db = db.And("name like ?", "%"+data.Name+"%")
	}

	if data.Type != 0 {
		db = db.And("type=?", data.Type)
	}

	err := db.Desc("userid").Limit(lim.PageSize, lim.PageIndex).Find(&datas)
	if err != nil {
		panic(err)
	}

	dbs := this.DB.NewSession()

	if data.Name != "" {
		dbs = dbs.And("name like ?", "%"+data.Name+"%")
	}

	if data.Type != 0 {
		dbs = dbs.And("type=?", data.Name)
	}

	total, errs := dbs.Count(&User{})
	if errs != nil {
		panic(errs)
	}

	return Users{
		Count: int(total),
		Data:  datas,
	}
}

// 添加用户
func (this *UserDbModel) Add(data User) {
	_, err := this.DB.Insert(data)
	if err != nil {
		panic(err)
	}
}

// 修改用户信息
func (this *UserDbModel) Update(data User) {
	var err error
	_, err = this.DB.Where("userid=?", data.UserId).Update(&data)
	if err != nil {
		panic(err)
	}
}

// 删除用户信息
func (this *UserDbModel) Delete(useid int) {
	data := User{}
	_, err := this.DB.Where("userid=?", useid).Delete(&data)
	if err != nil {
		panic(err)
	}
}
