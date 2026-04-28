var minOperations = function (grid, x) {
  let nums = [];

  // Flatten grid
  for (let row of grid) {
    for (let val of row) {
      nums.push(val);
    }
  }

  // Check feasibility
  let rem = nums[0] % x;
  for (let num of nums) {
    if (num % x !== rem) return -1;
  }

  // Sort
  nums.sort((a, b) => a - b);

  // Median
  let median = nums[Math.floor(nums.length / 2)];

  // Count operations
  let ops = 0;
  for (let num of nums) {
    ops += Math.abs(num - median) / x;
  }

  return ops;
};
