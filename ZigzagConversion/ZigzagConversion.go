package ZigzagConversion

import "fmt"

type Holder struct {
	str_arr  string
	next     *Holder
	previous *Holder
}

func main() {
	fmt.Println(convert("A", 1))
}

func convert(s string, numRows int) string {
	if len(s) == 1 || numRows == 1 {
		return s
	}
	holders := make([]Holder, numRows)

	for i := 0; i < numRows-1; i++ {
		holders[i].next = &holders[i+1]
	}

	for i := numRows - 1; i > 0; i-- {
		holders[i].previous = &holders[i-1]
	}

	// fmt.Println(holders)
	char := 0
	flag := true
	point := &holders[0]
	for len(s) > char {

		if point.next != nil && flag == true {
			point.str_arr = point.str_arr + string(s[char])
			char += 1
			point = point.next
		} else if flag == true {
			flag = false
		}
		if point.previous != nil && flag == false {
			point.str_arr = point.str_arr + string(s[char])
			char += 1
			point = point.previous
		} else if flag == false {
			flag = true
		}
	}
	// fmt.Println(holders)

	word := ""
	for i := 0; i < numRows; i++ {
		word += holders[i].str_arr
	}

	return word
}
