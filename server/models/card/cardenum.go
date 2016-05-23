package card

import (
	. "github.com/fishedee/language"
)

var CardQueueEnum struct {
	EnumStructString
	DELEVENT string `enum:"/card/del,银行卡被删除了"`
}

func init() {
	InitEnumStructString(&CardQueueEnum)
}
