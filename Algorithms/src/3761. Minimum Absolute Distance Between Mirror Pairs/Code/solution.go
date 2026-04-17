func reverseNum(x int) int {
    rev := 0

    for x > 0 {
        rev = rev*10 + (x % 10)
        x /= 10
    }

    return rev
}

func minMirrorPairDistance(nums []int) int {
    lastIndex := make(map[int]int)
    ans := int(^uint(0) >> 1) // Maximum int value

    for i, num := range nums {
        // If current number exists in map,
        // then we found a mirror pair
        if prevIndex, exists := lastIndex[num]; exists {
            if i-prevIndex < ans {
                ans = i - prevIndex
            }
        }

        // Store reverse(num) with current index
        rev := reverseNum(num)
        lastIndex[rev] = i
    }

    if ans == int(^uint(0)>>1) {
        return -1
    }

    return ans
}