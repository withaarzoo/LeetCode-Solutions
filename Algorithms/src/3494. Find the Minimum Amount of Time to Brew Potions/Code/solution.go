package main

import (
    // "fmt"
)

func minTime(skill []int, mana []int) int64 {
    n := len(skill)
    m := len(mana)
    if m == 0 {
        return 0
    }

    // prefix sums (int64)
    pref := make([]int64, n)
    for i := 0; i < n; i++ {
        if i == 0 {
            pref[0] = int64(skill[0])
        } else {
            pref[i] = pref[i-1] + int64(skill[i])
        }
    }

    S := int64(0)
    // use a large negative initial constant for best
    const NEG_INF int64 = -1 << 62
    for j := 1; j < m; j++ {
        prev := int64(mana[j-1])
        cur  := int64(mana[j])
        best := NEG_INF
        for i := 0; i < n; i++ {
            prev_pref := int64(0)
            if i > 0 {
                prev_pref = pref[i-1]
            }
            cand := pref[i]*prev - prev_pref*cur
            if cand > best {
                best = cand
            }
        }
        S += best
    }

    ans := S + pref[n-1]*int64(mana[m-1])
    return ans
}
