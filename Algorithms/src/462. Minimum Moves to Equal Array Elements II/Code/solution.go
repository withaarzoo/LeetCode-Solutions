func minMoves2(nums []int) int {
    sort.Ints(nums)
    median := nums[len(nums)/2]
    moves := 0
    for _, num := range nums {
        diff := num - median
        if diff < 0 {
            diff = -diff
        }
        moves += diff
    }
    return moves
}