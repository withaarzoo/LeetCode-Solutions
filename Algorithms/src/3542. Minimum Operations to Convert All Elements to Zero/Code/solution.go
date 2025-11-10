func minOperations(nums []int) int {
	ans := 0
	stk := make([]int, 0) // non-decreasing stack
	for _, x := range nums {
		// pop while top > x
		for len(stk) > 0 && stk[len(stk)-1] > x {
			stk = stk[:len(stk)-1]
		}
		if x == 0 {
			continue
		}
		if len(stk) == 0 || stk[len(stk)-1] < x {
			ans++
			stk = append(stk, x)
		}
	}
	return ans
}
