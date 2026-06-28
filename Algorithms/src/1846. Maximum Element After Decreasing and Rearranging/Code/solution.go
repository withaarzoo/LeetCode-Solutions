func maximumElementAfterDecrementingAndRearranging(arr []int) int {

    // Sort the array in ascending order.
    sort.Ints(arr)

    // The first element must become 1.
    arr[0] = 1

    // Build the largest valid sequence.
    for i := 1; i < len(arr); i++ {

        // The current value cannot exceed previous + 1.
        if arr[i] > arr[i-1]+1 {
            arr[i] = arr[i-1] + 1
        }
    }

    // The last element is the answer.
    return arr[len(arr)-1]
}