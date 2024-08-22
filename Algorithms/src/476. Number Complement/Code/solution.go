func findComplement(num int) int {
    // Initialize mask to 0. The mask will be used to cover all the bits of 'num'.
    mask := 0
    
    // Create a temporary variable 'temp' to store the value of 'num'.
    // This will be used to determine the length of 'num' in bits.
    temp := num
    
    // Loop to create a mask where all bits are set to 1, covering the length of 'num'.
    // The idea is to match the bit-length of 'num' with all bits as 1s.
    for temp != 0 {
        // Shift the current mask to the left by 1 bit and set the rightmost bit to 1.
        // For example, if the mask is 0000 and temp has 3 bits, the mask will become 111.
        mask = (mask << 1) | 1
        
        // Right shift 'temp' by 1 bit to eventually bring 'temp' down to 0.
        // This helps to determine when we have matched the length of 'num'.
        temp >>= 1
    }
    
    // Use XOR between 'num' and 'mask' to flip all the bits of 'num'.
    // XOR with 1 inverts the bits, so we get the complement of the number.
    return num ^ mask
}
