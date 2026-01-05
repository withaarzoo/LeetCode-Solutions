func maxMatrixSum(matrix [][]int) int64 {
    var sum int64 = 0
    negativeCount := 0
    minAbs := int(1e9)

    for _, row := range matrix {
        for _, val := range row {
            if val < 0 {
                negativeCount++
                val = -val
            }
            sum += int64(val)
            if val < minAbs {
                minAbs = val
            }
        }
    }

    if negativeCount%2 == 1 {
        sum -= int64(2 * minAbs)
    }

    return sum
}
