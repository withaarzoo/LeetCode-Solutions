var numberOfStableArrays = function (zero, one, limit) {
  const MOD = 1e9 + 7;
  const n = zero + one;

  const fact = new Array(n + 1).fill(1);
  const invFact = new Array(n + 1).fill(1);

  const modPow = (a, b) => {
    let res = 1n;
    let base = BigInt(a);
    let exp = BigInt(b);

    while (exp > 0n) {
      if (exp & 1n) res = (res * base) % BigInt(MOD);
      base = (base * base) % BigInt(MOD);
      exp >>= 1n;
    }
    return Number(res);
  };

  for (let i = 1; i <= n; i++) fact[i] = (fact[i - 1] * i) % MOD;

  invFact[n] = modPow(fact[n], MOD - 2);

  for (let i = n - 1; i >= 0; i--)
    invFact[i] = (invFact[i + 1] * (i + 1)) % MOD;

  const C = (n, k) => {
    if (k < 0 || k > n) return 0;
    return (((fact[n] * invFact[k]) % MOD) * invFact[n - k]) % MOD;
  };

  const F = (N, K) => {
    if (K <= 0 || K > N) return 0;

    let ans = 0;
    const maxJ = Math.floor((N - K) / limit);

    for (let j = 0; j <= maxJ; j++) {
      let term = (C(K, j) * C(N - j * limit - 1, K - 1)) % MOD;

      if (j % 2) ans = (ans - term + MOD) % MOD;
      else ans = (ans + term) % MOD;
    }

    return ans;
  };

  const maxK = Math.min(zero, one + 1);

  const oneWays = new Array(maxK + 3).fill(0);

  for (let k = 1; k <= maxK + 1; k++) oneWays[k] = F(one, k);

  let ans = 0;

  for (let k = 1; k <= maxK; k++) {
    let z = F(zero, k);

    let o = (oneWays[k - 1] + 2 * oneWays[k] + oneWays[k + 1]) % MOD;

    ans = (ans + z * o) % MOD;
  }

  return ans;
};
