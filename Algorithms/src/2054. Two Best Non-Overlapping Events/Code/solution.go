func maxTwoEvents(events [][]int) int {
    // Sort by start time
    sort.Slice(events, func(i, j int) bool {
        return events[i][0] < events[j][0]
    })

    // Sort by end time
    endSorted := make([][]int, len(events))
    copy(endSorted, events)
    sort.Slice(endSorted, func(i, j int) bool {
        return endSorted[i][1] < endSorted[j][1]
    })

    n := len(events)
    maxValueTill := make([]int, n)

    maxValueTill[0] = endSorted[0][2]
    for i := 1; i < n; i++ {
        maxValueTill[i] = max(maxValueTill[i-1], endSorted[i][2])
    }

    ans := 0
    j := 0

    for i := 0; i < n; i++ {
        start := events[i][0]
        value := events[i][2]

        for j < n && endSorted[j][1] < start {
            j++
        }

        ans = max(ans, value)
        if j > 0 {
            ans = max(ans, value+maxValueTill[j-1])
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
