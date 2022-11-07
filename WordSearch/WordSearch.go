package main

func main() {

}

type wordSearch struct {
	board [][]byte
	word  string
}

func exist(board [][]byte, word string) bool {
	row := len(board)
	col := len(board[0])
	ws := wordSearch{
		board: board,
		word:  word,
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if ws.find(i, j, 0) {
				return true
			}
		}
	}

	return false
}

func (ws wordSearch) find(row int, col int, index int) bool {

	if index == len(ws.word) {
		return true
	} else if row < 0 || row > (len(ws.board)-1) || col < 0 || col > (len(ws.board[0])-1) || ws.word[index] != ws.board[row][col] {
		return false
	}
	tmp := ws.board[row][col]
	ws.board[row][col] = 0
	if ws.find(row-1, col, index+1) || ws.find(row+1, col, index+1) || ws.find(row, col-1, index+1) || ws.find(row, col+1, index+1) {
		return true
	}
	ws.board[row][col] = tmp
	return false
}
