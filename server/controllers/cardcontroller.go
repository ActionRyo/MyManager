package controllers

import (
	. "mymanager/models/card"
	. "mymanager/models/condition"
)

type CardController struct {
	BaseController
	CardAo CardAoModel
}

// 查询银行卡信息
func (this *CardController) Search_Json() interface{} {
	var lim Limit
	this.CheckGet(&lim)

	var onecard Card
	this.CheckGet(&onecard)

	cards := this.CardAo.QueryCardByCondition(onecard, lim)

	return cards
}

// 根据银行卡ID查询银行卡信息
func (this *CardController) Get_Json() interface{} {
	var one Card
	this.CheckGet(&one)

	cards := this.CardAo.QueryCardByCardId(one.CardId)

	return cards[0]
}

// 添加银行卡信息
func (this *CardController) Add_Json() interface{} {
	var one Card

	this.CheckPost(&one)

	this.CardAo.Add(one)

	return nil
}

// 修改银行卡信息
func (this *CardController) Mod_Json() /*interface{} */ {
	var oneCard Card
	this.CheckPost(&oneCard)

	// 验证Session
	// to do

	this.CardAo.Update(oneCard)

	//return nil
}

// 删除银行卡信息
func (this *CardController) Del_Json() interface{} {
	var input struct {
		CardId int
	}
	this.CheckPost(&input)

	this.CardAo.Delete(input.CardId)

	return nil
}
