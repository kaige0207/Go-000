package mysqldb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kaige0207/Go-000/Week05/account/internal/pkg/configreader"
)

var (
	db  *sql.DB
	err error
)

func NewDB() (*sql.DB, error) {
	config := configreader.GetConfig()
	mysql := config.Mysql.Username + ":" + config.Mysql.Password + "@tcp(" + config.Mysql.Url + ")/" + config.Mysql.Database + "?charset=utf8"
	fmt.Println(mysql)
	if db == nil {
		db, err = sql.Open("mysql", mysql)
	}

	return db, err
}
