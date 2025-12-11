/**
 * @param {number} n
 * @param {number[][]} buildings
 * @return {number}
 */
var countCoveredBuildings = function (n, buildings) {
  // Maps: row x -> array of y's, col y -> array of x's
  const row = new Map();
  const col = new Map();

  for (const b of buildings) {
    const x = b[0],
      y = b[1];
    if (!row.has(x)) row.set(x, []);
    row.get(x).push(y);
    if (!col.has(y)) col.set(y, []);
    col.get(y).push(x);
  }

  // Sort each group's array
  for (const [k, arr] of row) arr.sort((a, b) => a - b);
  for (const [k, arr] of col) arr.sort((a, b) => a - b);

  // Helper binary search (returns index)
  function lowerBound(arr, val) {
    let l = 0,
      r = arr.length;
    while (l < r) {
      const mid = (l + r) >> 1;
      if (arr[mid] < val) l = mid + 1;
      else r = mid;
    }
    return l;
  }

  let ans = 0;
  for (const b of buildings) {
    const x = b[0],
      y = b[1];
    const ys = row.get(x);
    const xs = col.get(y);
    const posY = lowerBound(ys, y);
    const posX = lowerBound(xs, x);
    const insideRow = posY > 0 && posY < ys.length - 1;
    const insideCol = posX > 0 && posX < xs.length - 1;
    if (insideRow && insideCol) ans++;
  }
  return ans;
};
