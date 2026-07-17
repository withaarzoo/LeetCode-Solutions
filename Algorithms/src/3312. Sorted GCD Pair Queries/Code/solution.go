func gcdValues(nums []int, queries []int64) []int {

	// Find maximum value.
	mx := 0
	for _, x := range nums {
		if x > mx {
			mx = x
		}
	}

	// Frequency array.
	freq := make([]int, mx+1)
	for _, x := range nums {
		freq[x]++
	}

	// exact[g] = pairs with GCD exactly g.
	exact := make([]int64, mx+1)

	// Process from large divisor to small.
	for g := mx; g >= 1; g-- {

		var cnt int64 = 0

		// Count divisible numbers.
		for m := g; m <= mx; m += g {
			cnt += int64(freq[m])
		}

		// Total pairs whose GCD is multiple of g.
		pairs := cnt * (cnt - 1) / 2

		// Remove larger exact GCDs.
		for m := g * 2; m <= mx; m += g {
			pairs -= exact[m]
		}

		exact[g] = pairs
	}

	// Prefix sums.
	prefix := make([]int64, mx+1)
	for g := 1; g <= mx; g++ {
		prefix[g] = prefix[g-1] + exact[g]
	}

	ans := make([]int, len(queries))

	for i, q := range queries {

		l, r := 1, mx

		// Binary search answer.
		for l < r {
			mid := (l + r) / 2

			if prefix[mid] >= q+1 {
				r = mid
			} else {
				l = mid + 1
			}
		}

		ans[i] = l
	}

	return ans
}