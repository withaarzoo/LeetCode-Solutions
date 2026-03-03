func findKthBit(n int, k int) byte {
    // Base case
    if n == 1 {
        return '0'
    }

    length := (1 << n) - 1
    mid := (length + 1) / 2

    if k == mid {
        return '1'
    } else if k < mid {
        return findKthBit(n-1, k)
    } else {
        bit := findKthBit(n-1, length-k+1)
        if bit == '0' {
            return '1'
        }
        return '0'
    }
}