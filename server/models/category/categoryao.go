package category

import (
	. "github.com/fishedee/language"
	"github.com/fishedee/web"
	"mymanager/models/condition"
	. "mymanager/models/user"
	"time"
)

type CategoryAoModel struct {
	web.Model
	CategoryDb CategoryDbModel
	UserAo     UserAoModel
}

// 根据条件获取分类信息
func (this *CategoryAoModel) QueryCategoryByCondition(data Category, lim condition.Limit) Categorys {

	userid := this.UserAo.GetUserIdFromSession()

	data.UserId = userid
	datas := this.CategoryDb.QueryCategoryByCondition(data, lim)
	this.Log.Debug("%v", lim)
	this.Log.Debug("%v", datas.Data)

	return datas
}

// 根据分类ID查询分类信息
func (this *CategoryAoModel) QueryCategoryByCategoryId(cid int) []Category {
	//userid := this.UserAo.GetUserIdFromSession()
	datas := this.CategoryDb.QueryCategoryByCategoryId(cid)

	return datas
}

// 根据用户ID查询分类信息
func (this *CategoryAoModel) QueryCategoryByUserid(userid int) []Category {
	datas := this.CategoryDb.QueryCategoryByUserid(userid)

	return datas
}

// 添加分类
func (this *CategoryAoModel) Add(name string, rem string) {
	//this.IsAdmin()
	userid := this.UserAo.GetUserIdFromSession()
	if userid == 0 {
		Throw(1, "获取用户信息失败")
	}

	if name == "" {
		Throw(1, "请输入分类名称")
	}

	if rem == "" {
		Throw(1, "请输入备注信息")
	}

	// 验证是否已经存在
	datas := this.CategoryDb.QueryCategoryByName(name, userid)
	if len(datas) > 0 {
		Throw(1, "已经存在相同分类名称")
	}

	var cat Category

	cat.Name = name
	cat.Remark = rem
	cat.UserId = userid
	cat.CreateTime = time.Now()
	cat.ModifyTime = time.Now()

	this.CategoryDb.Add(cat)
}

// 修改分类信息
func (this *CategoryAoModel) Update(aCategory Category) {
	userid := this.UserAo.GetUserIdFromSession()

	if userid == 0 {
		Throw(1, "获取用户信息失败")
	}

	if aCategory.Name == "" {
		Throw(1, "请输入分类名称")
	}

	if aCategory.Remark == "" {
		Throw(1, "请输入备注信息")
	}

	// 验证是否已经存在
	datas := this.CategoryDb.QueryCategoryByNameForUpd(aCategory.Name, aCategory.CategoryId, userid)
	if len(datas) > 0 {
		Throw(1, "已经存在相同分类名称")
	}

	aCategory.UserId = userid
	aCategory.ModifyTime = time.Now()

	this.CategoryDb.Update(aCategory)
}

// 删除分类信息
func (this *CategoryAoModel) Delete(Categoryid int) {
	this.UserAo.IsAdmin()
	this.CategoryDb.Delete(Categoryid)
}
