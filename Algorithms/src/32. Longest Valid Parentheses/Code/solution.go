func longestValidParentheses(s string) int {
    left, right, maxLen := 0, 0, 0
    for _, c := range s {
        if c == '(' {
            left++
        } else {
            right++
        }
        if left == right {
            if 2*right > maxLen {
                maxLen = 2 * right
            }
        } else if right > left {
            left, right = 0, 0
        }
    }
    left, right = 0, 0
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == '(' {
            left++
        } else {
            right++
        }
        if left == right {
            if 2*left > maxLen {
                maxLen = 2 * left
            }
        } else if left > right {
            left, right = 0, 0
        }
    }
    return maxLen
}