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
)

//Init init utils, must be ran before anything that uses the config
func Init() {
	var err error
	Config, err = GetConfig()
	ErrCheck("Error getting config", err)

	if IsStructEmpty(Config.Discord) {
		log.Fatal("A value in config.Discord is empty")
	}

	if IsStructEmpty(Config.TD) {
		log.Fatal("A value in config.TD is empty")
	}
}

//ErrCheck check if there is an error
func ErrCheck(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %+v", msg, err)
		panic(err)
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
func PrintOrder(order OrderStruct) {
	fmt.Println("----------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("Buy: " + fmt.Sprint(order.Buy) + ", Ticker: " + order.Ticker + ", Date: " + order.ExpDate.Format("1/2/2006") + ", StrikerPrice: " + fmt.Sprint(order.StrikPrice) + ", ContractType: " + order.ContractType + ", Buy Price: " + fmt.Sprint(order.Price) + ", Risky: " + fmt.Sprint(order.Risky) + ", Stop Loss: " + fmt.Sprint(order.StopLoss))
	fmt.Println("----------------------------------------------------------------------------------------------------------------------------------")
}

//InTimeSpan Stuff
func InTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}
