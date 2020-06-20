package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//SetTradeBal stuff
func SetTradeBal(ballance float64) error {
	if _, err := os.Stat("./info"); os.IsNotExist(err) {
		os.Mkdir("./info", 0700)
	}
	file, err := os.Create("./info/tradebal.fake")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, fmt.Sprint(ballance))
	if err != nil {
		return err
	}
	return file.Sync()
}

//GetTradeBal stuff
func GetTradeBal(ballance float64) (float64, error) {
	if FileExists("./info/tradebal.fake") {
		data, err := ioutil.ReadFile("./info/tradebal.fake")
		if err != nil {
			return 0, err
		}
		return ToNumeric(string(data))
	}
	return ballance, SetTradeBal(ballance)
}
