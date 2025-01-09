func prefixCount(words []string, pref string) int {
    count := 0
    for _, word := range words {
        // Check if the prefix matches the start of the word
        if len(word) >= len(pref) && word[:len(pref)] == pref {
            count++
        }
    }
    return count
}
