func sumAndMultiply(n int) int64 {
    // This stores the number formed by all non-zero digits.
    var x int64 = 0

    // This stores the sum of all non-zero digits.
    var sum int64 = 0

    // Find the highest place value to read digits left to right.
    divisor := 1
    for n/divisor >= 10 {
        divisor *= 10
    }

    // Process every digit from left to right.
    for divisor > 0 {
        // Extract the digit at the current place value.
        digit := n / divisor

        // Remove the current digit from n.
        n %= divisor

        // Ignore zero digits completely.
        if digit != 0 {
            // Append the current digit to x.
            x = x*10 + int64(digit)

            // Add the current digit to the sum.
            sum += int64(digit)
        }

        // Move to the next smaller place value.
        divisor /= 10
    }

    // Return the concatenated number multiplied by its digit sum.
    return x * sum
}