/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var shortestSubarray = function (nums, k) {
  const n = nums.length;
  const prefix = new Array(n + 1).fill(0);

  // Step 1: Compute prefix sums
  for (let i = 0; i < n; i++) {
    prefix[i + 1] = prefix[i] + nums[i];
  }

  const deque = [];
  let minLength = Infinity;

  // Step 2: Process prefix sums
  for (let i = 0; i <= n; i++) {
    while (deque.length > 0 && prefix[i] - prefix[deque[0]] >= k) {
      minLength = Math.min(minLength, i - deque.shift());
    }

    while (deque.length > 0 && prefix[i] <= prefix[deque[deque.length - 1]]) {
      deque.pop();
    }

    deque.push(i);
  }

  return minLength === Infinity ? -1 : minLength;
};
