func earliestFinishTime(landStartTime []int, landDuration []int,
	waterStartTime []int, waterDuration []int) int {

	// Large initial value
	ans := int(1e9)

	n := len(landStartTime)
	m := len(waterStartTime)

	// Try every land-water pair
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {

			// Land -> Water

			// Finish time of land ride
			landFinish := landStartTime[i] + landDuration[i]

			// Water ride starts after land ride finishes
			// and after it becomes available
			waterStart := max(landFinish, waterStartTime[j])

			// Final finish time
			finish1 := waterStart + waterDuration[j]

			// Water -> Land

			// Finish time of water ride
			waterFinish := waterStartTime[j] + waterDuration[j]

			// Land ride starts after water ride finishes
			// and after it becomes available
			landStart := max(waterFinish, landStartTime[i])

			// Final finish time
			finish2 := landStart + landDuration[i]

			// Keep minimum answer
			ans = min(ans, min(finish1, finish2))
		}
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}