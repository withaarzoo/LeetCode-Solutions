func maxDistance(colors []int) int {
    n := len(colors)
    ans := 0

    // Check farthest house from the first house
    for i := n - 1; i >= 0; i-- {
        if colors[i] != colors[0] {
            if i > ans {
                ans = i
            }
            break
        }
    }

    // Check farthest house from the last house
    for i := 0; i < n; i++ {
        if colors[i] != colors[n-1] {
            distance := n - 1 - i
            if distance > ans {
                ans = distance
            }
            break
        }
    }

    return ans
}