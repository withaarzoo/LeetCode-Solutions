/**
 * @param {number[]} nums
 * @return {number}
 */
var minMirrorPairDistance = function (nums) {
  // Function to reverse digits of a number
  const reverseNum = (x) => {
    let rev = 0;

    while (x > 0) {
      rev = rev * 10 + (x % 10);
      x = Math.floor(x / 10);
    }

    return rev;
  };

  const lastIndex = new Map();
  let ans = Infinity;

  for (let i = 0; i < nums.length; i++) {
    // If current number exists in map,
    // then we found a mirror pair
    if (lastIndex.has(nums[i])) {
      ans = Math.min(ans, i - lastIndex.get(nums[i]));
    }

    // Store reverse(nums[i]) with current index
    const rev = reverseNum(nums[i]);
    lastIndex.set(rev, i);
  }

  return ans === Infinity ? -1 : ans;
};
