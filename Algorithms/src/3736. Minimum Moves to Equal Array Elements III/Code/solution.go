func minMoves(nums []int) int {
    mx := nums[0]
    for _, x := range nums {
        if x > mx {
            mx = x
        }
    }
    ans := 0
    for _, x := range nums {
        ans += mx - x
    }
    return ans
}