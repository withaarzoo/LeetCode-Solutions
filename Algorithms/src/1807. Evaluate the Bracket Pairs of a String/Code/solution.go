func evaluate(s string, knowledge [][]string) string {
    mp := make(map[string]string, len(knowledge))
    for _, p := range knowledge {
        mp[p[0]] = p[1]
    }
    var res strings.Builder
    n := len(s)
    for i := 0; i < n; {
        if s[i] == '(' {
            j := i + 1
            for s[j] != ')' {
                j++
            }
            key := s[i+1 : j]
            if v, ok := mp[key]; ok {
                res.WriteString(v)
            } else {
                res.WriteByte('?')
            }
            i = j + 1
        } else {
            res.WriteByte(s[i])
            i++
        }
    }
    return res.String()
}