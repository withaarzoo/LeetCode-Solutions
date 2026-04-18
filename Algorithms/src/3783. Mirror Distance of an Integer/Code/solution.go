func mirrorDistance(n int) int {
    rev := 0
    temp := n

    // Reverse the digits of n
    for temp > 0 {
        digit := temp % 10      // Get last digit
        rev = rev*10 + digit    // Add digit to reversed number
        temp /= 10              // Remove last digit
    }

    // Return absolute difference
    if n > rev {
        return n - rev
    }
    return rev - n
}