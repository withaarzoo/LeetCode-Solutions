func makeFancyString(s string) string {
    result := []byte{}
    for i := 0; i < len(s); i++ {
        n := len(result)
        if n < 2 || !(result[n-1] == s[i] && result[n-2] == s[i]) {
            result = append(result, s[i])
        }
    }
    return string(result)
}