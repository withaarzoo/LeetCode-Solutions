/**
 * @param {number[]} A
 * @param {number[]} B
 * @return {number[]}
 */
var findThePrefixCommonArray = function (A, B) {
  let n = A.length;

  // Frequency array
  let freq = new Array(n + 1).fill(0);

  // Result array
  let ans = new Array(n);

  // Stores current common count
  let common = 0;

  for (let i = 0; i < n; i++) {
    // Add current value from A
    freq[A[i]]++;

    // If count becomes 2,
    // it is now common in both arrays
    if (freq[A[i]] === 2) {
      common++;
    }

    // Add current value from B
    freq[B[i]]++;

    // Same check for B
    if (freq[B[i]] === 2) {
      common++;
    }

    // Save current answer
    ans[i] = common;
  }

  return ans;
};
