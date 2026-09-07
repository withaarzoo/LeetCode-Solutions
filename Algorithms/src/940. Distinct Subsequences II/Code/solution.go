func distinctSubseqII(s string) int {
	const MOD int64 = 1000000007 // I use this modulo because the answer can be very large.

	var dp int64 = 1 // I start with the empty subsequence as the only subsequence.
	last := make([]int64, 26) // last[c] stores dp from before the previous occurrence of c.

	for i := 0; i < len(s); i++ { // I process every character exactly once.
		index := int(s[i] - 'a') // I convert the lowercase character into an index from 0 to 25.

		oldDp := dp // I save the old count before changing dp.

		dp = (2*dp - last[index] + MOD) % MOD // I double the count and remove duplicate subsequences.

		last[index] = oldDp // I remember the old count for future occurrences of this character.
	}

	return int((dp - 1 + MOD) % MOD) // I remove the empty subsequence from the final answer.
}