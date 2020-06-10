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

func isNumeric(s string) bool {
	s = strings.Replace(s, "@", "", 1)
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
