package category

import "time"

type Category struct {
	CategoryId int
	UserId     int
	Name       string
	Remark     string
	CreateTime time.Time
	ModifyTime time.Time
}

// 用于统计
type Categorys struct {
	Count int
	Data  []Category
}
