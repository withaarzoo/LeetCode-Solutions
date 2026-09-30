func maxDepthAfterSplit(seq string) []int {
    ans := make([]int, len(seq))
    depth := 0
    for i, ch := range seq {
        if ch == '(' {
            depth++
            ans[i] = depth % 2
        } else {
            ans[i] = depth % 2
            depth--
        }
    }
    return ans
}