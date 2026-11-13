func countNonDecreasingSubarrays(nums []int, k int) int64 {
    ans := int64(0)
    cost := int64(0)
    type pair struct{ val, cnt int }
    dq := []pair{}
    n := len(nums)
    j := n - 1
    for i := n - 1; i >= 0; i-- {
        num := nums[i]
        count := 1
        for len(dq) > 0 && dq[len(dq)-1].val < num {
            back := dq[len(dq)-1]
            dq = dq[:len(dq)-1]
            count += back.cnt
            cost += int64(num-back.val) * int64(back.cnt)
        }
        dq = append(dq, pair{num, count})
        for cost > int64(k) {
            front := dq[0]
            dq = dq[1:]
            cost -= int64(front.val - nums[j])
            j--
            if front.cnt > 1 {
                dq = append([]pair{{front.val, front.cnt - 1}}, dq...)
            }
        }
        ans += int64(j - i + 1)
    }
    return ans
}