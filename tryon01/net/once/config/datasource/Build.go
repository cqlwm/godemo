package datasource

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Instance *sql.DB

func Build() error {
	username := "root"
	password := "WJttYV8anCwbjc6l"
	url := "rm-2vca8eh1prx298nrxoo.mysql.cn-chengdu.rds.aliyuncs.com:3306"
	dbName := "look"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",username, password, url, dbName)
	var err error
	if Instance, err = sql.Open("mysql", dsn); err != nil {
		return err
	}
	return Instance.Ping()
}
