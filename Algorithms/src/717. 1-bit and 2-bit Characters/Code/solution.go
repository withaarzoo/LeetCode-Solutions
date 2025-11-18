package main

// isOneBitCharacter returns true if the last character must be a one-bit character.
func isOneBitCharacter(bits []int) bool {
    n := len(bits)
    i := 0
    // iterate until before the last bit
    for i < n-1 {
        if bits[i] == 1 {
            // 1 starts a two-bit character
            i += 2
        } else {
            // 0 is a one-bit character
            i += 1
        }
    }
    // if we land at the last index, it's a one-bit character
    return i == n-1
}
