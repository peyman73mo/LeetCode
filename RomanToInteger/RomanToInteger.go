package main

func romanToInt(s string) int {
	var roman_numbers map[string]int = map[string]int{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}
	result := 0

	for i := 0; i < len(s); {
		if (i + 1) < len(s) {
			if roman_numbers[string(s[i])] < roman_numbers[string(s[i+1])] {
				result += roman_numbers[string(s[i])+string(s[i+1])]
				i += 2
			} else {
				result += roman_numbers[string(s[i])]
				i++
			}

		} else {
			result += roman_numbers[string(s[i])]
			i++
		}
	}

	return result
}
