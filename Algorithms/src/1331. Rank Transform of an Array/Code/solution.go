func arrayRankTransform(arr []int) []int {
	// Create a copy so the original order is preserved
	sorted := append([]int(nil), arr...)

	// Sort the copied array
	sort.Ints(sorted)

	// Store value -> rank
	rank := make(map[int]int)
	currentRank := 1

	// Assign ranks only to unique values
	for _, num := range sorted {
		if _, exists := rank[num]; !exists {
			rank[num] = currentRank
			currentRank++
		}
	}

	// Replace every element with its assigned rank
	for i, num := range arr {
		arr[i] = rank[num]
	}

	// Return the transformed array
	return arr
}