func shortestSubarray(nums []int, k int) int {
    n := len(nums)
    prefix := make([]int64, n+1)

    // Step 1: Compute prefix sums
    for i := 0; i < n; i++ {
        prefix[i+1] = prefix[i] + int64(nums[i])
    }

    dq := []int{}
    minLength := n + 1

    // Step 2: Process prefix sums
    for i := 0; i <= n; i++ {
        for len(dq) > 0 && prefix[i]-prefix[dq[0]] >= int64(k) {
            if i-dq[0] < minLength {
                minLength = i - dq[0]
            }
            dq = dq[1:]
        }

        for len(dq) > 0 && prefix[i] <= prefix[dq[len(dq)-1]] {
            dq = dq[:len(dq)-1]
        }

        dq = append(dq, i)
    }

    if minLength == n+1 {
        return -1
    }
    return minLength
}
