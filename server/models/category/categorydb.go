package category

import (
	"github.com/fishedee/web"
	"mymanager/models/condition"
)

type CategoryDbModel struct {
	web.Model
}

// 获取分类信息
func (this *CategoryDbModel) QueryCategoryByCondition(data Category, lim condition.Limit) Categorys {

	datas := []Category{}

	db := this.DB.NewSession()
	db = db.And("userid=?", data.UserId)
	if data.Name != "" {
		db = db.And("name like ?", "%"+data.Name+"%")
	}

	if data.Remark != "" {
		db = db.And("remark like ?", "%"+data.Remark+"%")
	}

	if lim.PageSize != 0 {
		db = db.Limit(lim.PageSize, lim.PageIndex)
	}

	err := db.Desc("categoryid").Find(&datas)
	if err != nil {
		panic(err)
	}

	dbs := this.DB.NewSession()
	dbs = dbs.And("userid=?", data.UserId)
	if data.Name != "" {
		dbs = dbs.And("name like ?", "%"+data.Name+"%")
	}

	if data.Remark != "" {
		dbs = dbs.And("remark like ?", "%"+data.Remark+"%")
	}

	total, errs := dbs.Count(&Category{})
	if errs != nil {
		panic(errs)
	}

	return Categorys{
		Count: int(total),
		Data:  datas,
	}
}

// 根据分类ID查询分类信息
func (this *CategoryDbModel) QueryCategoryByCategoryId(categoryid int) []Category {
	data := []Category{}
	err := this.DB.Where("Categoryid=?", categoryid).Find(&data)
	if err != nil {
		panic(err)
	}

	return data
}

// 根据分类ID查询分类信息
func (this *CategoryDbModel) QueryCategoryByUserid(userid int) []Category {
	data := []Category{}
	err := this.DB.Where("userid=?", userid).Find(&data)
	if err != nil {
		panic(err)
	}

	return data
}

// 根据分类名称查询分类信息
func (this *CategoryDbModel) QueryCategoryByName(name string, userid int) []Category {
	data := []Category{}
	db := this.DB.NewSession()
	db = db.And("name=?", name)
	db = db.And("userid=?", userid)
	err := db.Find(&data)
	if err != nil {
		panic(err)
	}

	return data
}

// 根据分类名称查询分类信息
func (this *CategoryDbModel) QueryCategoryByNameForUpd(name string, uid int, cid int) []Category {
	data := []Category{}
	db := this.DB.NewSession()
	err := db.And("name=?", name).And("Categoryid<>?", cid).And("userid=?", uid).Find(&data)
	if err != nil {
		panic(err)
	}

	return data
}

// 添加分类信息
func (this *CategoryDbModel) Add(data Category) {
	_, err := this.DB.Insert(data)
	if err != nil {
		panic(err)
	}
}

// 修改分类信息
func (this *CategoryDbModel) Update(data Category) {
	_, err := this.DB.Where("Categoryid=?", data.CategoryId).Update(&data)
	if err != nil {
		panic(err)
	}
}

// 删除分类信息
func (this *CategoryDbModel) Delete(categoryid int) {
	data := Category{}
	_, err := this.DB.Where("Categoryid=?", categoryid).Delete(&data)
	if err != nil {
		panic(err)
	}
}
