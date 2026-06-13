func mapWordWeights(words []string, weights []int) string {
	var result []byte

	// Process each word
	for _, word := range words {
		sumWeight := 0

		// Add weights of all characters
		for _, ch := range word {
			sumWeight += weights[ch-'a']
		}

		// Reduce into range [0, 25]
		value := sumWeight % 26

		// Reverse alphabet mapping:
		// 0 -> z, 1 -> y, ..., 25 -> a
		result = append(result, byte('z'-value))
	}

	return string(result)
}