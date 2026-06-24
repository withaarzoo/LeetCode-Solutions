/**
 * @param {number} n
 * @param {number} l
 * @param {number} r
 * @return {number}
 */
var zigZagArrays = function (n, l, r) {
  const MOD = 1000000007n;

  const m = r - l + 1;
  const sz = 2 * m;

  // Matrix multiplication
  const multiply = (A, B) => {
    const C = Array.from({ length: sz }, () => Array(sz).fill(0n));

    for (let i = 0; i < sz; i++) {
      for (let k = 0; k < sz; k++) {
        if (A[i][k] === 0n) continue;

        const cur = A[i][k];

        for (let j = 0; j < sz; j++) {
          if (B[k][j] === 0n) continue;

          C[i][j] = (C[i][j] + cur * B[k][j]) % MOD;
        }
      }
    }

    return C;
  };

  let T = Array.from({ length: sz }, () => Array(sz).fill(0n));

  for (let x = 0; x < m; x++) {
    // up[x] -> down[y]
    for (let y = x + 1; y < m; y++) {
      T[x][m + y] = 1n;
    }

    // down[x] -> up[y]
    for (let y = 0; y < x; y++) {
      T[m + x][y] = 1n;
    }
  }

  let result = Array.from({ length: sz }, (_, i) =>
    Array.from({ length: sz }, (_, j) => (i === j ? 1n : 0n)),
  );

  let power = BigInt(n - 1);

  while (power > 0n) {
    if (power & 1n) {
      result = multiply(result, T);
    }

    T = multiply(T, T);
    power >>= 1n;
  }

  let answer = 0n;

  for (let i = 0; i < sz; i++) {
    let rowSum = 0n;

    for (let j = 0; j < sz; j++) {
      rowSum = (rowSum + result[i][j]) % MOD;
    }

    answer = (answer + rowSum) % MOD;
  }

  return Number(answer);
};
