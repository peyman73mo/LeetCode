package main

import "fmt"

func main() {

	fmt.Println(longestPalindrome("asdalakjd"))
}

func longestPalindrome(s string) string {
	var buffer = ""

	for index, _ := range s {
		for i := 0; i < len(s)-index; i++ {
			if check(s[index : len(s)-i]) {
				if len(buffer) < len(s[index:len(s)-i]) {
					buffer = s[index : len(s)-i]
				}
			}
		}
	}

	return buffer
}

func check(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
