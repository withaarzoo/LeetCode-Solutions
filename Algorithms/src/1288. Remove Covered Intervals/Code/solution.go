func removeCoveredIntervals(intervals [][]int) int {
    // Sort by start in ascending order.
    // For equal starts, place the larger end first
    // so the covering interval is processed before the covered one.
    sort.Slice(intervals, func(i, j int) bool {
        if intervals[i][0] == intervals[j][0] {
            return intervals[i][1] > intervals[j][1]
        }
        return intervals[i][0] < intervals[j][0]
    })

    // Every valid end is greater than 0, so -1 is a safe initial value.
    maxEnd := -1

    // This stores how many intervals are not covered.
    remaining := 0

    // Check every interval in the sorted order.
    for _, interval := range intervals {
        // A larger end means no earlier interval can cover this one.
        if interval[1] > maxEnd {
            remaining++

            // Remember the farthest ending point seen so far.
            maxEnd = interval[1]
        }
        // Otherwise, the interval is covered, so I skip it.
    }

    // Return the number of intervals that remain.
    return remaining
} 