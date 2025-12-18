func maxProfit(prices []int, strategy []int, k int) int64 {
	n := len(prices)

	var base int64 = 0
	for i := 0; i < n; i++ {
		base += int64(strategy[i] * prices[i])
	}

	prefPrice := make([]int64, n+1)
	prefProfit := make([]int64, n+1)

	for i := 0; i < n; i++ {
		prefPrice[i+1] = prefPrice[i] + int64(prices[i])
		prefProfit[i+1] = prefProfit[i] + int64(strategy[i]*prices[i])
	}

	var bestDelta int64 = 0
	half := k / 2

	for l := 0; l+k <= n; l++ {
		m := l + half
		r := l + k

		oldProfit := prefProfit[r] - prefProfit[l]
		newProfit := prefPrice[r] - prefPrice[m]

		if newProfit-oldProfit > bestDelta {
			bestDelta = newProfit - oldProfit
		}
	}

	return base + bestDelta
}
