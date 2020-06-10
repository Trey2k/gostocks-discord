package main

import (
	"fmt"
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
	return i, err
}

//toNumeric(s string) (float64, error)
func toNumeric(s string) (float64, error) {
	i, err := strconv.ParseFloat(s, 64)
	return i, err
}

var isLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
