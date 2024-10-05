func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }

    s1Count := make([]int, 26)
    s2Count := make([]int, 26)

    // Initialize counts for s1 and the first window of s2
    for i := 0; i < len(s1); i++ {
        s1Count[s1[i]-'a']++
        s2Count[s2[i]-'a']++
    }

    // Slide the window over s2
    for i := 0; i < len(s2)-len(s1); i++ {
        if match(s1Count, s2Count) {
            return true
        }
        s2Count[s2[i]-'a']--
        s2Count[s2[i+len(s1)]-'a']++
    }

    // Check the last window
    return match(s1Count, s2Count)
}

func match(a, b []int) bool {
    for i := 0; i < 26; i++ {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}