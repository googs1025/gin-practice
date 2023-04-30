package dbinit

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
	"strings"
)

var (
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

func init() {
	workdir, err := os.Getwd()
	wrpath := "/pkg/db/"
	//wrpath := "/"	  // 执行测试文件要用这个路径，不然会报错
	var str = []string{workdir, wrpath, "config.ini"}
	path := strings.Join(str, "")

	config, err := ini.Load(path)
	checkErr(err)
	LoadMysqlData(config)
	dbConfig := []string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":" + DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}
	pathDB := strings.Join(dbConfig, "")
	initDB(pathDB)

	LoadMysqlDataSlave(config)
	dbConfig = []string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":" + DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}
	pathDB = strings.Join(dbConfig, "")
	initDBSlave(pathDB)

}

func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassword = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()

}

func LoadMysqlDataSlave(file *ini.File) {
	Db = file.Section("mysql_slave").Key("Db").String()
	DbHost = file.Section("mysql_slave").Key("DbHost").String()
	DbPort = file.Section("mysql_slave").Key("DbPort").String()
	DbUser = file.Section("mysql_slave").Key("DbUser").String()
	DbPassword = file.Section("mysql_slave").Key("DbPassWord").String()
	DbName = file.Section("mysql_slave").Key("DbName").String()

}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
