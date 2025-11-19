func findFinalValue(nums []int, original int) int {
    // Build a map to emulate a set for constant-time checks
    seen := make(map[int]bool, len(nums))
    for _, v := range nums {
        seen[v] = true
    }
    // While original is present, double it
    for seen[original] {
        original *= 2
    }
    return original
}
