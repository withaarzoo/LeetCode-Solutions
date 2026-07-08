func sumAndMultiply(s string, queries [][]int) []int {
    // I use int64 because intermediate multiplication can exceed 32-bit integer limits.
    const MOD int64 = 1000000007
    n := len(s)

    // nonZeroCount[i] = number of non-zero digits in s[0..i-1].
    nonZeroCount := make([]int, n+1)

    // I store only digits that remain after removing zeroes.
    digits := make([]int64, 0, n)

    // I build the original-position mapping and compressed digits together.
    for i := 0; i < n; i++ {
        // I carry the previous prefix count forward.
        nonZeroCount[i+1] = nonZeroCount[i]

        // Only non-zero digits enter the compressed sequence.
        if s[i] != '0' {
            nonZeroCount[i+1]++
            digits = append(digits, int64(s[i]-'0'))
        }
    }

    k := len(digits)

    // prefixValue[i] stores the first i compressed digits modulo MOD.
    prefixValue := make([]int64, k+1)

    // prefixSum[i] stores the sum of the first i compressed digits.
    prefixSum := make([]int64, k+1)

    // power10[i] stores 10^i modulo MOD.
    power10 := make([]int64, k+1)
    power10[0] = 1

    // I build all prefix information in one pass.
    for i := 0; i < k; i++ {
        // I append the current digit to the previous prefix number.
        prefixValue[i+1] = (prefixValue[i]*10 + digits[i]) % MOD

        // I extend the digit-sum prefix.
        prefixSum[i+1] = prefixSum[i] + digits[i]

        // I compute the next power of 10.
        power10[i+1] = (power10[i] * 10) % MOD
    }

    // I allocate exactly one result slot for each query.
    answer := make([]int, len(queries))

    // I answer every query using constant-time prefix calculations.
    for i, query := range queries {
        l := query[0]
        r := query[1]

        // left counts non-zero digits before l.
        left := nonZeroCount[l]

        // right counts non-zero digits through r.
        right := nonZeroCount[r+1]

        // length is the number of compressed digits inside the query.
        length := right - left

        // I subtract the shifted earlier prefix to isolate the query number.
        x := (prefixValue[right] - (prefixValue[left]*power10[length])%MOD + MOD) % MOD

        // I get the digit sum with prefix subtraction.
        sum := prefixSum[right] - prefixSum[left]

        // I multiply x by its digit sum and store the result.
        answer[i] = int((x * sum) % MOD)
    }

    return answer
}