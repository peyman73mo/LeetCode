package main

func main() {

}

func findWords(board [][]byte, words []string) []string {

	result := []string{}
	for _, word := range words {
		if check(board, word) {
			result = append(result, word)
		}
	}

	return result
}

func check(board [][]byte, word string) bool {
	pre_indexes := []int{-1, -1}
	for i := 0; i < len(word); i++ {
		if findchar(word[i], pre_indexes, board) != true {
			return false
		}
	}
	return true
}

func findchar(char byte, pre_index []int, board [][]byte) bool {
	if pre_index[0] == -1 {
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[i]); j++ {
				if board[i][j] == char {
					pre_index[0] = i
					pre_index[1] = j
					return true
				}
			}
		}
	} else {
		adjacents := [][]int{{}}

	}

	return false
}
