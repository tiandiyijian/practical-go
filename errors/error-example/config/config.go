package config

import (
	"io/ioutil"
)

func Load() (string, error) {
	data, err := ioutil.ReadFile("C:/Users/Administrator/tmp/myHotelApp/config.txt")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// func Print() {
// 	fmt.Println(string(load()))
// }
