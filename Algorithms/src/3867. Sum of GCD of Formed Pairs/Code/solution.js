/**
 * @param {number[]} nums
 * @return {number}
 */
var gcdSum = function (nums) {
  // Euclidean algorithm
  const gcd = (a, b) => {
    while (b !== 0) {
      let t = a % b;
      a = b;
      b = t;
    }
    return a;
  };

  const n = nums.length;

  // Store prefix gcd values
  const prefixGcd = new Array(n);

  // Running prefix maximum
  let prefixMax = 0;

  // Build prefixGcd
  for (let i = 0; i < n; i++) {
    prefixMax = Math.max(prefixMax, nums[i]);
    prefixGcd[i] = gcd(nums[i], prefixMax);
  }

  // Sort in ascending order
  prefixGcd.sort((a, b) => a - b);

  let ans = 0;

  // Pair smallest with largest
  let left = 0;
  let right = n - 1;

  while (left < right) {
    ans += gcd(prefixGcd[left], prefixGcd[right]);
    left++;
    right--;
  }

  return ans;
};
