func braceExpansionII(expression string) []string {
    i := 0
    var parse func() map[string]struct{}
    product := func(a, b map[string]struct{}) map[string]struct{} {
        res := make(map[string]struct{})
        for x := range a {
            for y := range b {
                res[x+y] = struct{}{}
            }
        }
        return res
    }
    parse = func() map[string]struct{} {
        res := make(map[string]struct{})
        cur := map[string]struct{}{"": {}}
        for i < len(expression) && expression[i] != '}' {
            if expression[i] == '{' {
                i++
                next := parse()
                i++
                cur = product(cur, next)
            } else if expression[i] == ',' {
                for s := range cur {
                    res[s] = struct{}{}
                }
                cur = map[string]struct{}{"": {}}
                i++
            } else {
                next := map[string]struct{}{string(expression[i]): {}}
                i++
                cur = product(cur, next)
            }
        }
        for s := range cur {
            res[s] = struct{}{}
        }
        return res
    }
    result := parse()
    ans := make([]string, 0, len(result))
    for s := range result {
        ans = append(ans, s)
    }
    sort.Strings(ans)
    return ans
}