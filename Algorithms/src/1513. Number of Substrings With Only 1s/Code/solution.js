/**
 * @param {string} s
 * @return {number}
 */
var numSub = function (s) {
  const MOD = 1000000007n; // use BigInt for safety
  let res = 0n;
  let cnt = 0n;

  for (let i = 0; i < s.length; ++i) {
    if (s[i] === "1") {
      cnt += 1n;
    } else {
      res = (res + (((cnt * (cnt + 1n)) / 2n) % MOD)) % MOD;
      cnt = 0n;
    }
  }
  res = (res + (((cnt * (cnt + 1n)) / 2n) % MOD)) % MOD;
  return Number(res); // LeetCode expects a number
};
