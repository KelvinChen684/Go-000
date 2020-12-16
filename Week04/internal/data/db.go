package data

import (
	"database/sql"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	_ "github.com/lib/pq"
	"os"
	"fmt"
)

func NewDB() (db *sql.DB, err error)  {
	type option struct {
		driver string	`json:"driver, required"`
		db    string `json:"db, required"`
		host string `json:"host"`
		port string `json:"port"`
		user string `json:"user"`
		sslmode string `json:"sslmode"`
		password string `json:"password, required"`
	}

	var opt option
	var jsonFile *os.File
	if jsonFile, err = os.Open("../configs/postgre.json"); err != nil {
		err = errors.Wrapf(errors.New("postgre database config error"), "%s", err.Error())
		return nil, err
	}
	defer jsonFile.Close()

	byteVal, _ := ioutil.ReadAll(jsonFile)
	if err = json.Unmarshal([]byte(byteVal), &opt); err != nil {
		err = errors.Wrapf(errors.New("postgre database config error"), "%s", err.Error())
		return nil, err
	}

	pgconfig := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		opt.host, opt.port, opt.user, opt.password, opt.db, opt.sslmode)
	if db, err = sql.Open(opt.driver, pgconfig); err != nil {
		err = errors.Wrapf(errors.New("postgre database connect error"), "%s", err.Error())
		return nil, err
	}
	return db, nil
}
