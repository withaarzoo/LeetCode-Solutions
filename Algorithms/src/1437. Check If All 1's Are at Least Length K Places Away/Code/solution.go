func kLengthApart(nums []int, k int) bool {
    prev := -1 // index of last seen 1; -1 means none yet
    for i, v := range nums {
        if v == 1 {
            if prev != -1 {
                // zeros between current and previous 1 = i - prev - 1
                if i - prev - 1 < k {
                    return false
                }
            }
            prev = i
        }
    }
    return true
}
