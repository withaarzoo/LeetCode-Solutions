/**
 * @param {number[]} arr
 * @param {number} k
 * @return {boolean}
 */
var canArrange = function (arr, k) {
  // Frequency array to store the count of remainders
  let remainderFreq = new Array(k).fill(0);

  // Step 1: Calculate the remainder for each element and store the frequency
  for (let num of arr) {
    let remainder = ((num % k) + k) % k; // Ensure non-negative remainder
    remainderFreq[remainder]++;
  }

  // Step 2: Check if the pairing condition holds
  for (let i = 0; i <= Math.floor(k / 2); i++) {
    if (i === 0) {
      // Elements with remainder 0 must pair among themselves
      if (remainderFreq[i] % 2 !== 0) return false;
    } else {
      // Remainder i must pair with remainder k-i
      if (remainderFreq[i] !== remainderFreq[k - i]) return false;
    }
  }

  return true;
};
