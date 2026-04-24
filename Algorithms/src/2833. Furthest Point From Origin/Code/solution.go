func furthestDistanceFromOrigin(moves string) int {
    left, right, blank := 0, 0, 0

    for _, c := range moves {
        if c == 'L' {
            left++
        } else if c == 'R' {
            right++
        } else {
            blank++
        }
    }

    position := right - left

    if position < 0 {
        position = -position
    }

    return position + blank
}