package mysql

import(
	"github.com/gohouse/gorose"
	_"github.com/gohouse/gorose/driver/mysql"
	"fmt"
	"os"

)
var dbConfig = &gorose.DbConfigSingle{
	Driver: "mysql",
	EnableQueryLog: true,
	SetMaxOpenConns: 0,
	SetMaxIdleConns: 0,
	Prefix: "",
	Dsn:"root:rengenshenghsa@tcp(localhost:3306)/blog?charset=utf8",
}
var db *gorose.Connection
func init() {
	connection, err := gorose.Open(dbConfig)
	if err != nil {
		fmt.Printf("数据库连接失败----> %s", err.Error())
		os.Exit(1)
	}
	db = connection
}
func Conn() *gorose.Connection {
	return db
}