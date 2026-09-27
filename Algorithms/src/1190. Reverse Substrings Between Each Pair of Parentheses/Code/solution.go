func reverseParentheses(s string) string {
    n := len(s)
    pair := make([]int, n)
    st := []int{}
    for i := 0; i < n; i++ {
        if s[i] == '(' {
            st = append(st, i)
        } else if s[i] == ')' {
            j := st[len(st)-1]
            st = st[:len(st)-1]
            pair[i] = j
            pair[j] = i
        }
    }
    var res []byte
    i, dir := 0, 1
    for i >= 0 && i < n {
        if s[i] == '(' || s[i] == ')' {
            i = pair[i]
            dir = -dir
        } else {
            res = append(res, s[i])
        }
        i += dir
    }
    return string(res)
}