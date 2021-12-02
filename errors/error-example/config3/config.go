package config

import (
	"fmt"
	"io/ioutil"
)

type HeaderError struct {
	FaultyHeader string
}

func (e *HeaderError) Error() string {
	return fmt.Sprintf("Bad header. Provided %s, expected : APPCONF", e.FaultyHeader)
}

const fileHeader = "APPCONF"

func Load() (string, error) {
	data, err := ioutil.ReadFile("config.txt")
	if err != nil {
		return "", err
	}
	conf := string(data)
	if len(conf) < 8 {
		return "", &HeaderError{conf}
	}
	if conf[0:8] != fileHeader {
		// return "", errors.New("the config file do not begin by accepted header")
		// return "", fmt.Errorf("the config file do not begin by accepted header")
		return "", &HeaderError{conf[:8]}
	}
	return conf, nil
}
