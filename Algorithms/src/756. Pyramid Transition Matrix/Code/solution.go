func pyramidTransition(bottom string, allowed []string) bool {
    rules := map[string][]byte{}
    for _, s := range allowed {
        rules[s[:2]] = append(rules[s[:2]], s[2])
    }

    bad := map[string]bool{}

    var dfs func(string, int, []byte) bool
    dfs = func(row string, idx int, next []byte) bool {
        if len(row) == 1 {
            return true
        }

        if idx == len(row)-1 {
            nxt := string(next)
            if bad[nxt] {
                return false
            }
            ok := dfs(nxt, 0, nil)
            if !ok {
                bad[nxt] = true
            }
            return ok
        }

        key := row[idx : idx+2]
        for _, c := range rules[key] {
            if dfs(row, idx+1, append(next, c)) {
                return true
            }
        }
        return false
    }

    return dfs(bottom, 0, nil)
}
