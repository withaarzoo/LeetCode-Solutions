func resultArray(nums []int) []int {
    // I create arr1 with the first number because the first operation is fixed.
    arr1 := []int{nums[0]}

    // I create arr2 with the second number because the second operation is fixed.
    arr2 := []int{nums[1]}

    // I process every remaining number starting from index 2.
    for i := 2; i < len(nums); i++ {
        // I compare the last element currently stored in both arrays.
        if arr1[len(arr1)-1] > arr2[len(arr2)-1] {
            // I append nums[i] to arr1 when its last value is greater.
            arr1 = append(arr1, nums[i])
        } else {
            // Otherwise, I append nums[i] to arr2.
            arr2 = append(arr2, nums[i])
        }
    }

    // I create the final slice with capacity for all elements to reduce reallocations.
    result := make([]int, 0, len(nums))

    // I append arr1 first because the required result starts with arr1.
    result = append(result, arr1...)

    // I append arr2 after arr1 to complete the concatenation.
    result = append(result, arr2...)

    // I return the final result.
    return result
}