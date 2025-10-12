/**
 * JavaScript implementation
 * @param {number} m
 * @param {number} k
 * @param {number[]} nums
 * @return {number}
 */
var magicalSum = function (m, k, nums) {
  const MOD = BigInt(1000000007);
  const n = nums.length;
  // combinations as BigInt
  const C = Array.from({ length: m + 1 }, () => Array(m + 1).fill(0n));
  for (let i = 0; i <= m; i++) {
    C[i][0] = 1n;
    C[i][i] = 1n;
    for (let j = 1; j < i; j++) {
      C[i][j] = (C[i - 1][j - 1] + C[i - 1][j]) % MOD;
    }
  }
  // powers powA[i][t] = nums[i]^t (BigInt)
  const powA = Array.from({ length: n }, () => Array(m + 1).fill(1n));
  for (let i = 0; i < n; i++) {
    powA[i][0] = 1n;
    const a = BigInt(nums[i]) % MOD;
    for (let t = 1; t <= m; t++) {
      powA[i][t] = (powA[i][t - 1] * a) % MOD;
    }
  }
  const M = m;
  // cur[r][carry][ones] : BigInt
  let cur = Array.from({ length: M + 1 }, () =>
    Array.from({ length: M + 1 }, () => Array(M + 1).fill(0n))
  );
  cur[M][0][0] = 1n;

  for (let i = 0; i < n; i++) {
    let nxt = Array.from({ length: M + 1 }, () =>
      Array.from({ length: M + 1 }, () => Array(M + 1).fill(0n))
    );
    for (let r = 0; r <= M; r++) {
      for (let carry = 0; carry <= M; carry++) {
        for (let ones = 0; ones <= M; ones++) {
          let val = cur[r][carry][ones];
          if (val === 0n) continue;
          for (let t = 0; t <= r; t++) {
            let newr = r - t;
            let sum = carry + t;
            let bit = sum & 1;
            let newones = ones + bit;
            if (newones > M) continue;
            let newcarry = sum >>> 1;
            let mult = (C[r][t] * powA[i][t]) % MOD;
            let add = (val * mult) % MOD;
            nxt[newr][newcarry][newones] =
              (nxt[newr][newcarry][newones] + add) % MOD;
          }
        }
      }
    }
    cur = nxt;
  }

  let ans = 0n;
  for (let carry = 0; carry <= M; carry++) {
    for (let ones = 0; ones <= M; ones++) {
      let val = cur[0][carry][ones];
      if (val === 0n) continue;
      // leftover carry bits contribute their popcount
      let extra = popcount(carry);
      if (ones + extra === k) {
        ans = (ans + val) % MOD;
      }
    }
  }
  // convert BigInt to Number safely mod 1e9+7 (fits into Number)
  return Number(ans);

  function popcount(x) {
    let c = 0;
    while (x > 0) {
      c += x & 1;
      x >>>= 1;
    }
    return c;
  }
};
