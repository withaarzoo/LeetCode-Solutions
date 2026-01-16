/**
 * @param {number} m
 * @param {number} n
 * @param {number[]} hFences
 * @param {number[]} vFences
 * @return {number}
 */
var maximizeSquareArea = function (m, n, hFences, vFences) {
  const MOD = 1000000007n; // BigInt

  // Add boundary fences
  hFences.push(1, m);
  vFences.push(1, n);

  // Sort fences
  hFences.sort((a, b) => a - b);
  vFences.sort((a, b) => a - b);

  const horizontal = new Set();
  const vertical = new Set();

  // All horizontal distances
  for (let i = 0; i < hFences.length; i++) {
    for (let j = i + 1; j < hFences.length; j++) {
      horizontal.add(hFences[j] - hFences[i]);
    }
  }

  // All vertical distances
  for (let i = 0; i < vFences.length; i++) {
    for (let j = i + 1; j < vFences.length; j++) {
      vertical.add(vFences[j] - vFences[i]);
    }
  }

  let maxSide = 0;

  // Find maximum common distance
  for (let d of horizontal) {
    if (vertical.has(d)) {
      maxSide = Math.max(maxSide, d);
    }
  }

  if (maxSide === 0) return -1;

  // IMPORTANT: use BigInt for square
  const side = BigInt(maxSide);
  const area = (side * side) % MOD;

  return Number(area);
};
