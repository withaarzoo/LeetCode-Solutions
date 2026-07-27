func maxProduct(nums []int) int {

    // Store the largest value found so far
    first := 0

    // Store the second largest value found so far
    second := 0

    // Traverse the array once
    for _, num := range nums {

        // If current number becomes the largest
        if num >= first {

            // Old largest becomes second largest
            second = first

            // Update largest
            first = num

        } else if num > second {
            // Otherwise update second largest if needed
            second = num
        }
    }

    // Return the required product
    return (first - 1) * (second - 1)
}