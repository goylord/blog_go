package mysql

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose"
)

var dbConfig = &gorose.Config{
	Driver:          "mysql",
	SetMaxOpenConns: 0,
	SetMaxIdleConns: 0,
	Prefix:          "",
	Dsn:             "root:rengenshenghsa@tcp(106.13.67.87:3306)/blog?charset=utf8",
}

// var db gorose.IOrm
var connection *gorose.Engin

func init() {
	var err error
	connection, err = gorose.Open(dbConfig)
	if err != nil {
		fmt.Printf("数据库连接失败----> %s", err.Error())
		os.Exit(1)
	}
	fmt.Printf("打开成功", connection)
}
func Conn() gorose.IOrm {
	return connection.NewOrm()
}
