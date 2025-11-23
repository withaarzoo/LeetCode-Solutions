func maxSumDivThree(nums []int) int {
    sum := 0
    const INF = int(1e9)
    r1Min1, r1Min2 := INF, INF // two smallest remainder-1 numbers
    r2Min1, r2Min2 := INF, INF // two smallest remainder-2 numbers

    for _, x := range nums {
        sum += x
        r := x % 3
        if r == 1 {
            if x < r1Min1 {
                r1Min2 = r1Min1
                r1Min1 = x
            } else if x < r1Min2 {
                r1Min2 = x
            }
        } else if r == 2 {
            if x < r2Min1 {
                r2Min2 = r2Min1
                r2Min1 = x
            } else if x < r2Min2 {
                r2Min2 = x
            }
        }
    }

    mod := sum % 3
    if mod == 0 {
        return sum
    }

    removeCost := int64(1e18)

    if mod == 1 {
        if r1Min1 != INF {
            if int64(r1Min1) < removeCost {
                removeCost = int64(r1Min1)
            }
        }
        if r2Min2 != INF {
            if int64(r2Min1+r2Min2) < removeCost {
                removeCost = int64(r2Min1 + r2Min2)
            }
        }
    } else { // mod == 2
        if r2Min1 != INF {
            if int64(r2Min1) < removeCost {
                removeCost = int64(r2Min1)
            }
        }
        if r1Min2 != INF {
            if int64(r1Min1+r1Min2) < removeCost {
                removeCost = int64(r1Min1 + r1Min2)
            }
        }
    }

    if removeCost >= int64(1e18) {
        return 0
    }
    return sum - int(removeCost)
}
