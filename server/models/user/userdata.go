package user

import "time"

type User struct {
	UserId     int
	Name       string
	Password   string
	Type       int
	CreateTime time.Time
	ModifyTime time.Time
}

// 用于统计
type Users struct {
	Count int
	Data  []User
}
