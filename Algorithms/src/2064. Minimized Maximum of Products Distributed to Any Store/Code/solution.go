import "math"

func minimizedMaximum(n int, quantities []int) int {
    canDistribute := func(maxProducts int) bool {
        storesNeeded := 0
        for _, quantity := range quantities {
            storesNeeded += int(math.Ceil(float64(quantity) / float64(maxProducts)))
            if storesNeeded > n {
                return false
            }
        }
        return storesNeeded <= n
    }

    low, high := 1, 0
    for _, quantity := range quantities {
        if quantity > high {
            high = quantity
        }
    }
    answer := high

    for low <= high {
        mid := low + (high - low) / 2
        if canDistribute(mid) {
            answer = mid
            high = mid - 1
        } else {
            low = mid + 1
        }
    }

    return answer
}