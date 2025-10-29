func smallestNumber(n int) int {
    // Find smallest k with (1<<k) - 1 >= n
    k := 1
    for {
        val := (1 << k) - 1
        if val >= n {
            return val
        }
        k++
    }
}
