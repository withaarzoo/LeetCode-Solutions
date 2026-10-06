func minAddToMakeValid(s string) int {
    open, add := 0, 0
    for _, c := range s {
        if c == '(' {
            open++
        } else {
            if open > 0 {
                open--
            } else {
                add++
            }
        }
    }
    return add + open
}