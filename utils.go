package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func errCheck(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %+v", msg, err)
		panic(err)
	}
}

//isNumericIgnore(s string, ig string, x int)
func isNumericIgnore(s string, ig string, x int) bool {
	s = strings.Replace(s, ig, "", x)
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

//isNumeric(s string)
func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

//toNumericIgnore(s string, ig string, x int) (float64, error)
func toNumericIgnore(s string, ig string, x int) (float64, error) {
	s = strings.Replace(s, ig, "", x)
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

//toNumeric(s string) (float64, error)
func toNumeric(s string) (float64, error) {
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

//noNumbers : test if string contains any numbers
var noNumbers = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

//test if config is empty
func isStructEmpty(x interface{}) bool {
	v := reflect.ValueOf(x)

	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i).Interface()
		if value == "" || value == nil {
			return true
		}
	}
	return false
}
