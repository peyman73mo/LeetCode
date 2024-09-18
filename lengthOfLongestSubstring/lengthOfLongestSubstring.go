package main

func lengthOfLongestSubstring(s string) int {
	var max_len int
	var tmp int = 0

	for index := 0; index < len(s); index++ {
		var checked_map = make(map[byte]int)
		max_len = 0
		for i := index; i < len(s); i++ {
			if checked_map[s[i]] == 0 {
				checked_map[s[i]] = 1
				max_len++
			} else {
				break
			}
		}
		if tmp < max_len {
			tmp = max_len
		}
	}

	return tmp
}

func main() {
	println(lengthOfLongestSubstring("abcabcbb"))
	println(lengthOfLongestSubstring("bbbbb"))
	println(lengthOfLongestSubstring("pwwkew"))
	println(lengthOfLongestSubstring("dvdf"))
	println(lengthOfLongestSubstring("asjrgapa"))
	println(lengthOfLongestSubstring("abcabcbb"))
}
