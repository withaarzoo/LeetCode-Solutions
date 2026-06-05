/**
 * @param {number} num1
 * @param {number} num2
 * @return {number}
 */
var totalWaviness = function (num1, num2) {
  function solve(n) {
    if (n < 0) return 0;

    const s = String(n);
    const memo = new Map();

    function dfs(pos, started, last, secondLast, tight) {
      if (pos === s.length) {
        return [1n, 0n];
      }

      const key = `${pos}|${started}|${last}|${secondLast}`;

      if (!tight && memo.has(key)) {
        return memo.get(key);
      }

      const limit = tight ? Number(s[pos]) : 9;

      let cnt = 0n;
      let wav = 0n;

      for (let d = 0; d <= limit; d++) {
        const ntight = tight && d === limit;

        if (!started && d === 0) {
          const [c, w] = dfs(pos + 1, false, 10, 10, ntight);

          cnt += c;
          wav += w;
        } else {
          let add = 0n;

          if (started && secondLast !== 10) {
            if (
              (last > secondLast && last > d) ||
              (last < secondLast && last < d)
            ) {
              add = 1n;
            }
          }

          const nSecondLast = started ? last : 10;

          const [c, w] = dfs(pos + 1, true, d, nSecondLast, ntight);

          cnt += c;
          wav += w + add * c;
        }
      }

      const res = [cnt, wav];

      if (!tight) {
        memo.set(key, res);
      }

      return res;
    }

    return Number(dfs(0, false, 10, 10, true)[1]);
  }

  return solve(num2) - solve(num1 - 1);
};
