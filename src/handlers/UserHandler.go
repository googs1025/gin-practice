package handlers

import (
	"gin-practice/src/data/Getter"
	"gin-practice/src/models/UserModel"
	"gin-practice/src/result"
	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {
	user := UserModel.NewUserModel()

	result.Result(c.ShouldBind(user)).UnWrap()

	// 比较不优雅的方法！
	//okFunc := OK(c)
	//errFunc := Error(c)
	// 已修改的方法
	//resFunc("UserList", "10001", result.Result(c.ShouldBind(user)).UnWrap())




	//if user.UserID > 100 {
	//	okFunc("UserList", "10001", result.Result(test.GetInfo(user.UserID)).UnWrap())
	//} else {
	//	errFunc("UserList", "10001", result.Result(test.GetInfo(user.UserID)).UnWrap())
	//}


	//resFunc := R(c)
	//if user.UserID > 10 {
	//	resFunc("UserList", "10001", result.Result(test.GetInfo(user.UserID)).UnWrap())(OK)
	//} else {
	//	resFunc("UserList", "10001", result.Result(test.GetInfo(user.UserID)).UnWrap())(Error)
	//}

	resFunc := R(c)
	resFunc("UserList", "10001", Getter.UserGetter.GetUserList())(OK)

	//resFunc := R(c)
	//if user.UserID > 10 {
	//	resFunc("UserList", "10001", Getter.UserGetter.GetUserList())(OK)
	//} else {
	//	resFunc("UserList", "10001", result.Result(test.GetInfo(user.UserID)).UnWrap())(Error)
	//}


}

func UserDetail(c *gin.Context) {
	// 用gin自带的来验证id
	id := &struct {
		Id int `uri:"id" binding:"required,gt=10"`
	}{}
	//id := c.Param("id")
	result.Result(c.ShouldBindUri(id)).UnWrap()

	resFunc := R(c)
	resFunc("UserDetail", "10002", Getter.UserGetter.GetUserByID(id.Id).UnWrap())(OK)


}

func UserSave(c *gin.Context) {
	u := UserModel.NewUserModel()
	result.Result(c.ShouldBindJSON(u)).UnWrap()

	resFunc := R(c)
	resFunc("UserSave", "10003", "true")(OK)
}
