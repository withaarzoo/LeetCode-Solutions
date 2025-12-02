/**
 * @param {number[][]} points
 * @return {number}
 */
var countTrapezoids = function (points) {
  const MOD = 1000000007n;
  const INV2 = (MOD + 1n) / 2n; // modular inverse of 2

  // 1. Count points per y
  const freq = new Map();
  for (const [x, y] of points) {
    const key = BigInt(y); // use BigInt key to be safe
    freq.set(key, (freq.get(key) || 0n) + 1n);
  }

  let sumF = 0n; // S
  let sumF2 = 0n; // SQ

  // 2. For each y, compute C(c,2)
  for (const cBig of freq.values()) {
    const c = cBig;
    if (c >= 2n) {
      const f = ((c * (c - 1n)) / 2n) % MOD; // C(c,2)
      sumF = (sumF + f) % MOD;
      sumF2 = (sumF2 + ((f * f) % MOD)) % MOD;
    }
  }

  // 3. ((S^2 - SQ) / 2) mod MOD
  let ans = (sumF * sumF) % MOD;
  ans = (ans - sumF2 + MOD) % MOD;
  ans = (ans * INV2) % MOD;

  return Number(ans); // LeetCode expects a Number
};
