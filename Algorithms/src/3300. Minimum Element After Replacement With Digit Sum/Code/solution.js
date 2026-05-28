/**
 * @param {number[]} nums
 * @return {number}
 */
var minElement = function (nums) {
  // Helper function to calculate digit sum
  const digitSum = (num) => {
    let sum = 0;

    // Process every digit
    while (num > 0) {
      sum += num % 10; // Add last digit
      num = Math.floor(num / 10); // Remove last digit
    }

    return sum;
  };

  let ans = Infinity;

  // Check digit sum of every element
  for (const num of nums) {
    ans = Math.min(ans, digitSum(num));
  }

  return ans;
};
