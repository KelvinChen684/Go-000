package data

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
)

func NewRedis() (r redis.Conn, err error) {
	type option struct {
		Network string	`json:"network, required"`
		Addr    string `json:"addr, required"`
	}
	var opt option
	var jsonFile *os.File
	if jsonFile, err = os.Open("/../configs/redis.json"); err != nil {
		err = errors.Wrapf(errors.New("json config error"), "%s", err.Error())
		return nil, err
	}
	defer jsonFile.Close()

	byteVal, _ := ioutil.ReadAll(jsonFile)
	if err = json.Unmarshal([]byte(byteVal), &opt); err != nil {
		err = errors.Wrapf(errors.New("json config error"), "%s", err.Error())
		return nil, err
	}

	if r, err = redis.Dial(opt.Network, opt.Addr); err != nil {
		err = errors.Wrapf(errors.New("redis connect error"), "%s", err.Error())
		return nil, err
	}
	return r, nil
}
