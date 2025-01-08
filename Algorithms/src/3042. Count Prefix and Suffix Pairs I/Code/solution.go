func countPrefixSuffixPairs(words []string) int {
    count := 0
    n := len(words)

    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            prefix := words[i]
            word := words[j]
            if len(prefix) <= len(word) && 
               strings.HasPrefix(word, prefix) && 
               strings.HasSuffix(word, prefix) {
                count++
            }
        }
    }

    return count
}
