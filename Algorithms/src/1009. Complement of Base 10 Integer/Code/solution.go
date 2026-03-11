func bitwiseComplement(n int) int {
    
    // Edge case
    if n == 0 {
        return 1
    }

    mask := 0

    // Build mask with all bits = 1
    for mask < n {
        mask = (mask << 1) | 1
    }

    // XOR flips bits
    return mask ^ n
}