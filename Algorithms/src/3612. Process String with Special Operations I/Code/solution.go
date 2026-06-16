func processStr(s string) string {
    // Stores the current result being built
    result := ""

    for _, c := range s {
        // Lowercase letter -> append to result
        if c >= 'a' && c <= 'z' {
            result += string(c)

        // Remove last character if it exists
        } else if c == '*' {
            if len(result) > 0 {
                result = result[:len(result)-1]
            }

        // Duplicate current result
        } else if c == '#' {
            result += result

        // Reverse current result
        } else if c == '%' {
            chars := []rune(result)

            // Two-pointer reversal
            for l, r := 0, len(chars)-1; l < r; l, r = l+1, r-1 {
                chars[l], chars[r] = chars[r], chars[l]
            }

            result = string(chars)
        }
    }

    return result
}