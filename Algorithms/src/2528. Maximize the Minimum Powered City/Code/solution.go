package main

func maxPower(stations []int, r int, k int) int64 {
    n := len(stations)

    // base power with diff
    diff := make([]int64, n+1)
    for i, v := range stations {
        L := i - r
        if L < 0 {
            L = 0
        }
        R := i + r + 1
        if R > n {
            R = n
        }
        diff[L] += int64(v)
        diff[R] -= int64(v)
    }

    base := make([]int64, n)
    var run int64
    for i := 0; i < n; i++ {
        run += diff[i]
        base[i] = run
    }

    // binary search
    var sum int64
    for _, v := range stations {
        sum += int64(v)
    }
    lo, hi := int64(0), sum+int64(k)
    var ans int64

    can := func(T int64) bool {
        add := make([]int64, n+1)
        var extra, used int64
        for i := 0; i < n; i++ {
            extra += add[i]
            curr := base[i] + extra
            if curr < T {
                need := T - curr
                used += need
                if used > int64(k) {
                    return false
                }
                extra += need
                end := i + 2*r + 1
                if end > n {
                    end = n
                }
                add[end] -= need
            }
        }
        return true
    }

    for lo <= hi {
        mid := (lo + hi) / 2
        if can(mid) {        // ← THIS WAS THE syntax fix
            ans = mid
            lo = mid + 1
        } else {
            hi = mid - 1
        }
    }
    return ans
}
