package main

import (
	"fmt"
)

func lengthOfLongestSubstring(str string) int {

	lenght := 0

	for i := 0; i < len(str); i++ {
		arr := make(map[string]int)
		tmp := 0
		for _, char := range str[i:] {
			if _, ok := arr[string(char)]; ok {
				break
			} else {
				arr[string(char)] = int(char)
				tmp += 1
			}
		}
		if tmp > lenght {
			lenght = tmp
		}
	}
	return lenght
}

func main() {
	fmt.Println(lengthOfLongestSubstring("au"))

}
