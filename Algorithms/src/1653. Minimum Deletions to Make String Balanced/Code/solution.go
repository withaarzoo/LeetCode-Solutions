func minimumDeletions(s string) int {
    countB := 0
    deletions := 0

    for _, ch := range s {
        if ch == 'b' {
            countB++
        } else {
            if deletions+1 < countB {
                deletions = deletions + 1
            } else {
                deletions = countB
            }
        }
    }
    return deletions
}
