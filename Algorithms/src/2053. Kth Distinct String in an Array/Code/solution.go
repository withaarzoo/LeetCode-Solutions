func kthDistinct(arr []string, k int) string {
    // Create a map to count the occurrences of each string
    count := make(map[string]int)
    
    // Create a slice to store distinct strings
    distinct := []string{}

    // Loop through each string in the array and count its occurrences
    for _, str := range arr {
        count[str]++
    }

    // Loop through the array again to collect distinct strings in order
    for _, str := range arr {
        if count[str] == 1 {
            // If a string appears exactly once, add it to the distinct slice
            distinct = append(distinct, str)
        }
    }

    // Check if k is within the range of distinct strings
    if k <= len(distinct) {
        // Return the k-th distinct string (1-based index)
        return distinct[k-1]
    } else {
        // If k is out of range, return an empty string
        return ""
    }
}
