package account

import (
	"github.com/fishedee/web"
	"mymanager/models/condition"
	"time"
)

type AccountDbModel struct {
	web.Model
}

// 获取财务信息
func (this *AccountDbModel) QueryAccountByCondition(data Account, lim condition.Limit) Accounts {

	datas := []Account{}

	db := this.DB.NewSession()
	db = db.And("userid=?", data.UserId)
	if data.Name != "" {
		db = db.And("name like ?", "%"+data.Name+"%")
	}

	if data.Remark != "" {
		db = db.And("remark like ?", "%"+data.Remark+"%")
	}

	if data.CategoryId != 0 {
		db = db.And("categoryId = ?", data.CategoryId)
	}

	if data.CardId != 0 {
		db = db.And("cardid = ?", data.CardId)
	}

	if data.Type != 0 {
		db = db.And("type = ?", data.Type)
	}

	err := db.Desc("accountid").Limit(lim.PageSize, lim.PageIndex).Find(&datas)
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

	if data.CategoryId != 0 {
		dbs = dbs.And("categoryId = ?", data.CategoryId)
	}

	if data.CardId != 0 {
		dbs = dbs.And("cardid = ?", data.CardId)
	}

	if data.Type != 0 {
		dbs = dbs.And("type = ?", data.Type)
	}

	total, errs := dbs.Count(&Account{})
	if errs != nil {
		panic(errs)
	}

	return Accounts{
		Count: int(total),
		Data:  datas,
	}
}

// 根据财务ID查询财务信息
func (this *AccountDbModel) QueryAccountByAccountId(accountid int) []Account {
	data := []Account{}
	err := this.DB.Where("accountid=?", accountid).Find(&data)
	if err != nil {
		panic(err)
	}

	return data
}

func (this *AccountDbModel) QueryAccountByCardId(cardid int) []Account {
	data := []Account{}
	err := this.DB.Where("cardid=?", cardid).Find(&data)
	if err != nil {
		panic(err)
	}

	return data
}

func (this *AccountDbModel) QueryAccountByCategoryId(categoryid int) []Account {
	data := []Account{}
	err := this.DB.Where("categoryid=?", categoryid).Find(&data)
	if err != nil {
		panic(err)
	}

	return data
}

// 添加财务信息
func (this *AccountDbModel) Add(data Account) {
	_, err := this.DB.Insert(data)
	if err != nil {
		panic(err)
	}
}

// 修改财务信息
func (this *AccountDbModel) Update(data Account) {
	_, err := this.DB.Where("accountid=?", data.AccountId).Update(&data)
	if err != nil {
		panic(err)
	}
}

// 删除财务信息
func (this *AccountDbModel) Delete(accountid int) {
	data := Account{}
	_, err := this.DB.Where("accountid=?", accountid).Delete(&data)
	if err != nil {
		panic(err)
	}
}

func (this *AccountDbModel) DeleteAccByCardID(userid int, card int) {
	data := Account{}
	db := this.DB
	_, err := db.Where("userid=?", userid).And("cardid=?", card).Delete(&data)
	if err != nil {
		panic(err)
	}
}

func (this *AccountDbModel) DeleteAccByCategoryId(userid int, categoryid int) {
	data := Account{}
	db := this.DB
	_, err := db.Where("userid=?", userid).And("categoryid=?", categoryid).Delete(&data)
	if err != nil {
		panic(err)
	}
}

// 根据用户ID获取财务
func (this *AccountDbModel) QueryAccountByUserId(userid int) []Account {
	data := []Account{}
	err := this.DB.Where("userid=?", userid).Find(&data)
	if err != nil {
		panic(err)
	}

	return data
}

// 根据用户ID获取财务
func (this *AccountDbModel) QueryAccountWeekByUserId(userid int) []AccountWithWeek {
	data := []AccountWithWeek{}
	//err := this.DB.Where("userid=?", userid).Find(&data)
	err := this.DB.Sql("select *,year(Createtime) as year,week(createtime) as week from t_account where userid=?", userid).Find(&data)
	if err != nil {
		panic(err)
	}

	return data
}

// 获取开始时间
func (this *AccountDbModel) QueryWeekTypeStartDate(userid int, t int, year int, week int) []WeekCreateTime {

	var data []WeekCreateTime

	err := this.DB.Sql("select DATE_FORMAT(DATE_ADD(createtime,INTERVAL-WEEKDAY(createtime) DAY),'%Y-%m-%d') as createtime from (select year(createtime) as year,week(createtime) as week,createtime from t_account where type =? and userid =?) t where t.year = ? and t.week =? order by createtime asc limit 1", t, userid, year, week).Find(&data)

	if err != nil {
		panic(err)
	}

	return data
}

// 根据类别获取详细的统计信息
func (this *AccountDbModel) QueryWeekTypeDetail(userId int, wtype int, startDate time.Time, endDate time.Time) []Account {
	var data []Account
	//err := this.DB.Sql("select * from t_account where userId=? AND type=? AND createTime between ? and ?", userId, wtype, startDate, endDate).Find(&data)
	err := this.DB.Sql("select * from t_account where userId=? and type=? and createTime>=? and createTime<?", userId, wtype, startDate, endDate).Find(&data)
	if err != nil {
		panic(err)
	}
	return data
}

// 获取银卡周的第一天的日期
func (this *AccountDbModel) QueryWeekCardStartDate(userid int, year int, week int, cardid int) []WeekCreateTime {
	var data []WeekCreateTime

	err := this.DB.Sql("select DATE_FORMAT(DATE_ADD(createtime,INTERVAL-WEEKDAY(createtime) DAY),'%Y-%m-%d') as createtime from (select year(createtime) as year,week(createtime) as week,createtime from t_account where cardid =? and userid =?) t where t.year = ? and t.week =? order by createtime asc limit 1", cardid, userid, year, week).Find(&data)

	if err != nil {
		panic(err)
	}

	return data
}

// 根据银行卡获取详细的统计信息
func (this *AccountDbModel) QueryWeekCardDetail(userId int, cardid int, startDate time.Time, endDate time.Time) []Account {
	var data []Account
	//err := this.DB.Sql("select * from t_account where userId=? AND type=? AND createTime between ? and ?", userId, wtype, startDate, endDate).Find(&data)
	err := this.DB.Sql("select * from t_account where userId=? and cardid=? and createTime>=? and createTime<?", userId, cardid, startDate, endDate).Find(&data)
	if err != nil {
		panic(err)
	}
	return data
}
