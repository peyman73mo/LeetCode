package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(validIPAddress("2001:0db8:85a3:0:0:8A2E:0370:7334"))
}

func validIPAddress(queryIP string) string {

	if strings.Contains(queryIP, ".") {

		if len(strings.Split(queryIP, ".")) == 4 && IPv4Checker(queryIP) {
			return "IPv4"
		}
	} else if strings.Contains(queryIP, ":") {
		if len(strings.Split(queryIP, ":")) == 8 && IPv6Cheker(queryIP) {
			return "IPv6"
		}
	}

	return "Neither"
}

func IPv4Checker(ip string) bool {
	regstr := `([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`

	re, _ := regexp.Compile(`^(` + regstr + `\.){3}` + regstr + `$`)
	if re.MatchString(ip) {
		return true
	} else {
		return false
	}

}
func IPv6Cheker(ip string) bool {
	regstr := `([a-fA-F0-9])`
	re, _ := regexp.Compile(`^(` + regstr + `{1,4}\:){7}` + regstr + `{1,4}$`)
	if re.MatchString(ip) {
		return true
	} else {
		return false
	}
}
