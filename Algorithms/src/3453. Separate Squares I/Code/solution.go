func separateSquares(squares [][]int) float64 {
    totalArea := 0.0
    low, high := 1e18, -1e18

    for _, s := range squares {
        y := float64(s[1])
        l := float64(s[2])
        totalArea += l * l
        if y < low {
            low = y
        }
        if y+l > high {
            high = y + l
        }
    }

    for i := 0; i < 80; i++ {
        mid := (low + high) / 2
        areaBelow := 0.0

        for _, s := range squares {
            y := float64(s[1])
            l := float64(s[2])

            if mid <= y {
                continue
            } else if mid >= y+l {
                areaBelow += l * l
            } else {
                areaBelow += l * (mid - y)
            }
        }

        if areaBelow*2 < totalArea {
            low = mid
        } else {
            high = mid
        }
    }
    return low
}
