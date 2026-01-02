func repeatedNTimes(nums []int) int {
    seen := make(map[int]bool)

    for _, x := range nums {
        // If already seen, this is the repeated element
        if seen[x] {
            return x
        }
        // Otherwise, mark as seen
        seen[x] = true
    }
    return -1 // Guaranteed answer exists
}
