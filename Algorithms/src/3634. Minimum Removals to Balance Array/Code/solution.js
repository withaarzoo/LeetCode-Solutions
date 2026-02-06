var minRemoval = function (nums, k) {
  nums.sort((a, b) => a - b);

  let left = 0;
  let maxKeep = 1;

  for (let right = 0; right < nums.length; right++) {
    while (nums[right] > nums[left] * k) {
      left++;
    }
    maxKeep = Math.max(maxKeep, right - left + 1);
  }

  return nums.length - maxKeep;
};
