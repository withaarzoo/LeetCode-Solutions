func zigZagArrays(n int, l int, r int) int {
	const MOD int64 = 1000000007

	m := r - l + 1
	sz := 2 * m

	// Matrix multiplication
	multiply := func(A, B [][]int64) [][]int64 {
		C := make([][]int64, sz)

		for i := 0; i < sz; i++ {
			C[i] = make([]int64, sz)

			for k := 0; k < sz; k++ {
				if A[i][k] == 0 {
					continue
				}

				cur := A[i][k]

				for j := 0; j < sz; j++ {
					if B[k][j] == 0 {
						continue
					}

					C[i][j] = (C[i][j] + cur*B[k][j]) % MOD
				}
			}
		}

		return C
	}

	T := make([][]int64, sz)

	for i := 0; i < sz; i++ {
		T[i] = make([]int64, sz)
	}

	for x := 0; x < m; x++ {

		// up[x] -> down[y]
		for y := x + 1; y < m; y++ {
			T[x][m+y] = 1
		}

		// down[x] -> up[y]
		for y := 0; y < x; y++ {
			T[m+x][y] = 1
		}
	}

	result := make([][]int64, sz)

	for i := 0; i < sz; i++ {
		result[i] = make([]int64, sz)
		result[i][i] = 1
	}

	power := n - 1

	for power > 0 {
		if power&1 == 1 {
			result = multiply(result, T)
		}

		T = multiply(T, T)
		power >>= 1
	}

	var answer int64 = 0

	for i := 0; i < sz; i++ {
		var rowSum int64 = 0

		for j := 0; j < sz; j++ {
			rowSum = (rowSum + result[i][j]) % MOD
		}

		answer = (answer + rowSum) % MOD
	}

	return int(answer)
}