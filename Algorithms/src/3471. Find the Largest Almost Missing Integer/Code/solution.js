/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var largestInteger = function (nums, k) {
  const n = nums.length;

  // nums[i] is between 0 and 50, so a fixed frequency array is enough.
  // Using an array keeps the extra space constant.
  const freq = new Array(51).fill(0);

  // Count the frequency of every value in nums.
  for (const x of nums) {
    freq[x]++;
  }

  // When k = 1, every subarray contains exactly one element.
  // Therefore, a valid value must occur exactly once in nums.
  if (k === 1) {
    // Search from 50 downwards so the first valid value is the largest.
    for (let x = 50; x >= 0; x--) {
      if (freq[x] === 1) {
        return x;
      }
    }

    // No value appears exactly once.
    return -1;
  }

  // When k = n, the entire array is the only subarray.
  // Hence, every distinct value appears in exactly one subarray.
  if (k === n) {
    let answer = 0;

    // Find the largest value in nums.
    for (const x of nums) {
      answer = Math.max(answer, x);
    }

    return answer;
  }

  // For 1 < k < n, only the first and last elements
  // can belong to exactly one subarray of size k.
  let answer = -1;

  // The first value is valid only if it occurs once in nums.
  if (freq[nums[0]] === 1) {
    answer = Math.max(answer, nums[0]);
  }

  // The last value is valid only if it occurs once in nums.
  if (freq[nums[n - 1]] === 1) {
    answer = Math.max(answer, nums[n - 1]);
  }

  // Return the largest valid endpoint, or -1 if neither is valid.
  return answer;
};
