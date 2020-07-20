package utils

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

//Init init utils, must be ran before anything that uses the config
func Init() {
	var err error

	if _, err := os.Stat("./logs"); os.IsNotExist(err) {
		os.Mkdir("./logs", 0700)
	}

	Config, err = GetConfig()
	ErrCheck("Error getting config", err)

	if IsStructEmpty(Config.Discord) {
		Log("A value in config.Discord is empty", LogError)
		panic("Fatal")
	}

	if IsStructEmpty(Config.TD) {
		Log("A value in config.TD is empty", LogError)
		panic("Fatal")
	}
}

//ErrCheck check if there is an error
func ErrCheck(msg string, err error) {
	if err != nil {
		Log(fmt.Sprintf("%s: %+v", msg, errors.WithStack(err)), LogError)
	}
}

//IsNumericIgnore (s string, ig string, x int)
func IsNumericIgnore(str string, ignore string, count int) bool {
	str = strings.Replace(str, ignore, "", count)
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}

//IsNumeric (s string)
func IsNumeric(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}

//ToNumericIgnore (s string, ig string, x int) (float64, error)
func ToNumericIgnore(str string, ignore string, count int) (float64, error) {
	strFilterd := strings.Replace(str, ignore, "", count)
	number, err := strconv.ParseFloat(strFilterd, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}

//ToIntIgnore (s string, ig string, x int) (float64, error)
func ToIntIgnore(str string, ignore string, count int) (int64, error) {
	strFilterd := strings.Replace(str, ignore, "", count)
	number, err := strconv.ParseInt(strFilterd, 10, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}

//ToNumeric (s string) (float64, error)
func ToNumeric(str string) (float64, error) {
	number, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}

//NoNumbers : test if string contains any numbers
var NoNumbers = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

//IsStructEmpty test if config is empty
func IsStructEmpty(x interface{}) bool {
	v := reflect.ValueOf(x)

	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i).Interface()
		if value == "" || value == nil {
			return true
		}
	}
	return false
}

//FileExists check if a file exists
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

//PrintOrder nicley print order
func PrintOrder(order OrderStruct) string {
	return fmt.Sprintf("----------------------------------------------------------------------------------------------------------------------------------\n"+
		"Buy: %v, Ticker: %s, Date: %s, StrikerPrice: %v, ContractType: %s, Alerted Price: %v, Risky: %v, Stop Loss: %v\n"+
		"----------------------------------------------------------------------------------------------------------------------------------\n",
		order.Buy, order.Ticker, order.ExpDate.Format("1/2/2006"), order.StrikPrice, order.ContractType, order.Price, order.Risky, order.StopLoss)
}

//InTimeSpan Stuff
func InTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

const (
	_ = iota
	LogError
	LogOrder
	LogInfo
)

//Log log events
func Log(msg string, logType uint) {
	switch logType {
	case LogError:
		f, _ := os.OpenFile("./logs/error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0700)
		defer f.Close()
		log.SetOutput(f)
		log.Printf("Error: %s\n", msg)
		fmt.Printf("Error: %s\n", msg)
	case LogOrder:
		f, _ := os.OpenFile("./logs/info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0700)
		defer f.Close()
		log.SetOutput(f)
		log.Printf("New order:\n%s", msg)
		fmt.Printf("New order:\n%s", msg)
	case LogInfo:
		f, _ := os.OpenFile("./logs/info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0700)
		defer f.Close()
		log.SetOutput(f)
		log.Printf("Info: %s\n", msg)
		fmt.Printf("Info: %s\n", msg)
	}
}
