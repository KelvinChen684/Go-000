package data

import (
	"database/sql"
	"github.com/KelvinChen684/Go-000/tree/main/Week04/internal/biz"
	"github.com/gomodule/redigo/redis"
)

var _ biz.UserData = (biz.UserData)(nil)

type data struct {
	db *sql.DB
	redis redis.Conn
}

func NewData(db *sql.DB, redis redis.Conn) (d biz.UserData, err error)  {
	return newData(db, redis)
}

func newData(db *sql.DB, redis redis.Conn) (d *data, err error)  {
	d = &data{
		db: db,
		redis: redis,
	}
	return
}

func (d *data) Query(user *biz.User) (ug *biz.UserGender, err error)  {
	var gender interface{}
	if gender, err = d.redis.Do("Get", user.Name); err == nil {
		ug.Name = user.Name
		ug.Gender = gender.(string)
		return
	}

	if err = d.db.QueryRow("SELECT name, gender FROM users WHERE user = $1", user.Name).
		Scan(&ug.Name, &ug.Gender); err != nil {
			return nil, err
	}
	return ug, nil
}

