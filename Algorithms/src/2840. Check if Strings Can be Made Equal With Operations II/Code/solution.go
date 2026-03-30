func checkStrings(s1 string, s2 string) bool {
    // Frequency arrays for even and odd positions
    even := make([]int, 26)
    odd := make([]int, 26)

    for i := 0; i < len(s1); i++ {
        if i%2 == 0 {
            // Count characters at even indexes
            even[s1[i]-'a']++
            even[s2[i]-'a']--
        } else {
            // Count characters at odd indexes
            odd[s1[i]-'a']++
            odd[s2[i]-'a']--
        }
    }

    // Check if all frequencies become zero
    for i := 0; i < 26; i++ {
        if even[i] != 0 || odd[i] != 0 {
            return false
        }
    }

    return true
}