package account

import (
	"fmt"
	. "github.com/fishedee/language"
	"github.com/fishedee/web"
	. "mymanager/models/card"
	. "mymanager/models/category"
	"mymanager/models/condition"
	. "mymanager/models/user"
	"time"
)

type AccountAoModel struct {
	web.Model
	AccountDb  AccountDbModel
	UserAo     UserAoModel
	CategoryAo CategoryAoModel
	CardAo     CardAoModel
}

// 根据条件获取分类信息
func (this *AccountAoModel) QueryAccountByCondition(data Account, lim condition.Limit) Accounts {

	userid := this.UserAo.GetUserIdFromSession()
	data.UserId = userid
	datas := this.AccountDb.QueryAccountByCondition(data, lim)

	return datas
}

// 根据分类ID查询分类信息
func (this *AccountAoModel) QueryAccountByAccountId(aid int) []Account {
	//userid := this.UserAo.GetUserIdFromSession()
	datas := this.AccountDb.QueryAccountByAccountId(aid)

	return datas
}

// 添加分类
func (this *AccountAoModel) Add(acc Account) {
	//this.IsAdmin()
	userid := this.UserAo.GetUserIdFromSession()
	if userid == 0 {
		Throw(1, "获取用户信息失败")
	}

	if acc.Name == "" {
		Throw(1, "请输入分类名称")
	}

	if acc.Money <= 0 {
		Throw(1, "金额必须大于0")
	}

	if acc.CategoryId == 0 {
		Throw(1, "请选择分类")
	}

	if acc.CardId == 0 {
		Throw(1, "请选择银行卡")
	}

	if acc.Remark == "" {
		Throw(1, "请输入备注")
	}

	acc.UserId = userid
	acc.CreateTime = time.Now()
	acc.ModifyTime = time.Now()

	this.AccountDb.Add(acc)
}

// 修改分类信息
func (this *AccountAoModel) Update(aAccount Account) {
	userid := this.UserAo.GetUserIdFromSession()
	if userid == 0 {
		Throw(1, "获取用户信息失败")
	}

	// 判断数据是否存在DB
	datas := this.QueryAccountByAccountId(aAccount.AccountId)
	if len(datas) == 0 {
		Throw(1, "不存在相关数据")
	}

	if aAccount.Name == "" {
		Throw(1, "请输入分类名称")
	}

	if aAccount.Money <= 0 {
		Throw(1, "金额必须大于0")
	}

	if aAccount.CategoryId == 0 {
		Throw(1, "请选择分类")
	}

	if aAccount.CardId == 0 {
		Throw(1, "请选择银行卡")
	}

	if aAccount.Remark == "" {
		Throw(1, "请输入备注")
	}

	aAccount.UserId = userid
	aAccount.ModifyTime = time.Now()

	this.AccountDb.Update(aAccount)
}

// 删除分类信息
func (this *AccountAoModel) Delete(accountId int) {
	this.UserAo.IsAdmin()
	datas := this.QueryAccountByAccountId(accountId)
	if len(datas) == 0 {
		Throw(1, "不存在相关数据")
	}
	this.AccountDb.Delete(accountId)
}

func (this *AccountAoModel) DeleteAccByCardID(cardid int) {
	this.UserAo.IsAdmin()
	userid := this.UserAo.GetUserIdFromSession()
	if userid == 0 {
		Throw(1, "获取用户信息失败")
	}

	datas := this.AccountDb.QueryAccountByCardId(cardid)
	if len(datas) == 0 {
		Throw(1, "不存在相关数据")
	}
	this.AccountDb.DeleteAccByCardID(userid, cardid)
}

func (this *AccountAoModel) DeleteAccByCategoryId(categoryid int) {
	this.UserAo.IsAdmin()
	userid := this.UserAo.GetUserIdFromSession()
	if userid == 0 {
		Throw(1, "获取用户信息失败")
	}

	datas := this.AccountDb.QueryAccountByCategoryId(categoryid)
	if len(datas) == 0 {
		Throw(1, "不存在相关数据")
	}
	this.AccountDb.DeleteAccByCategoryId(userid, categoryid)
}

// 统计每周的数据信息
func (this *AccountAoModel) GetWeekTypeStatistic(userId int) []WeekType {
	userid := this.UserAo.GetUserIdFromSession()
	if userid == 0 {
		Throw(1, "获取用户信息失败")
	}

	enum := AccountWeekType.Datas()

	//accData := this.AccountDb.QueryAccountByUserId(userid)
	accData := this.AccountDb.QueryAccountWeekByUserId(userid)
	lstWeekType := QuerySelect(accData, func(one AccountWithWeek) WeekType {
		//year, week := one.CreateTime.ISOWeek()
		result := WeekType{
			Money: one.Money,
			Type:  one.Type,
			Week:  one.Week,
			Year:  one.Year,
		}

		return result
	}).([]WeekType)

	yearWeekType := QueryGroup(lstWeekType, "Year,Week", func(list []WeekType) []WeekType {
		rdate := QueryGroup(list, "Type", func(list []WeekType) []WeekType {
			sum := QuerySum(QueryColumn(list, "Money"))
			list[0].Money = sum.(int)
			return []WeekType{list[0]}
		}).([]WeekType)

		one := list[0]
		rdate = QueryLeftJoin(enum, rdate, "Id = Type", func(all EnumData, target WeekType) WeekType {
			return WeekType{
				Year: one.Year,
				Week: one.Week,
				Name: fmt.Sprintf(
					"%4d年%02d周",
					one.Year,
					one.Week,
				),
				Type:     all.Id,
				TypeName: all.Name,
				Money:    target.Money,
			}
		}).([]WeekType)

		return rdate
	}).([]WeekType)

	//排序
	rdate := QuerySort(yearWeekType, "Year,Week,Type").([]WeekType)

	return rdate
}

// 详细周数据类型统计
func (this *AccountAoModel) GetWeekDetailTypeStatistic(year int, week int, wt int) []WeekTypeDetail {
	userid := this.UserAo.GetUserIdFromSession()
	if userid == 0 {
		Throw(1, "获取用户信息失败")
	}

	getDate := this.AccountDb.QueryWeekTypeStartDate(userid, wt, year, week)
	if len(getDate) == 0 {
		Throw(1, "没找到相关数据")
	}

	endDate := getDate[0].CreateTime.AddDate(0, 0, 8)

	// 获取该段时间范围内的统计数据
	data := this.AccountDb.QueryWeekTypeDetail(userid, wt, getDate[0].CreateTime, endDate)

	//this.Log.Debug("%v", data)

	allCateData := this.CategoryAo.QueryCategoryByUserid(userid)

	//this.Log.Debug("%v", allCateData)

	result := QueryGroup(data, "CategoryId", func(one []Account) []Account {
		sum := QuerySum(QueryColumn(one, "Money"))
		one[0].Money = sum.(int)
		return []Account{one[0]}
	}).([]Account)

	totalPrice := float64((QuerySum(QueryColumn(result, "Money")).(int)))
	resultData := QueryLeftJoin(result, allCateData, "CategoryId=CategoryId", func(one Account, all Category) WeekTypeDetail {
		categoryMoney := float64(one.Money)
		categoryMoneyPrec := categoryMoney / totalPrice * 100
		return WeekTypeDetail{
			CategoryId:   one.CategoryId,
			CategoryName: all.Name,
			Precent:      categoryMoneyPrec,
			Money:        one.Money,
		}
	}).([]WeekTypeDetail)
	//this.Log.Debug("%v", resultData)
	return resultData
}

// 统计银行卡数据信息
func (this *AccountAoModel) GetWeekCardStatistic() []WeekType {
	userid := this.UserAo.GetUserIdFromSession()
	if userid == 0 {
		Throw(1, "获取用户信息失败")
	}

	// 获取统计的相关数据
	accData := this.AccountDb.QueryAccountWeekByUserId(userid)
	lstWeekType := QuerySelect(accData, func(one AccountWithWeek) WeekType {
		result := WeekType{
			Money:  one.Money,
			Type:   one.Type,
			Week:   one.Week,
			Year:   one.Year,
			CardId: one.CardId,
		}

		return result
	}).([]WeekType)

	// 获取用户旗下的所有银行卡信息
	cardData := this.CardAo.QueryCardByUserId(userid)

	this.Log.Debug("%#v", cardData)

	// 分类计算金额
	startMoney := map[int]int{}
	for _, aCard := range cardData.Data {
		startMoney[aCard.CardId] = aCard.Money
	}

	lstWeekType = QueryGroup(lstWeekType, "Year,Week", func(lstWt []WeekType) []WeekType {
		result := QueryGroup(lstWt, "CardId", func(lstWt []WeekType) []WeekType {
			cardTotal := QuerySum(QuerySelect(lstWt, func(onewt WeekType) int {
				//if onewt.Type%2 == 0 {
				//	return -onewt.Money
				//} else {
				//	return onewt.Money
				//}

				//var targetMoney int
				if onewt.Type == 1 {
					return onewt.Money
				} else if onewt.Type == 2 {
					return -onewt.Money
				} else if onewt.Type == 3 {
					return onewt.Money
				} else if onewt.Type == 4 {
					return -onewt.Money
				} else if onewt.Type == 5 {
					return onewt.Money
				} else if onewt.Type == 6 {
					return -onewt.Money
				} else {
					Throw(1, "类别错误")
					return 0
				}

				//return targetMoney

			})).(int)
			lstWt[0].Money = cardTotal
			this.Log.Debug("Year:%v Week:%v Card:%v Money:%v", lstWt[0].Year, lstWt[0].Week, lstWt[0].CardId, lstWt[0].Money)
			return []WeekType{lstWt[0]}
		}).([]WeekType)

		onels := lstWt[0]
		result = QueryLeftJoin(cardData.Data, result, "CardId = CardId", func(c Card, wt WeekType) WeekType {
			sumMony := startMoney[c.CardId] + wt.Money
			return WeekType{
				Year: onels.Year,
				Week: onels.Week,
				Name: fmt.Sprintf(
					"%4d年%02d周",
					onels.Year,
					onels.Week,
				),
				Money:    sumMony,
				CardId:   c.CardId,
				CardName: c.Name,
			}
		}).([]WeekType)

		return result
	}).([]WeekType)

	//排序
	rdata := QuerySort(lstWeekType, "Year,Week,CardId").([]WeekType)

	return rdata
}

// 详细银行卡的收支信息
func (this *AccountAoModel) GetWeekDetailCardStatistic(year int, week int, cardid int) []WeekTypeDetail {
	userid := this.UserAo.GetUserIdFromSession()
	if userid == 0 {
		Throw(1, "获取用户信息失败")
	}

	startDate := this.AccountDb.QueryWeekCardStartDate(userid, year, week, cardid)
	endDate := startDate[0].CreateTime.AddDate(0, 0, 8)

	accData := this.AccountDb.QueryWeekCardDetail(userid, cardid, startDate[0].CreateTime, endDate)
	enums := AccountWeekType.Datas()

	mergerData := QueryGroup(accData, "Type", func(lstAcc []Account) []Account {
		sum := QuerySum(QueryColumn(lstAcc, "Money"))
		lstAcc[0].Money = sum.(int)
		return []Account{lstAcc[0]}
	}).([]Account)

	totalmoney := (float64)(QuerySum(QueryColumn(mergerData, "Money")).(int))
	rData := QueryLeftJoin(mergerData, enums, "Type = Id", func(one Account, enum EnumData) WeekTypeDetail {
		typeMoney := float64(one.Money) / totalmoney * 100
		return WeekTypeDetail{
			Type:     one.Type,
			TypeName: enum.Name,
			Precent:  typeMoney,
			Money:    one.Money,
		}
	}).([]WeekTypeDetail)

	return rData
}

func (this *AccountAoModel) WhenCardDel(cardId int) {
	this.Log.Debug("%v card is del", cardId)
	this.DeleteAccByCardID(cardId)
}

func (this *AccountAoModel) WhenCategoryDel(categoryId int) {
	this.Log.Debug("%v card is del", categoryId)
	this.DeleteAccByCategoryId(categoryId)
}

func init() {
	web.InitDaemon(func(this *AccountAoModel) {
		this.Queue.Subscribe(CardQueueEnum.DELEVENT, this.WhenCardDel)
	})

	web.InitDaemon(func(this *AccountAoModel) {
		this.Queue.Subscribe(CategoryQueueEnum.DELEVENT, this.WhenCategoryDel)
	})
}
