func maxFrequency(nums []int, k int, numOperations int) int {
    if len(nums) == 0 {
        return 0
    }
    mx := nums[0]
    for _, v := range nums {
        if v > mx {
            mx = v
        }
    }
    size := mx + k + 2
    count := make([]int, size)
    for _, v := range nums {
        count[v]++
    }
    for i := 1; i < size; i++ {
        count[i] += count[i-1]
    }

    ans := 0
    for t := 0; t < size; t++ {
        L := t - k
        if L < 0 {
            L = 0
        }
        R := t + k
        if R > size-1 {
            R = size-1
        }
        total := count[R]
        if L > 0 {
            total -= count[L-1]
        }
        freq_t := count[t]
        if t > 0 {
            freq_t -= count[t-1]
        }
        canConvert := total - freq_t
        take := numOperations
        if take > canConvert {
            take = canConvert
        }
        val := freq_t + take
        if val > ans {
            ans = val
        }
    }
    return ans
}
