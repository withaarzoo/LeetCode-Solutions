func bestClosingTime(customers string) int {
    totalY := 0
    for _, c := range customers {
        if c == 'Y' {
            totalY++
        }
    }

    openPenalty := 0
    closedPenalty := totalY
    minPenalty := closedPenalty
    answer := 0

    for i, c := range customers {
        if c == 'N' {
            openPenalty++
        } else {
            closedPenalty--
        }

        currentPenalty := openPenalty + closedPenalty
        if currentPenalty < minPenalty {
            minPenalty = currentPenalty
            answer = i + 1
        }
    }

    return answer
}
