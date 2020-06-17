package utils

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

//ErrCheck check if there is an error
func ErrCheck(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %+v", msg, err)
		panic(err)
	}
}

//IsNumericIgnore (s string, ig string, x int)
func IsNumericIgnore(s string, ig string, x int) bool {
	s = strings.Replace(s, ig, "", x)
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

//IsNumeric (s string)
func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

//ToNumericIgnore (s string, ig string, x int) (float64, error)
func ToNumericIgnore(s string, ig string, x int) (float64, error) {
	s = strings.Replace(s, ig, "", x)
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

//ToNumeric (s string) (float64, error)
func ToNumeric(s string) (float64, error) {
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
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
