func countMajoritySubarrays(nums []int, target int) int64 {

	n := len(nums)

	// Prefix sums after transformation
	pref := make([]int, n+1)
	for i := 0; i < n; i++ {
		if nums[i] == target {
			pref[i+1] = pref[i] + 1
		} else {
			pref[i+1] = pref[i] - 1
		}
	}

	// Coordinate compression
	values := append([]int{}, pref...)
	sort.Ints(values)

	k := 0
	for _, x := range values {
		if k == 0 || values[k-1] != x {
			values[k] = x
			k++
		}
	}
	values = values[:k]

	// Fenwick Tree
	bit := make([]int, k+2)

	update := func(idx int) {
		for idx < len(bit) {
			bit[idx]++
			idx += idx & -idx
		}
	}

	query := func(idx int) int64 {
		var sum int64
		for idx > 0 {
			sum += int64(bit[idx])
			idx -= idx & -idx
		}
		return sum
	}

	var ans int64

	for _, x := range pref {

		// Compressed index
		idx := sort.SearchInts(values, x) + 1

		// Count smaller prefix sums
		ans += query(idx - 1)

		// Insert current prefix sum
		update(idx)
	}

	return ans
}