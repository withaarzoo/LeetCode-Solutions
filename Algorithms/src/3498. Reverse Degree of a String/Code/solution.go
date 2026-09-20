func reverseDegree(s string) int {
    ans := 0
    for i, c := range s {
        ans += (26 - int(c - 'a')) * (i + 1)
    }
    return ans
}