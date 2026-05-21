func longestCommonPrefix(arr1 []int, arr2 []int) int {
    
    // Map used as a hash set to store prefixes
    prefixes := make(map[int]bool)

    // Generate all prefixes from arr1
    for _, num := range arr1 {

        x := num

        // Keep removing last digit
        for x > 0 {

            // Store current prefix
            prefixes[x] = true

            // Remove last digit
            x /= 10
        }
    }

    ans := 0

    // Process arr2
    for _, num := range arr2 {

        x := num

        // Keep checking prefixes
        for x > 0 {

            // Prefix found
            if prefixes[x] {

                // Count digits manually
                length := 0
                temp := x

                for temp > 0 {
                    length++
                    temp /= 10
                }

                // Update answer
                if length > ans {
                    ans = length
                }

                // Stop because larger prefix already found
                break
            }

            // Remove last digit
            x /= 10
        }
    }

    return ans
}