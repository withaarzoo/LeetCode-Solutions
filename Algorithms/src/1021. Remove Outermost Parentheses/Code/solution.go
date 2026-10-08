func removeOuterParentheses(s string) string {
    res := make([]byte, 0, len(s))
    bal := 0
    for i := 0; i < len(s); i++ {
        c := s[i]
        if c == '(' {
            if bal > 0 {
                res = append(res, c)
            }
            bal++
        } else {
            bal--
            if bal > 0 {
                res = append(res, c)
            }
        }
    }
    return string(res)
}