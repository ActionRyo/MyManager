package account

import (
	. "github.com/fishedee/language"
)

var AccountWeekType struct {
	EnumStruct
	Income         int `enum:"1,收入"`
	Pay            int `enum:"2,支出"`
	TransferIncome int `enum:"3,转账收入"`
	TransferPay    int `enum:"4,转账支出"`
	LendIncome     int `enum:"5,借还款收入"`
	LendPay        int `enum:"6,借还款支出"`
}

func init() {
	InitEnumStruct(&AccountWeekType)
}
