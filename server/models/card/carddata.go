package card

import "time"

type Card struct {
	CardId     int
	UserId     int
	Bank       string
	Name       string
	Card       string
	Money      int
	Remark     string
	CreateTime time.Time
	ModifyTime time.Time
}

// 用于统计
type Cards struct {
	Count int
	Data  []Card
}
