func minimumAbsDifference(arr []int) [][]int {
    // Step 1: Sort the array
    sort.Ints(arr)

    minDiff := math.MaxInt32
    result := [][]int{}

    // Step 2: Find minimum difference
    for i := 1; i < len(arr); i++ {
        diff := arr[i] - arr[i-1]
        if diff < minDiff {
            minDiff = diff
        }
    }

    // Step 3: Collect valid pairs
    for i := 1; i < len(arr); i++ {
        if arr[i]-arr[i-1] == minDiff {
            result = append(result, []int{arr[i-1], arr[i]})
        }
    }

    return result
}
