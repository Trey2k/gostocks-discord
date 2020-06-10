package main

import (
	"fmt"
	"strconv"
	"strings"
)

func errCheck(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %+v", msg, err)
		panic(err)
	}
}

/*s = string, i = ignore char, x = ignore char count
isNumericIgnore(s string, i string, x int)*/
func isNumericIgnore(s string, i string, x int) bool {
	s = strings.Replace(s, i, "", x)
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

//isNumeric(s string)
func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
