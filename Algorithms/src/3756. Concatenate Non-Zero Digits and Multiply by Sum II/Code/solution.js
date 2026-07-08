/**
 * @param {string} s
 * @param {number[][]} queries
 * @return {number[]}
 */
var sumAndMultiply = function (s, queries) {
  // I use BigInt because JavaScript Number cannot safely multiply values near MOD.
  const MOD = 1000000007n;
  const n = s.length;

  // nonZeroCount[i] = number of non-zero digits in s[0..i-1].
  const nonZeroCount = new Array(n + 1).fill(0);

  // I store only digits that can actually appear in a query result.
  const digits = [];

  // I build the original-position mapping and compressed digits together.
  for (let i = 0; i < n; i++) {
    // I start with the count from the previous prefix.
    nonZeroCount[i + 1] = nonZeroCount[i];

    // A non-zero digit increases the count and enters the compressed list.
    if (s[i] !== "0") {
      nonZeroCount[i + 1]++;
      digits.push(Number(s[i]));
    }
  }

  const k = digits.length;

  // prefixValue[i] stores the first i compressed digits modulo MOD.
  const prefixValue = new Array(k + 1).fill(0n);

  // prefixSum[i] stores the sum of the first i compressed digits.
  const prefixSum = new Array(k + 1).fill(0);

  // power10[i] stores 10^i modulo MOD.
  const power10 = new Array(k + 1).fill(1n);

  // I build all prefix information over the compressed sequence.
  for (let i = 0; i < k; i++) {
    // I append the current digit to the previous prefix number.
    prefixValue[i + 1] = (prefixValue[i] * 10n + BigInt(digits[i])) % MOD;

    // I extend the normal digit-sum prefix.
    prefixSum[i + 1] = prefixSum[i] + digits[i];

    // I compute the next power of 10.
    power10[i + 1] = (power10[i] * 10n) % MOD;
  }

  // I create the result array by processing every query independently.
  return queries.map(([l, r]) => {
    // left counts non-zero digits before l.
    const left = nonZeroCount[l];

    // right counts non-zero digits through r.
    const right = nonZeroCount[r + 1];

    // len is the compressed range length.
    const len = right - left;

    // I remove the earlier prefix after shifting it by len decimal places.
    const x =
      (prefixValue[right] - ((prefixValue[left] * power10[len]) % MOD) + MOD) %
      MOD;

    // I get the digit sum with prefix subtraction.
    const sum = prefixSum[right] - prefixSum[left];

    // The final value is below MOD, so converting it back to Number is safe.
    return Number((x * BigInt(sum)) % MOD);
  });
};
