/**
 * @param {number[]} arr
 * @return {number[]}
 */
var sortByBits = function (arr) {
  return arr.sort((a, b) => {
    // Count number of 1 bits
    const bitsA = a.toString(2).split("1").length - 1;
    const bitsB = b.toString(2).split("1").length - 1;

    if (bitsA !== bitsB) return bitsA - bitsB;

    return a - b;
  });
};
