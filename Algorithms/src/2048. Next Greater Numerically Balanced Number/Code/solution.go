func isBalanced(x int) bool {
    var cnt [10]int
    t := x
    for t > 0 {
        cnt[t%10]++
        t /= 10
    }
    if cnt[0] > 0 { // 0 must not appear
        return false
    }
    for d := 1; d <= 9; d++ {
        if cnt[d] != 0 && cnt[d] != d {
            return false
        }
    }
    return true
}

func nextBeautifulNumber(n int) int {
    x := n + 1
    for {
        if isBalanced(x) {
            return x
        }
        x++
    }
}
