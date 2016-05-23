package card

import (
	. "github.com/fishedee/language"
	"github.com/fishedee/web"
	"mymanager/models/condition"
	. "mymanager/models/user"
	"time"
)

type CardAoModel struct {
	web.Model
	CardDb CardDbModel
	UserAo UserAoModel
}

// 根据条件获取银行卡信息
func (this *CardAoModel) QueryCardByCondition(data Card, lim condition.Limit) Cards {

	userid := this.UserAo.GetUserIdFromSession()
	data.UserId = userid
	datas := this.CardDb.QueryCardByCondition(data, lim)

	return datas
}

// 根据银行卡ID查询银行卡信息
func (this *CardAoModel) QueryCardByCardId(cardid int) []Card {
	userid := this.UserAo.GetUserIdFromSession()
	datas := this.CardDb.QueryCardByCardId(cardid, userid)

	return datas
}

// 根据银行卡ID查询银行卡信息
func (this *CardAoModel) QueryCardBySingleCardId(cardid int) Card {
	datas := this.CardDb.QueryCardByCardId(cardid, 10001)
	if len(datas) == 0 {
		Throw(1, "没有此银行卡")
	}
	return datas[0]
}

// 根据用户ID获取所有的银行卡信息
func (this *CardAoModel) QueryCardByUserId(userid int) Cards {
	datas := this.CardDb.QueryCardByUserId(userid)

	return datas
}

// 添加银行卡
func (this *CardAoModel) Add(acard Card) {
	//this.IsAdmin()
	userid := this.UserAo.GetUserIdFromSession()
	if userid == 0 {
		Throw(1, "获取用户信息失败")
	}

	if acard.Name == "" {
		Throw(1, "请输入银行卡名称")
	}

	if acard.Bank == "" {
		Throw(1, "请输入银行名称")
	}

	if acard.Card == "" {
		Throw(1, "请输入银行卡账户")
	}

	if acard.Money < 0 {
		Throw(1, "初始余额不能小于0")
	}

	// 验证银卡号是否已经存在
	datas := this.CardDb.QueryCardByCardNum(acard.Card)
	if len(datas) > 0 {
		Throw(1, "已经存在相同卡号的银行卡")
	}

	acard.UserId = userid
	acard.CreateTime = time.Now()
	acard.ModifyTime = time.Now()

	this.CardDb.Add(acard)
}

// 修改银行卡信息
func (this *CardAoModel) Update(acard Card) {
	userid := this.UserAo.GetUserIdFromSession()

	if userid == 0 {
		Throw(1, "获取用户信息失败")
	}

	if acard.Name == "" {
		Throw(1, "请输入银行卡名称")
	}

	if acard.Bank == "" {
		Throw(1, "请输入银行名称")
	}

	if acard.Card == "" {
		Throw(1, "请输入银行卡账户")
	}

	if acard.Money < 0 {
		Throw(1, "初始余额不能小于0")
	}

	// 验证银卡号是否已经存在
	datas := this.CardDb.QueryCardByCardNumAndCardId(acard.Card, acard.CardId)
	if len(datas) > 0 {
		Throw(1, "已经存在相同卡号的银行卡")
	}

	acard.UserId = userid
	acard.ModifyTime = time.Now()

	this.CardDb.Update(acard)
}

// 删除银行卡信息
func (this *CardAoModel) Delete(cardid int) {
	this.UserAo.IsAdmin()
	this.CardDb.Delete(cardid)

	this.Queue.Publish(CardQueueEnum.DELEVENT, cardid)
}
