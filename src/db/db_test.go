package dbinit

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSquirrel(t *testing.T) {
	Convey("测试sql拼接！", t, func() {

		// SELECT user_id, user_name FROM users WHERE user_id=? ORDER BY user_id desc LIMIT 10 [12] <nil>
		// 返回 string, []interface{}, error: cd
		fmt.Println(squirrel.Select("user_id", "user_name").From("users").
			Where("user_id=?", 12).
			OrderBy("user_id desc").Limit(10).ToSql())

	})
}

type TestTable struct {
	Name string
	Age int

}

//func TestMySqlDB(t *testing.T) {
//	Convey("test db talbe", t, func() {
//		// 创建表结构
//		if !DB.Migrator().HasTable(&TestTable{}) {
//			_ = DB.AutoMigrate(&TestTable{})
//		}
//
//		user := &TestTable{Name: "Jinzhu111", Age: 18111}
//
//		// 插入数据。
//		result := DB.Create(user)
//		So(result.Error, ShouldBeNil)
//		fmt.Println(result.Error, result.RowsAffected)
//		// 删除数据。
//		result = DB.Where("name=?", "Jinzhu111").Delete(user)
//		fmt.Println(result.Error, result.RowsAffected)
//		So(result.Error, ShouldBeNil)
//		// 删除测试表。
//		_ = DB.Migrator().DropTable(user)
//	})
//
//
//
//}

func TestMySqlDB(t *testing.T) {
	Convey("test db talbe", t, func() {
		// 创建表结构
		if !DB.Migrator().HasTable(&TestTable{}) {
			_ = DB.AutoMigrate(&TestTable{})
		}

		if !DB_Slave.Migrator().HasTable(&TestTable{}) {
			_ = DB_Slave.AutoMigrate(&TestTable{})
		}


		user := &TestTable{Name: "Jinzhu111", Age: 18111}

		// 插入数据。
		result := DB.Create(user)
		So(result.Error, ShouldBeNil)
		fmt.Println(result.Error, result.RowsAffected)
		// 删除数据。
		result = DB_Slave.Where("name=?", "Jinzhu111").Delete(user)
		fmt.Println(result.Error, result.RowsAffected)
		So(result.Error, ShouldBeNil)
		// 删除测试表。
		//_ = DB.Migrator().DropTable(user)
	})



}
