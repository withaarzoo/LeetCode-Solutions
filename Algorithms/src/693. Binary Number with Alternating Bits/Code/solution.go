func hasAlternatingBits(n int) bool {

    prev := n & 1
    n >>= 1

    for n > 0 {
        curr := n & 1

        if curr == prev {
            return false
        }

        prev = curr
        n >>= 1
    }

    return true
}
