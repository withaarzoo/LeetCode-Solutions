func wordSubsets(words1 []string, words2 []string) []string {
    maxFreq := make([]int, 26)

    // Precompute the maximum frequency for each character in words2
    for _, word := range words2 {
        freq := make([]int, 26)
        for _, char := range word {
            freq[char-'a']++
        }
        for i := 0; i < 26; i++ {
            if freq[i] > maxFreq[i] {
                maxFreq[i] = freq[i]
            }
        }
    }

    result := []string{}
    // Check each word in words1
    for _, word := range words1 {
        freq := make([]int, 26)
        for _, char := range word {
            freq[char-'a']++
        }
        isUniversal := true
        for i := 0; i < 26; i++ {
            if freq[i] < maxFreq[i] {
                isUniversal = false
                break
            }
        }
        if isUniversal {
            result = append(result, word)
        }
    }

    return result
}
