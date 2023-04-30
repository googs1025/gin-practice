package UserModel

import (
	dbinit "gin-practice/pkg/db"
	_ "gin-practice/pkg/validators"
	"log"
)

func init() {
	initTable()
}

type UserModel struct {
	UserID   int    `gorm:"column:user_id;type:int;primaryKey;autoIncrement" json:"user_id" form:"id" `
	UserName string `gorm:"column:user_name;unique;type:varchar(50)" json:"user_name" form:"name" ` // 把valid校验的参数全都放入validator中。
	//UserName string `gorm:"column:user_name;unique;type:varchar(50)" json:"user_name" form:"name" binding:"required,min=4,UserName"`	// 都写在这里比较杂
	UserPwd     string `gorm:"column:user_pwd;type:varchar(50)" json:"user_pwd" `
	UserAddtime string `gorm:"column:user_addtime;type:datetime" json:"addTime" `
}

// initTable 自动创建表结构
func initTable() {
	err := dbinit.DB.AutoMigrate(&UserModel{})
	if err != nil {
		log.Fatal(err)
	}
}

func (u *UserModel) TableName() string {
	return "users"
}

func NewUserModel(attr ...UserModelsAttrFunc) *UserModel {
	u := &UserModel{}
	UserModelsAttrFuncList(attr).Apply(u)

	return u
}

//// NewUserModel 只有定义空对象。
//func NewUserModelWithNothing() *UserModel {
//	return &UserModel{}
//}
//
//// NewUserModelWithUserID 可以用这种方式，new对象，但是不太方便，可以换一种方式实现，里用传入可变参数的方式。
//func NewUserModelWithUserID(id int) *UserModel {
//	return &UserModel{
//		UserID: id,
//	}
//}

// 修改方法。
// Mutate new构建外，需要修改值的处理方法。
func (u *UserModel) Mutate(attr ...UserModelsAttrFunc) *UserModel {
	UserModelsAttrFuncList(attr).Apply(u)
	return u
}
