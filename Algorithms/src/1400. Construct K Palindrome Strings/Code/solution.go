func canConstruct(s string, k int) bool {
    if k > len(s) {
        return false // More palindromes than characters
    }
    freq := make([]int, 26) // Frequency array for lowercase letters
    for _, char := range s {
        freq[char-'a']++
    }
    oddCount := 0
    for _, count := range freq {
        if count%2 != 0 {
            oddCount++
        }
    }
    return oddCount <= k
}
