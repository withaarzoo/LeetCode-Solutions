import "sort"

// canBeEqual checks if the target array can be made equal to the arr array
// by sorting both arrays and comparing their elements.
func canBeEqual(target []int, arr []int) bool {
    // Step 1: Sort the target array in ascending order.
    sort.Ints(target)
    // Step 2: Sort the arr array in ascending order.
    sort.Ints(arr)
    
    // Step 3: Compare each element of the sorted target array with the sorted arr array.
    for i := range target {
        // Step 4: If any corresponding elements are not equal, return false.
        if target[i] != arr[i] {
            return false
        }
    }
    // Step 5: If all elements match, return true.
    return true
}
