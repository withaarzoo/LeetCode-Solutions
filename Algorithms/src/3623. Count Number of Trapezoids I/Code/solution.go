package main

func countTrapezoids(points [][]int) int {
	const MOD int64 = 1000000007
	const INV2 int64 = (MOD + 1) / 2 // modular inverse of 2

	// 1. Count points per y
	freq := make(map[int]int)
	for _, p := range points {
		y := p[1]
		freq[y]++
	}

	var sumF int64 = 0  // S
	var sumF2 int64 = 0 // SQ

	// 2. For each y, compute C(c,2) and accumulate
	for _, c := range freq {
		if c >= 2 {
			cc := int64(c)
			f := (cc * (cc - 1) / 2) % MOD // C(c,2) mod MOD
			sumF = (sumF + f) % MOD
			sumF2 = (sumF2 + (f*f)%MOD) % MOD
		}
	}

	// 3. ((S^2 - SQ) / 2) mod MOD
	ans := (sumF * sumF) % MOD
	ans = (ans - sumF2 + MOD) % MOD
	ans = (ans * INV2) % MOD

	return int(ans)
}
