func shiftingLetters(s string, shifts [][]int) string {
    n := len(s)
    diff := make([]int, n+1)

    // Build the difference array
    for _, shift := range shifts {
        start, end, direction := shift[0], shift[1], shift[2]
        delta := 1
        if direction == 0 {
            delta = -1
        }
        diff[start] += delta
        if end+1 < n {
            diff[end+1] -= delta
        }
    }

    // Calculate cumulative shifts
    shift := 0
    result := []byte(s)
    for i := 0; i < n; i++ {
        shift += diff[i]
        shift = (shift%26 + 26) % 26 // Normalize shift to [0, 25]
        result[i] = byte('a' + (result[i]-'a'+byte(shift))%26)
    }

    return string(result)
}
