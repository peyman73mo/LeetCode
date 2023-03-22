package main

func intToRoman(num int) string {

	roman_numbers := [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	roman_numbers_alphabet := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	result := ""
	for index, value := range roman_numbers {

		remaning_number := num / value
		result += concatenate_string(remaning_number, roman_numbers_alphabet[index])
		num = num % value
	}

	return result
}

func concatenate_string(num int, str string) string {
	result := ""
	for i := 0; i < num; i++ {
		result += str
	}
	return result
}
