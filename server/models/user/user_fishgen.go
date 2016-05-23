package user

import (
	. "github.com/fishedee/language"
	"mymanager/models/condition"
)

func (this *UserAoModel) Login_WithError(name string, password string) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Login(name, password)
	return
}

func (this *UserAoModel) Logout_WithError() (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Logout()
	return
}

func (this *UserAoModel) CheckMustLogin_WithError() (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.CheckMustLogin()
	return
}

func (this *UserAoModel) IsAdmin_WithError() (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.IsAdmin()
	return
}

func (this *UserAoModel) QueryUserByCondition_WithError(data User, lim condition.Limit) (_fishgen1 Users, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.QueryUserByCondition(data, lim)
	return
}

func (this *UserAoModel) QueryUserByUserId_WithError(userid int) (_fishgen1 []User, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.QueryUserByUserId(userid)
	return
}

func (this *UserAoModel) Add_WithError(name string, password string, t int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Add(name, password, t)
	return
}

func (this *UserAoModel) UpdatePassword_WithError(data User) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.UpdatePassword(data)
	return
}

func (this *UserAoModel) UpdateRole_WithError(data User) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.UpdateRole(data)
	return
}

func (this *UserAoModel) UpdateSelfPwd_WithError(oldPwd string, newPwd string) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.UpdateSelfPwd(oldPwd, newPwd)
	return
}

func (this *UserAoModel) Delete_WithError(userid int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Delete(userid)
	return
}

func (this *UserAoModel) SafePwd_WithError(password string) (_fishgen1 string, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.SafePwd(password)
	return
}

func (this *UserAoModel) GetUserIdFromSession_WithError() (_fishgen1 int, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.GetUserIdFromSession()
	return
}

func (this *UserDbModel) QueryUserByNameAndPassword_WithError(name string, pwd string) (_fishgen1 []User, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.QueryUserByNameAndPassword(name, pwd)
	return
}

func (this *UserDbModel) QueryUserByUserId_WithError(userid int) (_fishgen1 []User, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.QueryUserByUserId(userid)
	return
}

func (this *UserDbModel) QueryUserByName_WithError(name string) (_fishgen1 []User, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.QueryUserByName(name)
	return
}

func (this *UserDbModel) QueryUserByCondition_WithError(data User, lim condition.Limit) (_fishgen1 Users, _fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	_fishgen1 = this.QueryUserByCondition(data, lim)
	return
}

func (this *UserDbModel) Add_WithError(data User) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Add(data)
	return
}

func (this *UserDbModel) Update_WithError(data User) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Update(data)
	return
}

func (this *UserDbModel) Delete_WithError(useid int) (_fishgenErr Exception) {
	defer Catch(func(exception Exception) {
		_fishgenErr = exception
	})
	this.Delete(useid)
	return
}
