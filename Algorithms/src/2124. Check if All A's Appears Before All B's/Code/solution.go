func checkString(s string) bool {
    seenB := false
    for _, c := range s {
        if c == 'b' {
            seenB = true
        } else if seenB {
            return false
        }
    }
    return true
}