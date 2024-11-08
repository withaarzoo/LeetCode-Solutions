/**
 * @param {number[]} nums
 * @param {number} maximumBit
 * @return {number[]}
 */
var getMaximumXor = function (nums, maximumBit) {
  let n = nums.length;
  let answer = new Array(n);
  let XORed = 0;

  // Calculate the cumulative XOR of the entire nums array
  for (let num of nums) {
    XORed ^= num;
  }

  // max_k is 2^maximumBit - 1
  let max_k = (1 << maximumBit) - 1;

  // Process each query in reverse
  for (let i = 0; i < n; i++) {
    // Calculate the k that maximizes XOR
    answer[i] = XORed ^ max_k;

    // Update XORed by removing the effect of the last element
    XORed ^= nums[n - 1 - i];
  }

  return answer;
};
