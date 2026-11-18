func minChanges(s string) int {
    changes := 0
    for i := 0; i < len(s); i += 2 {
        if s[i] != s[i+1] {
            changes++
        }
    }
    return changes
}