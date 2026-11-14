func minimumSumSubarray(nums []int, l int, r int) int {
    n := len(nums)
    ans := math.MaxInt32
    for i := 0; i < n; i++ {
        sum := 0
        for j := i; j < n; j++ {
            sum += nums[j]
            length := j - i + 1
            if length >= l && length <= r && sum > 0 {
                if sum < ans {
                    ans = sum
                }
            }
        }
    }
    if ans == math.MaxInt32 {
        return -1
    }
    return ans
}