func minSwaps(nums []int) int {
    // Step 1: Calculate the total number of 1s in the array
    totalOnes := 0
    for _, num := range nums {
        if num == 1 {
            totalOnes++
        }
    }
    
    // Step 2: If there are no 1s, no swaps are needed
    if totalOnes == 0 {
        return 0
    }

    // Initialize the length of the array
    n := len(nums)
    
    // Step 3: Initialize the maximum number of 1s in any window and the current number of 1s in the current window
    maxOnesInWindow := 0
    currentOnesInWindow := 0

    // Step 4: Calculate the number of 1s in the initial window of size totalOnes
    for i := 0; i < totalOnes; i++ {
        currentOnesInWindow += nums[i]
    }

    // Set the initial window's number of 1s as the maximum found so far
    maxOnesInWindow = currentOnesInWindow

    // Step 5: Slide the window across the array to find the window with the maximum number of 1s
    for i := 1; i < n; i++ {
        // Subtract the element that is sliding out of the window
        currentOnesInWindow -= nums[i - 1]
        // Add the new element that is sliding into the window, using modulo to wrap around
        currentOnesInWindow += nums[(i + totalOnes - 1) % n]
        // Update the maximum number of 1s found in any window
        if currentOnesInWindow > maxOnesInWindow {
            maxOnesInWindow = currentOnesInWindow
        }
    }

    // Step 6: The minimum number of swaps needed is the difference between total 1s and the maximum 1s in any window
    return totalOnes - maxOnesInWindow
}
