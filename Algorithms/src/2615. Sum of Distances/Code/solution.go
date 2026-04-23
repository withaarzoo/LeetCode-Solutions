func distance(nums []int) []int64 {
    n := len(nums)
    mp := make(map[int][]int)

    // Group indices
    for i := 0; i < n; i++ {
        mp[nums[i]] = append(mp[nums[i]], i)
    }

    res := make([]int64, n)

    for _, idx := range mp {
        k := len(idx)

        var prefixSum int64 = 0
        var totalSum int64 = 0

        for _, x := range idx {
            totalSum += int64(x)
        }

        for i := 0; i < k; i++ {
            curr := int64(idx[i])

            left := curr*int64(i) - prefixSum
            right := (totalSum - prefixSum - curr) - curr*int64(k-i-1)

            res[curr] = left + right

            prefixSum += curr
        }
    }

    return res
}