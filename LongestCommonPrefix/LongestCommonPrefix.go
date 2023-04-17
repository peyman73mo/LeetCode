package main

import (
	"fmt"
	"regexp"
)

func main() {

	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(longestCommonPrefix([]string{"ab", "a"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	check_list := find_shortest_words(strs)
	tmp := ""
	// check only in shortest words
	for _, str := range check_list {
		// creating pattern

		tmp_pattern := ""
		i := 0
		for j := i; j < len(strs[str]); j++ {

			r, _ := regexp.Compile("^" + strs[str][i:j+1])

			if check_pattern(r, strs) {
				tmp_pattern = strs[str][i : j+1]
			} else {
				break
			}
		}
		if len(tmp_pattern) > len(tmp) {
			tmp = tmp_pattern
		}

	}

	return tmp
}

func find_shortest_words(strs []string) []int {
	//  this function finds shortest words in given list to check for pattern
	var tmp_list = []int{0}
	tmp := len(strs[0])

	for i := 1; i < len(strs); i++ {

		if len(strs[i]) == tmp {
			tmp_list = append(tmp_list, i)
		} else if len(strs[i]) < tmp {
			tmp_list = []int{i}
			tmp = len(strs[i])
		}
	}
	return tmp_list
}

func check_pattern(pattern *regexp.Regexp, strs []string) bool {
	// this function checks given pattern in all words in "strs"
	for _, word := range strs {
		if !pattern.MatchString(word) {
			return false
		}
	}
	return true
}
