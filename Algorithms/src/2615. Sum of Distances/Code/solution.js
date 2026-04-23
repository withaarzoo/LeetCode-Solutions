var distance = function (nums) {
  const n = nums.length;
  const map = new Map();

  // Group indices
  for (let i = 0; i < n; i++) {
    if (!map.has(nums[i])) map.set(nums[i], []);
    map.get(nums[i]).push(i);
  }

  const res = new Array(n).fill(0);

  for (let idx of map.values()) {
    let k = idx.length;

    let prefixSum = 0;
    let totalSum = idx.reduce((a, b) => a + b, 0);

    for (let i = 0; i < k; i++) {
      let curr = idx[i];

      let left = curr * i - prefixSum;
      let right = totalSum - prefixSum - curr - curr * (k - i - 1);

      res[curr] = left + right;

      prefixSum += curr;
    }
  }

  return res;
};
