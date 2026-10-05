func scoreOfParentheses(s string) int {
    score, depth := 0, 0
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            depth++
        } else {
            depth--
            if s[i-1] == '(' {
                score += 1 << depth
            }
        }
    }
    return score
}