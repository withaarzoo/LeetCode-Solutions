/**
 * @param {string} word
 * @return {number}
 */
var minimumDistance = function (word) {
  const memo = new Map();

  // Calculate Manhattan distance between two letters
  function getDist(a, b) {
    // 26 means finger is not placed yet
    if (a === 26 || b === 26) return 0;

    const row1 = Math.floor(a / 6);
    const col1 = a % 6;

    const row2 = Math.floor(b / 6);
    const col2 = b % 6;

    return Math.abs(row1 - row2) + Math.abs(col1 - col2);
  }

  function solve(idx, f1, f2) {
    // If all characters are typed
    if (idx === word.length) return 0;

    const key = `${idx},${f1},${f2}`;

    // Return memoized result
    if (memo.has(key)) return memo.get(key);

    const cur = word.charCodeAt(idx) - 65;

    // Option 1: Use finger 1
    const useFinger1 = getDist(f1, cur) + solve(idx + 1, cur, f2);

    // Option 2: Use finger 2
    const useFinger2 = getDist(f2, cur) + solve(idx + 1, f1, cur);

    const ans = Math.min(useFinger1, useFinger2);

    memo.set(key, ans);
    return ans;
  }

  // Both fingers initially not placed
  return solve(0, 26, 26);
};
