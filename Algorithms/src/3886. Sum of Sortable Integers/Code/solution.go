func sortableIntegers(nums []int) int {
    n := len(nums)
    sorted := make([]int, n)
    copy(sorted, nums)
    sort.Ints(sorted)
    divs := []int{}
    for i := 1; i*i <= n; i++ {
        if n%i == 0 {
            divs = append(divs, i)
            if i != n/i {
                divs = append(divs, n/i)
            }
        }
    }
    ans := 0
    for _, k := range divs {
        ok := true
        for s := 0; s < n && ok; s += k {
            desc, pos := 0, -1
            for j := 0; j < k; j++ {
                a := nums[s+j]
                b := nums[s+(j+1)%k]
                if a > b {
                    desc++
                    pos = j
                    if desc > 1 {
                        break
                    }
                }
            }
            if desc > 1 {
                ok = false
                break
            }
            r := 0
            if desc == 1 {
                r = (pos + 1) % k
            }
            for j := 0; j < k; j++ {
                if nums[s+(r+j)%k] != sorted[s+j] {
                    ok = false
                    break
                }
            }
        }
        if ok {
            ans += k
        }
    }
    return ans
}