/**
 * @param {number} n
 * @param {number} k
 * @return {character}
 */
var findKthBit = function (n, k) {
  // Base case
  if (n === 1) return "0";

  const length = (1 << n) - 1;
  const mid = Math.floor((length + 1) / 2);

  if (k === mid) {
    return "1";
  } else if (k < mid) {
    return findKthBit(n - 1, k);
  } else {
    const bit = findKthBit(n - 1, length - k + 1);
    return bit === "0" ? "1" : "0";
  }
};
