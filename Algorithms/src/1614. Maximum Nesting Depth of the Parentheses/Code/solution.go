func maxDepth(s string) int {
    depth, maxDepth := 0, 0
    for _, c := range s {
        if c == '(' {
            depth++
            if depth > maxDepth {
                maxDepth = depth
            }
        } else if c == ')' {
            depth--
        }
    }
    return maxDepth
}