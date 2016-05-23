package controllers

import (
	. "mymanager/models/account"
	. "mymanager/models/condition"
)

type AccountController struct {
	BaseController
	AccountAo AccountAoModel
}

// 查询分类信息
func (this *AccountController) Search_Json() interface{} {
	var lim Limit
	this.CheckGet(&lim)

	var cat Account
	this.CheckGet(&cat)

	cs := this.AccountAo.QueryAccountByCondition(cat, lim)

	return cs
}

// 根据分类ID查询分类信息
func (this *AccountController) Get_Json() interface{} {
	var one Account
	this.CheckGet(&one)

	cs := this.AccountAo.QueryAccountByAccountId(one.AccountId)

	return cs[0]
}

// 添加分类信息
func (this *AccountController) Add_Json() interface{} {
	var acco Account

	this.CheckPost(&acco)

	this.AccountAo.Add(acco)

	return nil
}

// 修改分类信息
func (this *AccountController) Mod_Json() /*interface{} */ {
	var oneCategory Account
	this.CheckPost(&oneCategory)

	// 验证Session
	// to do

	this.AccountAo.Update(oneCategory)

	//return nil
}

// 分类信息
func (this *AccountController) Del_Json() interface{} {
	var input struct {
		AccountId int
	}
	this.CheckPost(&input)

	this.AccountAo.Delete(input.AccountId)

	return nil
}

// 按周数据的收支情况
func (this *AccountController) GetWeekTypeStatistic_Json() interface{} {
	var input struct {
		UserId int
	}

	this.CheckGet(&input)
	datas := this.AccountAo.GetWeekTypeStatistic(input.UserId)
	return datas
}

// 详细的收支情况
func (this *AccountController) GetWeekDetailTypeStatistic_Json() interface{} {
	var input struct {
		Year int
		Week int
		Type int
	}

	this.CheckGet(&input)

	datas := this.AccountAo.GetWeekDetailTypeStatistic(input.Year, input.Week, input.Type)

	return datas
}

// 统计银行卡余额
func (this *AccountController) GetWeekCardStatistic_Json() interface{} {

	datas := this.AccountAo.GetWeekCardStatistic()

	return datas
}

// 详细银行卡余额
func (this *AccountController) GetWeekDetailCardStatistic_Json() interface{} {
	var input struct {
		Year   int
		Week   int
		CardId int
	}

	this.CheckGet(&input)

	datas := this.AccountAo.GetWeekDetailCardStatistic(input.Year, input.Week, input.CardId)

	return datas
}
