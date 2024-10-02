/**
 * @param {number[]} arr
 * @return {number[]}
 */
var arrayRankTransform = function (arr) {
  if (arr.length === 0) return [];

  // Step 1: Create a sorted copy of the array
  let sortedArr = [...arr];
  sortedArr.sort((a, b) => a - b);

  // Step 2: Create a map to store the rank of each element
  let rankMap = new Map();
  let rank = 1;

  // Step 3: Assign ranks to sorted elements
  for (let num of sortedArr) {
    if (!rankMap.has(num)) {
      rankMap.set(num, rank++);
    }
  }

  // Step 4: Replace each element in the original array with its rank
  for (let i = 0; i < arr.length; i++) {
    arr[i] = rankMap.get(arr[i]);
  }

  return arr;
};
