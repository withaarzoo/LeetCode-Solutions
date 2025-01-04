func countPalindromicSubsequence(s string) int {
    first := make([]int, 26)
    last := make([]int, 26)
    for i := range first {
        first[i] = -1
    }
    
    for i, char := range s {
        index := int(char - 'a')
        if first[index] == -1 {
            first[index] = i
        }
        last[index] = i
    }
    
    result := 0
    for i := 0; i < 26; i++ {
        if first[i] != -1 && last[i] > first[i] {
            middleChars := make(map[byte]bool)
            for j := first[i] + 1; j < last[i]; j++ {
                middleChars[s[j]] = true
            }
            result += len(middleChars)
        }
    }
    
    return result
}
