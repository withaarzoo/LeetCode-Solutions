func minMoves(nums []int) int {
    mn := nums[0]
    for _, x := range nums {
        if x < mn {
            mn = x
        }
    }
    moves := 0
    for _, x := range nums {
        moves += x - mn
    }
    return moves
}