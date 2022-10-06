package main


import (
	"gin-practice/src/common"
	"gin-practice/src/handlers"
	"gin-practice/src/models/UserModel"
	"gin-practice/src/result"
	"gin-practice/src/test"
	_ "gin-practice/src/validators"
	"github.com/gin-gonic/gin"
)

// TODO: 需要把server代码和路由代码拆开，可以搞一个router.go文件

func main() {
	r := gin.New()
	r.Use(common.ErrorHandler())


	// model初始化的方式推荐
	//user := UserModel.NewUserModel(UserModel.SetUserModelWithUserID(11), UserModel.SetUserModelWithUserName("jiang"))

	// 路由一
	r.GET("/", func(c *gin.Context) {
		//user := UserModel.NewUserModel()
		user := UserModel.NewUserModel(UserModel.SetUserModelWithUserID(11), UserModel.SetUserModelWithUserName("jiang"))

		// 可以自定义配置model参数。
		c.JSON(200, user)
	})

	// 路由二
	r.GET("/user", func(c *gin.Context) {
		user := UserModel.NewUserModel()
		user.Mutate(UserModel.SetUserModelWithUserID(111), UserModel.SetUserModelWithUserName("jiangjiang"))
		c.JSON(200, user)
	})
	
	r.POST("/user11", func(c *gin.Context) {
		user := UserModel.NewUserModel()

		result.Result(c.ShouldBind(user)).UnWrap()

		info := result.Result(test.GetInfo(user.UserID)).UnWrap()
		//c.JSON(200, user)
		c.JSON(200, info)
		// 统一错误处理前的方式
		//if err := c.ShouldBind(user); err != nil {
		//	panic(err)
		//	//c.JSON(400, gin.H{"message":err.Error()})
		//} else {
		//	c.JSON(200, user)
		//}
	})

	// 路由：获取user列表
	r.GET("/users", handlers.UserList)

	r.GET("/users/:id", handlers.UserDetail)

	r.POST("/user", handlers.UserSave)


	// 执行
	defer func() {
		_ = r.Run(":8080")
	}()



}
