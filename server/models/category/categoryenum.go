package category

import . "github.com/fishedee/language"

var CategoryQueueEnum struct {
	EnumStructString
	DELEVENT string `enum:"/category/del,分类被删除了"`
}

func init() {
	InitEnumStructString(&CategoryQueueEnum)
}
