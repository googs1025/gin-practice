package UserModel


type UserModelsAttrFunc func(u *UserModel)

type UserModelsAttrFuncList []UserModelsAttrFunc

// Apply 方法
func (l UserModelsAttrFuncList) Apply(u *UserModel) {
	for _, f := range l {
		f(u)
	}
}


// SetUserModelWithUserID 可以用这种方式，set对象，但是不太方便，可以换一种方式实现，里用传入可变参数的方式。
func SetUserModelWithUserID(id int) UserModelsAttrFunc {
	return func(u *UserModel) {
		u.UserID = id
	}
}

// SetUserModelWithUserName 设置name
func SetUserModelWithUserName(name string) UserModelsAttrFunc {
	return func(u *UserModel) {
		u.UserName = name
	}
}


