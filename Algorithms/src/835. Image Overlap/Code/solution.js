/**
 * @param {number[][]} img1
 * @param {number[][]} img2
 * @return {number}
 */
var largestOverlap = function (img1, img2) {
  const n = img1.length;
  // collect every coordinate that holds a 1
  const A = [],
    B = [];
  for (let i = 0; i < n; ++i) {
    for (let j = 0; j < n; ++j) {
      if (img1[i][j] === 1) A.push([i, j]);
      if (img2[i][j] === 1) B.push([i, j]);
    }
  }
  // count how many times each possible shift appears
  // shifts range from -(n-1) to +(n-1), so offset by n
  const cnt = Array.from({ length: 2 * n }, () => Array(2 * n).fill(0));
  let best = 0;
  for (const a of A) {
    for (const b of B) {
      const dx = b[0] - a[0] + n; // make index non-negative
      const dy = b[1] - a[1] + n;
      best = Math.max(best, ++cnt[dx][dy]);
    }
  }
  return best;
};
