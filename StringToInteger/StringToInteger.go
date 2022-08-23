package main

import "fmt"

func main() {
	fmt.Println(myAtoi("  678"))
	fmt.Println(myAtoi("  678 sdlekfus"))
	fmt.Println(myAtoi("  -678"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi("+-1"))
	fmt.Println(myAtoi("-+1"))
	fmt.Println(myAtoi("00000-42a1234"))
	fmt.Println(myAtoi("9223372036854775808"))
}

func myAtoi(s string) int {
	str := []byte(s)
	num := 0
	sign := 1
	flag := false

	if len(str) == 0 {
		return 0
	}
	for _, value := range str {
		if value == 32 && !flag {
			continue
		} else if (value > 57 || value < 48) && value != 45 && value != 43 {
			break
		}

		if value == 45 && !flag {
			flag = true
			sign = -1
		} else if value == 43 && !flag {
			flag = true
			sign = 1
		} else if flag && (value == 45 || value == 43) {
			break
		} else {
			flag = true
			num = num*10 + int(value-48)
		}
		if num > 2147483647 || num < -2147483648 {
			if sign == 1 {
				return 2147483647
			} else {
				return -2147483648
			}
		}
	}

	return num * sign
}
