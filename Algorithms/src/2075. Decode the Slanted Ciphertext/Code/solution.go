func decodeCiphertext(encodedText string, rows int) string {
    // Edge case: empty string
    if len(encodedText) == 0 {
        return ""
    }

    n := len(encodedText)
    cols := n / rows

    result := make([]byte, 0)

    // Start from every column in the first row
    for startCol := 0; startCol < cols; startCol++ {
        row, col := 0, startCol

        // Move diagonally down-right
        for row < rows && col < cols {
            index := row*cols + col
            result = append(result, encodedText[index])
            row++
            col++
        }
    }

    // Remove trailing spaces
    end := len(result) - 1
    for end >= 0 && result[end] == ' ' {
        end--
    }

    return string(result[:end+1])
}