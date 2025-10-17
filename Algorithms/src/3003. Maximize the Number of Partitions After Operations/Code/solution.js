/**
 * @param {string} s
 * @param {number} k
 * @return {number}
 */
var maxPartitionsAfterOperations = function (s, k) {
  const memo = new Map();

  const dp = (i, mask, canChange) => {
    if (i === s.length) return 0;
    const key = `${i},${mask},${canChange}`;
    if (memo.has(key)) return memo.get(key);

    const bit = s.charCodeAt(i) - 97;
    let newMask = mask | (1 << bit);
    let res = 0;

    if (countBits(newMask) > k) res = 1 + dp(i + 1, 1 << bit, canChange);
    else res = dp(i + 1, newMask, canChange);

    if (canChange) {
      for (let j = 0; j < 26; j++) {
        let changed = mask | (1 << j);
        if (countBits(changed) > k)
          res = Math.max(res, 1 + dp(i + 1, 1 << j, false));
        else res = Math.max(res, dp(i + 1, changed, false));
      }
    }

    memo.set(key, res);
    return res;
  };

  const countBits = (x) => x.toString(2).split("1").length - 1;

  return dp(0, 0, true) + 1;
};
