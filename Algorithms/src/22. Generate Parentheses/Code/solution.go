func generateParenthesis(n int) []string { 
    ans := []string{}
    cur := []byte{}
    var dfs func(open, close int)
    dfs = func(open, close int) {
        if open == 0 && close == 0 {
            ans = append(ans, string(cur))
            return
        }
        if open > 0 {
            cur = append(cur, '(')
            dfs(open-1, close)
            cur = cur[:len(cur)-1]
        }
        if close > open {
            cur = append(cur, ')')
            dfs(open, close-1)
            cur = cur[:len(cur)-1]
        }
    }
    dfs(n, n)
    return ans
}