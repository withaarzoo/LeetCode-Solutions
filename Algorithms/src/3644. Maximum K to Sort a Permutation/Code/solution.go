func sortPermutation(nums []int) int {
    ans := -1
    for i, v := range nums {
        if v != i {
            if ans == -1 {
                ans = v
            } else {
                ans &= v
            }
        }
    }
    if ans == -1 {
        return 0
    }
    return ans
}