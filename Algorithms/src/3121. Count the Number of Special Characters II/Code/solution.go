func numberOfSpecialChars(word string) int {
    
    // Store last occurrence of lowercase letters
    lower := make([]int, 26)

    // Store first occurrence of uppercase letters
    upper := make([]int, 26)

    // Initialize arrays with -1
    for i := 0; i < 26; i++ {
        lower[i] = -1
        upper[i] = -1
    }

    // Traverse the string
    for i, ch := range word {

        // If lowercase letter
        if ch >= 'a' && ch <= 'z' {

            // Update last occurrence
            lower[ch-'a'] = i
        } else {

            idx := ch - 'A'

            // Store only first occurrence
            if upper[idx] == -1 {
                upper[idx] = i
            }
        }
    }

    ans := 0

    // Check all letters
    for i := 0; i < 26; i++ {

        // Both lowercase and uppercase must exist
        if lower[i] != -1 && upper[i] != -1 {

            // Lowercase must come before uppercase
            if lower[i] < upper[i] {
                ans++
            }
        }
    }

    return ans
}