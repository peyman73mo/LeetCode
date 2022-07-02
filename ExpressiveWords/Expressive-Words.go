package main

import "fmt"

func main() {
	fmt.Println(expressiveWords("heeellooo", []string{"hello", "hi", "helo"}))
	fmt.Println(expressiveWords("zzzzzyyyyy", []string{"zzyy", "zy", "zyy"}))
	fmt.Println(expressiveWords("dddiiiinnssssssoooo", []string{"dinnssoo", "ddinso", "ddiinnso", "ddiinnssoo", "ddiinso", "dinsoo", "ddiinsso", "dinssoo", "dinso"}))
	fmt.Println(expressiveWords("lee", []string{"le"}))
}

func expressiveWords(s string, words []string) int {
	count := 0
	split_s := splitWord(s)
	// fmt.Println("checking ", split_s)
	for _, word := range words {
		if isExpressive(split_s, word) {
			fmt.Println(word)
			count++
		}
	}

	return count
}
func isExpressive(s []string, word string) bool {
	split_word := splitWord(word)
	if len(s) != len(split_word) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if len(s[i]) < len(split_word[i]) || s[i][0] != split_word[i][0] {
			// fmt.Println("[1] : ", split_word)
			return false
		}
		if len(split_word[i]) == 1 && len(s[i]) != len(split_word[i]) && len(s[i]) < 3 {
			// fmt.Println("[2] : ", split_word)
			return false
		}

	}

	return true
}

func splitWord(word string) []string {
	f_index := 0
	split_word := []string{}
	for i := 0; i < len(word)-1; i++ {
		if word[i] != word[i+1] {
			split_word = append(split_word, word[f_index:i+1])
			f_index = i + 1
		}
	}
	if f_index < len(word) {
		split_word = append(split_word, word[f_index:])
	}
	return split_word
}
