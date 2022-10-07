package Getter

import (
	"fmt"
	"gin-practice/src/data/mappers"
	dbinit "gin-practice/src/db"
	"gin-practice/src/models/UserModel"
	"gin-practice/src/result"
)

func init() {
	// 具体都是操作接口，不是操作具体实现
	UserGetter = NewUserGetterImpl()
}

var UserGetter IUserGetter

type IUserGetter interface {
	GetUserList() []*UserModel.UserModel
	//GetUserByID(id int) *UserModel.UserModel 不要只用这个，可以用error.Result
	GetUserByID(id int) *result.ErrorResult	// 只要涉及到error的，就用*result.ErrorResult返回
}

type UserGetterImpl struct {
	userMapper *mappers.UserMapper
}

func NewUserGetterImpl() *UserGetterImpl {
	return &UserGetterImpl{userMapper: &mappers.UserMapper{} }
}

func (u *UserGetterImpl) GetUserList () (users []*UserModel.UserModel) {
	dbinit.DB.Find(&users)

	//sqlMapper := u.userMapper.GetUserList()
	//dbinit.DB.Raw(sqlMapper.Sql, sqlMapper.Args).Find(&users)
	return
}

func (u *UserGetterImpl) GetUserByID (id int) *result.ErrorResult {
	user := &UserModel.UserModel{}
	res := dbinit.DB.Where("user_id = ?", id).Find(user)
	if res.Error != nil || res.RowsAffected == 0 {
		return result.Result(nil, fmt.Errorf("not found element, user id=%d", id))
	}

	return result.Result(user, nil)


}