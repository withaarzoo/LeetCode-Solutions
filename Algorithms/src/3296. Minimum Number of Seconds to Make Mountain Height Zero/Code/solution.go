func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {

    can := func(time int64) bool {

        totalHeight := int64(0)

        for _, t := range workerTimes {

            left := int64(0)
            right := int64(mountainHeight)

            for left <= right {

                mid := (left + right) / 2

                required := int64(t) * (mid * (mid + 1) / 2)

                if required <= time {
                    left = mid + 1
                } else {
                    right = mid - 1
                }
            }

            totalHeight += right

            if totalHeight >= int64(mountainHeight) {
                return true
            }
        }

        return false
    }

    left := int64(1)
    right := int64(1e18)
    ans := right

    for left <= right {

        mid := (left + right) / 2

        if can(mid) {
            ans = mid
            right = mid - 1
        } else {
            left = mid + 1
        }
    }

    return ans
}