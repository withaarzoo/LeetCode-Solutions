func maxLength(nums []int) int {
    n := len(nums)
    ans := 0
    for i := 0; i < n; i++ {
        prod := int64(1)
        g := 0
        lc := int64(1)
        for j := i; j < n; j++ {
            x := nums[j]
            if g == 0 {
                g = x
            } else {
                a, b := g, x
                for b != 0 {
                    a, b = b, a%b
                }
                g = a
            }
            gg := lc
            b := int64(x)
            for b != 0 {
                gg, b = b, gg%b
            }
            if lc/gg > (1<<62)/int64(x) {
                break
            }
            lc = lc / gg * int64(x)
            if prod > (1<<62)/int64(x) {
                break
            }
            prod *= int64(x)
            if prod == lc*int64(g) {
                if j-i+1 > ans {
                    ans = j - i + 1
                }
            }
        }
    }
    return ans
}