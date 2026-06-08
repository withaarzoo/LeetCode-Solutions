func pivotArray(nums []int, pivot int) []int {

    // Store elements smaller than pivot
    smaller := []int{}

    // Store elements equal to pivot
    equal := []int{}

    // Store elements greater than pivot
    greater := []int{}

    // Classify every element
    for _, num := range nums {
        if num < pivot {
            smaller = append(smaller, num)
        } else if num == pivot {
            equal = append(equal, num)
        } else {
            greater = append(greater, num)
        }
    }

    // Build final answer
    result := []int{}

    result = append(result, smaller...)
    result = append(result, equal...)
    result = append(result, greater...)

    return result
}