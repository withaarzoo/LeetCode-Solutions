func numOfStrings(patterns []string, word string) int {

    // Store the number of matching patterns
    count := 0

    // Check every pattern
    for _, pattern := range patterns {

        // strings.Contains() checks whether pattern exists in word
        if strings.Contains(word, pattern) {
            count++
        }
    }

    // Return the final answer
    return count
}