/**
 * @param {string} s
 * @return {number}
 */
var distinctSubseqII = function (s) {
  const MOD = 1000000007; // I use this modulo because the answer can be very large.

  let dp = 1; // I start with the empty subsequence.
  const last = new Array(26).fill(0); // last[c] stores dp from before the previous occurrence of c.

  for (const ch of s) {
    // I process every character exactly once.
    const index = ch.charCodeAt(0) - 97; // I convert the lowercase character into an index from 0 to 25.

    const oldDp = dp; // I save the old count because it is needed for last[index].

    dp = (2 * dp - last[index] + MOD) % MOD; // I double the count and remove duplicate subsequences.

    last[index] = oldDp; // I remember the old count for future occurrences of this character.
  }

  return (dp - 1 + MOD) % MOD; // I remove the empty subsequence from the answer.
};
