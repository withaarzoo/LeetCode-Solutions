func maximizeSquareHoleArea(n int, m int, hBars []int, vBars []int) int {

    getMaxGap := func(bars []int) int {
        sort.Ints(bars)

        maxLen := 1
        curLen := 1

        for i := 1; i < len(bars); i++ {
            if bars[i] == bars[i-1]+1 {
                curLen++
            } else {
                curLen = 1
            }
            if curLen > maxLen {
                maxLen = curLen
            }
        }
        return maxLen
    }

    hGap := getMaxGap(hBars) + 1
    vGap := getMaxGap(vBars) + 1

    side := min(hGap, vGap)
    return side * side
}
