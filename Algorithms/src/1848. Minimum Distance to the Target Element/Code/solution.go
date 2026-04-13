func getMinDistance(nums []int, target int, start int) int {
    // Store the minimum distance found so far
    answer := int(^uint(0) >> 1) // Maximum integer value

    // Traverse through the array
    for i := 0; i < len(nums); i++ {
        // Check if current element is the target
        if nums[i] == target {
            distance := i - start
            if distance < 0 {
                distance = -distance
            }

            if distance < answer {
                answer = distance
            }
        }
    }

    return answer
}