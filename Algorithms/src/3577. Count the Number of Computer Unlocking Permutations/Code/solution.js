/**
 * @param {number[]} complexity
 * @return {number}
 */
var countPermutations = function (complexity) {
  const MOD = 1_000_000_007n;
  const n = complexity.length;

  // 1. Find global minimum and its frequency
  let minVal = complexity[0];
  let cntMin = 0;
  for (const x of complexity) {
    if (x < minVal) {
      minVal = x;
      cntMin = 1;
    } else if (x === minVal) {
      cntMin++;
    }
  }

  // 2. Check if index 0 has unique minimum
  if (complexity[0] !== minVal || cntMin !== 1) {
    return 0;
  }

  // 3. Compute (n - 1)! % MOD using BigInt
  let ans = 1n;
  for (let i = 2n; i <= BigInt(n - 1); i++) {
    ans = (ans * i) % MOD;
  }
  return Number(ans);
};
