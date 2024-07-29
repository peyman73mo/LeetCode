package main

func maxProfit(prices []int) int {
	var invested bool
	var buy, profit int

	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] && !invested {
			invested = true
			buy = prices[i]
		} else if prices[i] > prices[i+1] && invested {
			profit += prices[i] - buy
			invested = false
		}
	}
	if invested {
		profit += prices[len(prices)-1] - buy
	}

	return profit
}

func main() {
	// prices := []int{7, 1, 5, 3, 6, 4}
	// prices := []int{1, 2, 3, 4, 5}
	prices := []int{7, 6, 4, 3, 1}
	println(maxProfit(prices))
}
