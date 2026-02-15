func addBinary(a string, b string) string {
    i := len(a) - 1  // pointer for a
    j := len(b) - 1  // pointer for b
    carry := 0
    
    result := []byte{}
    
    for i >= 0 || j >= 0 || carry > 0 {
        sum := carry
        
        // Add digit from a
        if i >= 0 {
            sum += int(a[i] - '0')
            i--
        }
        
        // Add digit from b
        if j >= 0 {
            sum += int(b[j] - '0')
            j--
        }
        
        // Append current bit
        result = append(result, byte(sum%2)+'0')
        
        carry = sum / 2
    }
    
    // Reverse result
    for left, right := 0, len(result)-1; left < right; left, right = left+1, right-1 {
        result[left], result[right] = result[right], result[left]
    }
    
    return string(result)
}
