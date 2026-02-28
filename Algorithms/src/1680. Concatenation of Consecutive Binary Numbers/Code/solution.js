/**
 * @param {number} n
 * @return {number}
 */
var concatenatedBinary = function (n) {
  const MOD = 1000000007n;
  let ans = 0n;
  let bitLength = 0n;

  for (let i = 1n; i <= BigInt(n); i++) {
    // If i is power of 2
    if ((i & (i - 1n)) === 0n) {
      bitLength++;
    }

    // Shift and add
    ans = (((ans << bitLength) % MOD) + i) % MOD;
  }

  return Number(ans);
};
