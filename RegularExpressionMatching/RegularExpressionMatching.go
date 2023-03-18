package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(is_match("aa", ".*"))
}

func is_match(s, p string) bool {
	pattern := "^" + p + "$"
	re, _ := regexp.Compile(pattern)
	if re.MatchString(s) {
		return true
	} else {
		return false
	}
}
