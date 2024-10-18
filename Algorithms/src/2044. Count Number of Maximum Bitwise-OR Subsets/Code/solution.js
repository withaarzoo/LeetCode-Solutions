/**
 * @param {number[]} nums
 * @return {number}
 */
var countMaxOrSubsets = function (nums) {
  let maxOR = 0;

  // Step 1: Compute the maximum OR
  for (let num of nums) {
    maxOR |= num;
  }

  let count = 0;

  const backtrack = (index, currentOR) => {
    if (currentOR === maxOR) {
      count++;
    }

    for (let i = index; i < nums.length; i++) {
      backtrack(i + 1, currentOR | nums[i]);
    }
  };

  // Step 2: Backtrack to count the subsets
  backtrack(0, 0);

  return count;
};
