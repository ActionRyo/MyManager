package account

import "time"

type Account struct {
	AccountId  int
	UserId     int
	CategoryId int
	CardId     int
	Type       int
	Name       string
	Money      int
	Remark     string
	CreateTime time.Time
	ModifyTime time.Time
}

type AccountWithWeek struct {
	AccountId  int
	UserId     int
	CategoryId int
	CardId     int
	Year       int
	Week       int
	Type       int
	Name       string
	Money      int
	Remark     string
	CreateTime time.Time
	ModifyTime time.Time
}

// 用于统计
type Accounts struct {
	Count int
	Data  []Account
}

type WeekType struct {
	CardId     int
	CardName   string
	Money      int
	Name       string
	Type       int
	TypeName   string
	Week       int
	Year       int
	CreateTime string
	ModifyTime string
}

type WeekTypeDetail struct {
	CategoryId   int
	CategoryName string
	Type         int
	TypeName     string
	Money        int
	Precent      float64
}

type WeekCreateTime struct {
	CreateTime time.Time
}
