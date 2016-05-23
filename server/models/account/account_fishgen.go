package account

import (
	. "github.com/fishedee/language"
	"mymanager/models/condition"
	"time"
)

func (this *AccountAoModel) QueryAccountByCondition_WithError(data Account, lim condition.Limit) (_fishgen1 Accounts, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.QueryAccountByCondition(data, lim)
	return
}

func (this *AccountAoModel) QueryAccountByAccountId_WithError(aid int) (_fishgen1 []Account, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.QueryAccountByAccountId(aid)
	return
}

func (this *AccountAoModel) Add_WithError(acc Account) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Add(acc)
	return
}

func (this *AccountAoModel) Update_WithError(aAccount Account) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Update(aAccount)
	return
}

func (this *AccountAoModel) Delete_WithError(accountId int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Delete(accountId)
	return
}

func (this *AccountAoModel) DeleteAccByCardID_WithError(cardid int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.DeleteAccByCardID(cardid)
	return
}

func (this *AccountAoModel) DeleteAccByCategoryId_WithError(categoryid int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.DeleteAccByCategoryId(categoryid)
	return
}

func (this *AccountAoModel) GetWeekTypeStatistic_WithError(userId int) (_fishgen1 []WeekType, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetWeekTypeStatistic(userId)
	return
}

func (this *AccountAoModel) GetWeekDetailTypeStatistic_WithError(year int, week int, wt int) (_fishgen1 []WeekTypeDetail, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetWeekDetailTypeStatistic(year, week, wt)
	return
}

func (this *AccountAoModel) GetWeekCardStatistic_WithError() (_fishgen1 []WeekType, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetWeekCardStatistic()
	return
}

func (this *AccountAoModel) GetWeekDetailCardStatistic_WithError(year int, week int, cardid int) (_fishgen1 []WeekTypeDetail, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetWeekDetailCardStatistic(year, week, cardid)
	return
}

func (this *AccountAoModel) WhenCardDel_WithError(cardId int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.WhenCardDel(cardId)
	return
}

func (this *AccountAoModel) WhenCategoryDel_WithError(categoryId int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.WhenCategoryDel(categoryId)
	return
}

func (this *AccountDbModel) QueryAccountByCondition_WithError(data Account, lim condition.Limit) (_fishgen1 Accounts, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.QueryAccountByCondition(data, lim)
	return
}

func (this *AccountDbModel) QueryAccountByAccountId_WithError(accountid int) (_fishgen1 []Account, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.QueryAccountByAccountId(accountid)
	return
}

func (this *AccountDbModel) QueryAccountByCardId_WithError(cardid int) (_fishgen1 []Account, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.QueryAccountByCardId(cardid)
	return
}

func (this *AccountDbModel) QueryAccountByCategoryId_WithError(categoryid int) (_fishgen1 []Account, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.QueryAccountByCategoryId(categoryid)
	return
}

func (this *AccountDbModel) Add_WithError(data Account) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Add(data)
	return
}

func (this *AccountDbModel) Update_WithError(data Account) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Update(data)
	return
}

func (this *AccountDbModel) Delete_WithError(accountid int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Delete(accountid)
	return
}

func (this *AccountDbModel) DeleteAccByCardID_WithError(userid int, card int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.DeleteAccByCardID(userid, card)
	return
}

func (this *AccountDbModel) DeleteAccByCategoryId_WithError(userid int, categoryid int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.DeleteAccByCategoryId(userid, categoryid)
	return
}

func (this *AccountDbModel) QueryAccountByUserId_WithError(userid int) (_fishgen1 []Account, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.QueryAccountByUserId(userid)
	return
}

func (this *AccountDbModel) QueryAccountWeekByUserId_WithError(userid int) (_fishgen1 []AccountWithWeek, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.QueryAccountWeekByUserId(userid)
	return
}

func (this *AccountDbModel) QueryWeekTypeStartDate_WithError(userid int, t int, year int, week int) (_fishgen1 []WeekCreateTime, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.QueryWeekTypeStartDate(userid, t, year, week)
	return
}

func (this *AccountDbModel) QueryWeekTypeDetail_WithError(userId int, wtype int, startDate time.Time, endDate time.Time) (_fishgen1 []Account, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.QueryWeekTypeDetail(userId, wtype, startDate, endDate)
	return
}

func (this *AccountDbModel) QueryWeekCardStartDate_WithError(userid int, year int, week int, cardid int) (_fishgen1 []WeekCreateTime, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.QueryWeekCardStartDate(userid, year, week, cardid)
	return
}

func (this *AccountDbModel) QueryWeekCardDetail_WithError(userId int, cardid int, startDate time.Time, endDate time.Time) (_fishgen1 []Account, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.QueryWeekCardDetail(userId, cardid, startDate, endDate)
	return
}
