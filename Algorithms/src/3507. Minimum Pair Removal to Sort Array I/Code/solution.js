/**
 * @param {number[]} nums
 * @return {number}
 */
var minimumPairRemoval = function (nums) {
  let operations = 0;

  const isSorted = () => {
    for (let i = 1; i < nums.length; i++) {
      if (nums[i] < nums[i - 1]) return false;
    }
    return true;
  };

  while (!isSorted()) {
    let minSum = Infinity;
    let index = 0;

    for (let i = 0; i < nums.length - 1; i++) {
      let sum = nums[i] + nums[i + 1];
      if (sum < minSum) {
        minSum = sum;
        index = i;
      }
    }

    nums[index] = minSum;
    nums.splice(index + 1, 1);
    operations++;
  }

  return operations;
};
