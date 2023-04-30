package mappers

import (
	"gin-practice/pkg/models/UserModel"
	"github.com/Masterminds/squirrel"
	"time"
)

type UserMapper struct {
}

// 可以进一步修改
//func (*UserMapper) GetUserList()  {
//	squirrel.Select("user_id", "user_name").From("users").
//		Where("user_id=?", 12).
//		OrderBy("user_id desc").Limit(10).ToSql()
//}

//
func (*UserMapper) GetUserList() *SqlMapper {
	// Mapper(squirrel.Select("user_id", "user_name").From("users").Where("user_id=?", 12).OrderBy("user_id desc").Limit(10).ToSql())
	return Mapper(squirrel.Select("user_id", "user_name").From("users").OrderBy("user_id desc").Limit(10).ToSql())
}

func (*UserMapper) AddNewUser(user *UserModel.UserModel) *SqlMapper {
	return Mapper(squirrel.Insert(user.TableName()).Columns("user_name", "user_pwd", "user_addtime").Values(user.UserName, user.UserPwd, time.Now()).ToSql())
}
