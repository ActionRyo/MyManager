package category

import (
	. "github.com/fishedee/language"
	"mymanager/models/condition"
)

func (this *CategoryAoModel) QueryCategoryByCondition_WithError(data Category, lim condition.Limit) (_fishgen1 Categorys, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.QueryCategoryByCondition(data, lim)
	return
}

func (this *CategoryAoModel) QueryCategoryByCategoryId_WithError(cid int) (_fishgen1 []Category, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.QueryCategoryByCategoryId(cid)
	return
}

func (this *CategoryAoModel) QueryCategoryByUserid_WithError(userid int) (_fishgen1 []Category, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.QueryCategoryByUserid(userid)
	return
}

func (this *CategoryAoModel) Add_WithError(name string, rem string) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Add(name, rem)
	return
}

func (this *CategoryAoModel) Update_WithError(aCategory Category) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Update(aCategory)
	return
}

func (this *CategoryAoModel) Delete_WithError(Categoryid int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Delete(Categoryid)
	return
}

func (this *CategoryDbModel) QueryCategoryByCondition_WithError(data Category, lim condition.Limit) (_fishgen1 Categorys, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.QueryCategoryByCondition(data, lim)
	return
}

func (this *CategoryDbModel) QueryCategoryByCategoryId_WithError(categoryid int) (_fishgen1 []Category, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.QueryCategoryByCategoryId(categoryid)
	return
}

func (this *CategoryDbModel) QueryCategoryByUserid_WithError(userid int) (_fishgen1 []Category, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.QueryCategoryByUserid(userid)
	return
}

func (this *CategoryDbModel) QueryCategoryByName_WithError(name string, userid int) (_fishgen1 []Category, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.QueryCategoryByName(name, userid)
	return
}

func (this *CategoryDbModel) QueryCategoryByNameForUpd_WithError(name string, uid int, cid int) (_fishgen1 []Category, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.QueryCategoryByNameForUpd(name, uid, cid)
	return
}

func (this *CategoryDbModel) Add_WithError(data Category) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Add(data)
	return
}

func (this *CategoryDbModel) Update_WithError(data Category) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Update(data)
	return
}

func (this *CategoryDbModel) Delete_WithError(categoryid int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Delete(categoryid)
	return
}
