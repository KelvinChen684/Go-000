package data

import (
	"database/sql"
	"github.com/pkg/errors"
	"log"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func init()  {
	var err error
	Db, err = sql.Open("postgres", "dbname=week01 sslmode=disable password=123456")
	if err != nil {
		log.Fatal(err)	// 数据库连接失败，程序启动失败
	}
	return
}


func UserByEmail(email string) (user string, err error) {
	err = Db.QueryRow("SELECT name FROM users WHERE email = $1", email).Scan(&user)
	if err != nil {
		err = errors.Wrap(err, "sql failed")
		return
	}
	return
}