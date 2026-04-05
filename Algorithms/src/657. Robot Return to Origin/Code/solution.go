func judgeCircle(moves string) bool {
    // x represents left/right position
    // y represents up/down position
    x, y := 0, 0

    // Traverse every move
    for _, move := range moves {
        if move == 'U' {
            y++ // Move up
        } else if move == 'D' {
            y-- // Move down
        } else if move == 'R' {
            x++ // Move right
        } else if move == 'L' {
            x-- // Move left
        }
    }

    // Robot returns to origin only if both coordinates are 0
    return x == 0 && y == 0
}