func removeInvalidParentheses(s string) []string {
    left, right := 0, 0
    for _, c := range s {
        if c == '(' {
            left++
        } else if c == ')' {
            if left > 0 {
                left--
            } else {
                right++
            }
        }
    }
    res := make(map[string]struct{})
    path := make([]byte, 0, len(s))
    var dfs func(i, leftRem, rightRem, open int)
    dfs = func(i, leftRem, rightRem, open int) {
        if i == len(s) {
            if leftRem == 0 && rightRem == 0 && open == 0 {
                res[string(path)] = struct{}{}
            }
            return
        }
        c := s[i]
        if c != '(' && c != ')' {
            path = append(path, c)
            dfs(i+1, leftRem, rightRem, open)
            path = path[:len(path)-1]
            return
        }
        if c == '(' {
            if leftRem > 0 {
                dfs(i+1, leftRem-1, rightRem, open)
            }
            path = append(path, c)
            dfs(i+1, leftRem, rightRem, open+1)
            path = path[:len(path)-1]
        } else {
            if rightRem > 0 {
                dfs(i+1, leftRem, rightRem-1, open)
            }
            if open > 0 {
                path = append(path, c)
                dfs(i+1, leftRem, rightRem, open-1)
                path = path[:len(path)-1]
            }
        }
    }
    dfs(0, left, right, 0)
    ans := make([]string, 0, len(res))
    for k := range res {
        ans = append(ans, k)
    }
    return ans
}