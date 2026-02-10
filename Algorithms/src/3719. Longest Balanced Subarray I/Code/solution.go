func longestBalanced(nums []int) int {
    n := len(nums)
    ans := 0

    for i := 0; i < n; i++ {
        evenSet := make(map[int]bool)
        oddSet := make(map[int]bool)

        for j := i; j < n; j++ {
            if nums[j]%2 == 0 {
                evenSet[nums[j]] = true
            } else {
                oddSet[nums[j]] = true
            }

            if len(evenSet) == len(oddSet) {
                if j-i+1 > ans {
                    ans = j - i + 1
                }
            }
        }
    }
    return ans
}
