package main

import "fmt"

func main() {
	for _, num := range []int{121, -121, 10, 123454321} {
		fmt.Println(isPalindrome(num))
	}
}

func isPalindrome(x int) bool {
	negetive := false
	var num []byte

	if x < 0 {
		x = -x
		negetive = true
	}

	for x > 0 {
		if negetive {
			num = append(num, byte('-'))
			negetive = false
		}
		num = append(num, byte(x%10))
		x = x / 10
	}
	fmt.Println(num)
	for i := 0; i < len(num)/2; i++ {
		if num[i] != num[len(num)-1-i] {
			return false
		}
	}
	return true
}
