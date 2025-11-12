package main

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	if a < 0 {
		return -a
	}
	return a
}

func minOperations(nums []int) int {
	n := len(nums)

	// 1) If there are ones, just convert others.
	ones := 0
	for _, x := range nums {
		if x == 1 {
			ones++
		}
	}
	if ones > 0 {
		return n - ones
	}

	// 2) If global gcd > 1, impossible.
	g := 0
	for _, x := range nums {
		g = gcd(g, x)
	}
	if g > 1 {
		return -1
	}

	// 3) Shortest subarray with gcd == 1.
	const INF = int(1e9)
	best := INF
	for i := 0; i < n; i++ {
		cur := 0
		for j := i; j < n; j++ {
			cur = gcd(cur, nums[j])
			if cur == 1 {
				if length := j - i + 1; length < best {
					best = length
				}
				break
			}
		}
	}

	// 4) Create first 1 + spread to all.
	return (best - 1) + (n - 1)
}
