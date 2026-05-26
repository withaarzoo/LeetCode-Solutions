func numberOfSpecialChars(word string) int {
    
    // Map used like a set to store characters
    st := make(map[rune]bool)

    // Store every character from the string
    for _, ch := range word {
        st[ch] = true
    }

    // Variable to store answer
    count := 0

    // Check all lowercase English letters
    for ch := 'a'; ch <= 'z'; ch++ {

        // Find corresponding uppercase character
        upper := ch - 'a' + 'A'

        // If both exist, increase answer
        if st[ch] && st[upper] {
            count++
        }
    }

    // Return final count
    return count
}