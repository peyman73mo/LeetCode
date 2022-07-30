package main

import "fmt"

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(1534236469))

}

func reverse(x int) int {
	var result int

	for x != 0 {
		result = result*10 + x%10
		x = x / 10
	}
	fmt.Println("result : ", int32(result))
	return result
}
