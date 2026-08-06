func smallestNumber(n int, t int) int {
    // Keep checking numbers starting from n
    for {
        product := 1
        x := n

        // Calculate the product of all digits
        for x > 0 {
            product *= x % 10
            x /= 10
        }

        // Return the first valid number
        if product%t == 0 {
            return n
        }

        // Check the next number
        n++
    }
}