func maxOperations(s string) int {
    var ans int64 = 0
    var ones int64 = 0
    n := len(s)
    for i := 0; i < n; i++ {
        if s[i] == '1' {
            ones++
        } else { // s[i] == '0'
            if i > 0 && s[i-1] == '1' {
                ans += ones
            }
        }
    }
    return int(ans)
}
