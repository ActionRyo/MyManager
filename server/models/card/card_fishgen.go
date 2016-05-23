package card

import (
	. "github.com/fishedee/language"
	"mymanager/models/condition"
)

func (this *CardAoModel) QueryCardByCondition_WithError(data Card, lim condition.Limit) (_fishgen1 Cards, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.QueryCardByCondition(data, lim)
	return
}

func (this *CardAoModel) QueryCardByCardId_WithError(cardid int) (_fishgen1 []Card, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.QueryCardByCardId(cardid)
	return
}

func (this *CardAoModel) QueryCardBySingleCardId_WithError(cardid int) (_fishgen1 Card, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.QueryCardBySingleCardId(cardid)
	return
}

func (this *CardAoModel) QueryCardByUserId_WithError(userid int) (_fishgen1 Cards, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.QueryCardByUserId(userid)
	return
}

func (this *CardAoModel) Add_WithError(acard Card) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Add(acard)
	return
}

func (this *CardAoModel) Update_WithError(acard Card) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Update(acard)
	return
}

func (this *CardAoModel) Delete_WithError(cardid int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Delete(cardid)
	return
}

func (this *CardDbModel) QueryCardByCondition_WithError(data Card, lim condition.Limit) (_fishgen1 Cards, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.QueryCardByCondition(data, lim)
	return
}

func (this *CardDbModel) QueryCardByCardId_WithError(cardid int, userid int) (_fishgen1 []Card, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.QueryCardByCardId(cardid, userid)
	return
}

func (this *CardDbModel) QueryCardByCardNum_WithError(cardNum string) (_fishgen1 []Card, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.QueryCardByCardNum(cardNum)
	return
}

func (this *CardDbModel) QueryCardByCardNumAndCardId_WithError(cardNum string, cid int) (_fishgen1 []Card, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.QueryCardByCardNumAndCardId(cardNum, cid)
	return
}

func (this *CardDbModel) QueryCardByUserId_WithError(userid int) (_fishgen1 Cards, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.QueryCardByUserId(userid)
	return
}

func (this *CardDbModel) Add_WithError(data Card) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Add(data)
	return
}

func (this *CardDbModel) Update_WithError(data Card) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Update(data)
	return
}

func (this *CardDbModel) Delete_WithError(cardid int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Delete(cardid)
	return
}
