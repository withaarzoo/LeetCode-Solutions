func maxNumOfSubstrings(s string) []string {
    n := len(s)
    // first and last occurrence of each letter
    left := make([]int, 26)
    right := make([]int, 26)
    for i := range left {
        left[i] = -1
        right[i] = -1
    }
    for i := 0; i < n; i++ {
        c := s[i] - 'a'
        if left[c] == -1 {
            left[c] = i   // record first time we see it
        }
        right[c] = i      // always update last time
    }

    // collect every valid closed interval
    type interval struct{ L, R int }
    var intervals []interval
    for i := 0; i < 26; i++ {
        if left[i] == -1 {
            continue // letter never appeared
        }
        L, R := left[i], right[i]
        valid := true
        // expand right end while scanning the current range
        for j := L; j <= R; j++ {
            c := s[j] - 'a'
            if left[c] < L { // needs to start earlier -> not closed
                valid = false
                break
            }
            if right[c] > R {
                R = right[c] // push right end if needed
            }
        }
        if valid {
            intervals = append(intervals, interval{L, R})
        }
    }

    // sort by ending position so greedy can pick earliest-ending first
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i].R < intervals[j].R
    })

    // greedy selection of non-overlapping intervals
    var ans []string
    lastEnd := -1
    for _, iv := range intervals {
        if iv.L > lastEnd { // completely after previous piece
            ans = append(ans, s[iv.L:iv.R+1])
            lastEnd = iv.R
        }
    }
    return ans
}