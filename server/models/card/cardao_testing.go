package card

import (
	//. "github.com/fishedee/language"
	. "github.com/fishedee/web"
	. "mymanager/models/condition"
	"time"
)

type CardAoTest struct {
	Test
	CardAo CardAoModel
}

func (this *CardAoTest) InitializationData() (Card, Card) {

	var fstCard Card

	fstCard.CardId = 10021
	fstCard.UserId = 10004
	fstCard.Name = "工资卡"
	fstCard.Bank = "顺德农商银行"
	fstCard.Card = "62278912313"
	fstCard.Money = 100000
	fstCard.Remark = "今晚来我家睡好吗"
	fstCard.CreateTime = time.Now()
	fstCard.ModifyTime = time.Now()

	var SndCard Card

	SndCard.CardId = 10022
	SndCard.UserId = 10004
	SndCard.Name = "信用卡"
	SndCard.Bank = "中国招商银行"
	SndCard.Card = "62222"
	SndCard.Money = 5000
	SndCard.Remark = "信用卡"
	SndCard.CreateTime = time.Now()
	SndCard.ModifyTime = time.Now()

	return fstCard, SndCard
}

func (this *CardAoTest) InitializationSession(userid int) {
	sess, err := this.Session.SessionStart()
	if err != nil {
		panic(err)
	}

	sess.Set("UserId", userid)
	defer sess.SessionRelease()
}

// 测试添加方法
func (this *CardAoTest) Func_Add(fstData Card, scdData Card) {

	this.CardAo.Add_WithError(fstData)
	this.CardAo.Add_WithError(scdData)

}

func (this *CardAoTest) Func_QueryCardByCondition(fstData Card, scdData Card) {
	var conditionData Card
	lim := Limit{
		PageIndex: 0,
		PageSize:  1000,
	}

	initDatas := []Card{scdData, fstData}

	getFstCard, _ := this.CardAo.QueryCardByCondition_WithError(conditionData, lim)

	this.AssertEqual(getFstCard.Count, len(initDatas))

	this.AssertEqual(initDatas, getFstCard.Data)
}

func (this *CardAoTest) Func_QueryCardByCardId(fstData Card, scdData Card) {
	getFst, _ := this.CardAo.QueryCardByCardId_WithError(fstData.CardId)

	getScd, _ := this.CardAo.QueryCardByCardId_WithError(scdData.CardId)

	this.AssertEqual(getFst[0], fstData)
	this.AssertEqual(getScd[0], scdData)
}

func (this *CardAoTest) Func_Update(fstData Card, scdData Card) {
	fstData.Bank = "中国建设银行"
	fstData.Card = "62278912313"
	fstData.Money = 1000
	fstData.Remark = "今晚来我家睡好吗？"

	this.CardAo.Update_WithError(fstData)
	getFst, _ := this.CardAo.QueryCardByCardId_WithError(fstData.CardId)

	this.AssertEqual(getFst[0], fstData)

	scdData.Card = "62278912313"
	err := this.CardAo.Update_WithError(scdData)

	this.AssertError(err, 1, "已经存在相同卡号的银行卡")
}

func (this *CardAoTest) Func_Delete(fstData Card, scdData Card) {

	getDataFst, _ := this.CardAo.QueryCardByCardId_WithError(fstData.CardId)
	getDataScd, _ := this.CardAo.QueryCardByCardId_WithError(scdData.CardId)

	this.AssertEqual(1, len(getDataFst))
	this.AssertEqual(1, len(getDataScd))

	this.CardAo.Delete(fstData.CardId)
	this.CardAo.Delete(scdData.CardId)

	getDataDelFst, _ := this.CardAo.QueryCardByCardId_WithError(fstData.CardId)
	getDataDelScd, _ := this.CardAo.QueryCardByCardId_WithError(scdData.CardId)

	this.AssertEqual(0, len(getDataDelFst))
	this.AssertEqual(0, len(getDataDelScd))
}

func (this *CardAoTest) TestBasic() {
	this.InitializationSession(10004)
	fstData, scdData := this.InitializationData()

	this.Func_Add(fstData, scdData)

	this.Func_QueryCardByCondition(fstData, scdData)
	this.Func_QueryCardByCardId(fstData, scdData)

	this.Func_Update(fstData, scdData)

	this.Func_Delete(fstData, scdData)
}

//func (this *CardAoTest) TestBasic() {
//	_, dd := this.CardAo.QueryCardBySingleCardId_WithError(100023)
//	this.Log.Debug("%v", dd)
//}

func init() {
	InitTest(&CardAoTest{})
}
