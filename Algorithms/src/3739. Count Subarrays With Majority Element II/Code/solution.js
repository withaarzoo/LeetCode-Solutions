/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var countMajoritySubarrays = function (nums, target) {
  const n = nums.length;

  // Prefix sums after converting target -> +1 and others -> -1
  const pref = new Array(n + 1).fill(0);

  for (let i = 0; i < n; i++) {
    pref[i + 1] = pref[i] + (nums[i] === target ? 1 : -1);
  }

  // Coordinate compression
  const values = [...new Set([...pref].sort((a, b) => a - b))];

  // Fenwick Tree
  const bit = new Array(values.length + 2).fill(0);

  function update(idx) {
    while (idx < bit.length) {
      bit[idx]++;
      idx += idx & -idx;
    }
  }

  function query(idx) {
    let sum = 0;
    while (idx > 0) {
      sum += bit[idx];
      idx -= idx & -idx;
    }
    return sum;
  }

  function lowerBound(arr, x) {
    let l = 0,
      r = arr.length;
    while (l < r) {
      const mid = (l + r) >> 1;
      if (arr[mid] < x) l = mid + 1;
      else r = mid;
    }
    return l;
  }

  let ans = 0;

  for (const x of pref) {
    const idx = lowerBound(values, x) + 1;

    ans += query(idx - 1);

    update(idx);
  }

  return ans;
};
