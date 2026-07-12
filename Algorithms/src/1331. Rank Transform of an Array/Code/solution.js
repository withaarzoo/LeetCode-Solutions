/**
 * @param {number[]} arr
 * @return {number[]}
 */
var arrayRankTransform = function (arr) {
  // Create a copy of the original array
  const sorted = [...arr];

  // Sort the copied array in increasing order
  sorted.sort((a, b) => a - b);

  // Store value -> rank
  const rank = new Map();
  let currentRank = 1;

  // Assign ranks only once for every unique value
  for (const num of sorted) {
    if (!rank.has(num)) {
      rank.set(num, currentRank++);
    }
  }

  // Replace every element with its rank
  for (let i = 0; i < arr.length; i++) {
    arr[i] = rank.get(arr[i]);
  }

  // Return the transformed array
  return arr;
};
