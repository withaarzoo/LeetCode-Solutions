func minTimeToVisitAllPoints(points [][]int) int {
    totalTime := 0

    for i := 1; i < len(points); i++ {
        dx := abs(points[i][0] - points[i-1][0])
        dy := abs(points[i][1] - points[i-1][1])

        if dx > dy {
            totalTime += dx
        } else {
            totalTime += dy
        }
    }

    return totalTime
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
