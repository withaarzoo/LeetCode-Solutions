func maximumSwap(num int) int {
    // Convert the number to a string array for manipulation
    numStr := []rune(fmt.Sprint(num))
    n := len(numStr)
    
    // Track the last occurrence of each digit (0-9)
    last := make([]int, 10)
    for i := 0; i < n; i++ {
        last[numStr[i]-'0'] = i
    }
    
    // Traverse the number from left to right
    for i := 0; i < n; i++ {
        // Check if a larger digit can be found later
        for d := 9; d > int(numStr[i]-'0'); d-- {
            if last[d] > i {
                // Swap the digits and return the new number
                numStr[i], numStr[last[d]] = numStr[last[d]], numStr[i]
                result, _ := strconv.Atoi(string(numStr))
                return result
            }
        }
    }
    
    // Return the original number if no swap is done
    return num
}