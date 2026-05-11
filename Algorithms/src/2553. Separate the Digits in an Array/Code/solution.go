func separateDigits(nums []int) []int {
    
    // Final array to store separated digits
    result := []int{}

    // Traverse every number
    for _, num := range nums {

        // Convert number into string
        str := strconv.Itoa(num)

        // Traverse every character in string
        for _, ch := range str {

            // Convert character digit into integer
            result = append(result, int(ch-'0'))
        }
    }

    // Return final answer
    return result
}