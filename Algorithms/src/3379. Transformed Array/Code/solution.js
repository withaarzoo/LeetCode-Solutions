var constructTransformedArray = function (nums) {
  const n = nums.length;
  const result = new Array(n);

  for (let i = 0; i < n; i++) {
    if (nums[i] === 0) {
      result[i] = nums[i];
    } else {
      let target = (i + nums[i]) % n;
      if (target < 0) target += n;
      result[i] = nums[target];
    }
  }
  return result;
};
