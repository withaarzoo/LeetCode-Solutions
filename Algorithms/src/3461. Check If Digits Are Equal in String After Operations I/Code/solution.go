package main

// hasSameDigits checks if final two digits after repeated operations are equal
func hasSameDigits(s string) bool {
    // convert to slice of ints
    digits := make([]int, len(s))
    for i := 0; i < len(s); i++ {
        digits[i] = int(s[i] - '0')
    }

    // reduce until length == 2
    for len(digits) > 2 {
        next := make([]int, len(digits)-1)
        for i := 0; i+1 < len(digits); i++ {
            next[i] = (digits[i] + digits[i+1]) % 10
        }
        digits = next
    }

    return len(digits) == 2 && digits[0] == digits[1]
}
