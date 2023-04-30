package main

import (
	"gin-practice/pkg/common"
	"gin-practice/pkg/handlers"
	"gin-practice/pkg/models/UserModel"
	"gin-practice/pkg/result"
	"gin-practice/pkg/test"
	_ "gin-practice/pkg/validators"
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

	// 练习的路由用。
	example := make(map[string]string)
	example["username"] = "jiangjiang"
	topic := &Topic{
		TopicID:    111,
		TopicTitle: "testingtesting",
	}

	r.GET("/try", func(c *gin.Context) {
		c.Writer.Write([]byte("hello\n")) // 直接写数据
		// 可以把对象直接转为json。
		c.JSON(200, example)
		c.JSON(200, topic)

	})

	// 执行
	defer func() {
		// 初始化日记
		common.InitLogger()
		_ = r.Run(":8080")
	}()

}

type Topic struct {
	TopicID    int
	TopicTitle string
}
