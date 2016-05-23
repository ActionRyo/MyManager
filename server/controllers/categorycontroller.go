package controllers

import (
	. "mymanager/models/category"
	. "mymanager/models/condition"
)

type CategoryController struct {
	BaseController
	CategoryAo CategoryAoModel
}

// 查询分类信息
func (this *CategoryController) Search_Json() interface{} {
	var lim Limit
	this.CheckGet(&lim)

	var cat Category
	this.CheckGet(&cat)

	cs := this.CategoryAo.QueryCategoryByCondition(cat, lim)

	return cs
}

// 根据分类ID查询分类信息
func (this *CategoryController) Get_Json() interface{} {
	var one Category
	this.CheckGet(&one)

	cs := this.CategoryAo.QueryCategoryByCategoryId(one.CategoryId)

	return cs[0]
}

// 添加分类信息
func (this *CategoryController) Add_Json() interface{} {
	var input struct {
		Name   string
		Remark string
	}

	this.CheckPost(&input)

	this.CategoryAo.Add(input.Name, input.Remark)

	return nil
}

// 修改分类信息
func (this *CategoryController) Mod_Json() /*interface{} */ {
	var oneCategory Category
	this.CheckPost(&oneCategory)

	// 验证Session
	// to do

	this.CategoryAo.Update(oneCategory)

	//return nil
}

// 查询分类信息
func (this *CategoryController) Del_Json() interface{} {
	var input struct {
		CategoryId int
	}
	this.CheckPost(&input)

	this.CategoryAo.Delete(input.CategoryId)

	return nil
}
