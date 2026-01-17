func largestSquareArea(bottomLeft [][]int, topRight [][]int) int64 {
    n := len(bottomLeft)
    var ans int64 = 0

    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {

            left   := max(bottomLeft[i][0], bottomLeft[j][0])
            right  := min(topRight[i][0], topRight[j][0])
            bottom := max(bottomLeft[i][1], bottomLeft[j][1])
            top    := min(topRight[i][1], topRight[j][1])

            if right > left && top > bottom {
                side := int64(min(right-left, top-bottom))
                if side*side > ans {
                    ans = side * side
                }
            }
        }
    }
    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
