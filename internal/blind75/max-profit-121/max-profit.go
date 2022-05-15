package max_profit_121

func maxProfit(prices []int) (profit int) {
	min, max := prices[0], prices[0]
	for i := range prices {
		if prices[i] > max {
			max = prices[i]
		}
		if prices[i] < min {
			min, max = prices[i], prices[i]
		}
		if max-min > profit {
			profit = max - min
		}
	}

	return profit
}
