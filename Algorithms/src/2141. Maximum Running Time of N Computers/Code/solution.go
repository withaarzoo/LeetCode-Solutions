package main

func maxRunTime(n int, batteries []int) int64 {
    // Compute total battery time as int64
    var total int64 = 0
    for _, b := range batteries {
        total += int64(b)
    }

    var low int64 = 0
    var high int64 = total / int64(n) // Upper bound on answer

    for low < high {
        mid := low + (high-low+1)/2 // upper mid

        var usable int64 = 0
        for _, b := range batteries {
            // Each battery contributes at most mid minutes
            if int64(b) < mid {
                usable += int64(b)
            } else {
                usable += mid
            }
            if usable >= mid*int64(n) {
                break
            }
        }

        if usable >= mid*int64(n) {
            // mid is feasible
            low = mid
        } else {
            // mid is too large
            high = mid - 1
        }
    }

    return low
}
