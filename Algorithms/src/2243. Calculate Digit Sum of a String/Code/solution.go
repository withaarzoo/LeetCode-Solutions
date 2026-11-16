func digitSum(s string, k int) string {
    for len(s) > k {
        next := ""
        for i := 0; i < len(s); i += k {
            sum := 0
            for j := i; j < len(s) && j < i+k; j++ {
                sum += int(s[j] - '0')
            }
            next += strconv.Itoa(sum)
        }
        s = next
    }
    return s
}