package config

import (
	"fmt"
	"io/ioutil"
)

const fileHeader = "APPCONF"

func Load() (string, error) {
	data, err := ioutil.ReadFile("C:/Users/Administrator/tmp/myHotelApp/config.txt")
	if err != nil {
		return "", err
	}
	conf := string(data)
	if len(conf) < 8 || conf[0:8] != fileHeader {
		// return "", errors.New("the config file do not begin by accepted header")
		return "", fmt.Errorf("the config file do not begin by accepted header")
	}
	return conf, nil
}
