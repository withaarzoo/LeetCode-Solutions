func maxSumOfSquares(num int, sum int) string {
    if sum > 9*num {
        return ""
    }
    q := sum / 9
    r := sum % 9
    b := make([]byte, 0, num)
    for i := 0; i < q; i++ {
        b = append(b, '9')
    }
    if r > 0 {
        b = append(b, byte('0'+r))
    }
    for i := q + map[bool]int{true: 1, false: 0}[r > 0]; i < num; i++ {
        b = append(b, '0')
    }
    return string(b)
}