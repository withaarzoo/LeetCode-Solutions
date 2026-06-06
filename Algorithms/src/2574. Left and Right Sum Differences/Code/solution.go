func leftRightDifference(nums []int) []int {
    
    n := len(nums)

    // Calculate total array sum
    rightSum := 0
    for _, num := range nums {
        rightSum += num
    }

    // Sum of elements on the left side
    leftSum := 0

    // Result array
    ans := make([]int, n)

    for i := 0; i < n; i++ {

        // Remove current element from right side sum
        rightSum -= nums[i]

        // Store absolute difference
        diff := leftSum - rightSum
        if diff < 0 {
            diff = -diff
        }
        ans[i] = diff

        // Add current element to left side sum
        leftSum += nums[i]
    }

    return ans
}