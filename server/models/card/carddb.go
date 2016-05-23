package card

import (
	"github.com/fishedee/web"
	"mymanager/models/condition"
)

type CardDbModel struct {
	web.Model
}

// 获取银行卡信息
func (this *CardDbModel) QueryCardByCondition(data Card, lim condition.Limit) Cards {

	datas := []Card{}

	db := this.DB.NewSession()
	db.And("userid=?", data.UserId)

	if data.Name != "" {
		db = db.And("name like ?", "%"+data.Name+"%")
	}

	if data.Remark != "" {
		db = db.And("remark like ?", "%"+data.Remark+"%")
	}

	if lim.PageSize != 0 {
		db = db.Limit(lim.PageSize, lim.PageIndex)
	}

	err := db.Desc("CardId").Find(&datas)
	if err != nil {
		panic(err)
	}

	dbs := this.DB.NewSession()
	dbs.And("userid=?", data.UserId)

	if data.Name != "" {
		dbs = dbs.And("name like ?", "%"+data.Name+"%")
	}

	if data.Remark != "" {
		dbs = dbs.And("remark like ?", "%"+data.Remark+"%")
	}

	total, errs := dbs.Count(&Card{})
	if errs != nil {
		panic(errs)
	}

	return Cards{
		Count: int(total),
		Data:  datas,
	}
}

// 根据银行卡ID查询银行卡信息
func (this *CardDbModel) QueryCardByCardId(cardid int, userid int) []Card {
	data := []Card{}
	db := this.DB.NewSession()

	err := db.And("cardid=?", cardid).And("userid=?", userid).Find(&data)
	if err != nil {
		panic(err)
	}

	return data
}

// 获取银行卡信息
func (this *CardDbModel) QueryCardByCardNum(cardNum string) []Card {
	data := []Card{}

	err := this.DB.Where("card=?", cardNum).Find(&data)
	if err != nil {
		panic(err)
	}

	return data
}

// 获取银行卡信息
func (this *CardDbModel) QueryCardByCardNumAndCardId(cardNum string, cid int) []Card {
	data := []Card{}
	db := this.DB.NewSession()
	err := db.And("cardid<>?", cid).And("card=?", cardNum).Find(&data)
	if err != nil {
		panic(err)
	}

	return data
}

// 根据用户ID获取所有的银行卡信息
func (this *CardDbModel) QueryCardByUserId(userid int) Cards {
	data := []Card{}
	err := this.DB.Where("userid=?", userid).Find(&data)
	if err != nil {
		panic(err)
	}

	return Cards{
		Count: 0,
		Data:  data,
	}
}

// 添加银行卡信息
func (this *CardDbModel) Add(data Card) {
	_, err := this.DB.Insert(data)
	if err != nil {
		panic(err)
	}
}

// 修改银行卡信息
func (this *CardDbModel) Update(data Card) {
	_, err := this.DB.Where("cardid=?", data.CardId).Update(&data)
	if err != nil {
		panic(err)
	}
}

// 删除银行卡信息
func (this *CardDbModel) Delete(cardid int) {
	data := Card{}
	_, err := this.DB.Where("cardid=?", cardid).Delete(&data)
	if err != nil {
		panic(err)
	}
}
