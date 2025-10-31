/**
 * @param {number[]} nums
 * @return {number[]}
 */
var getSneakyNumbers = function (nums) {
  const n = nums.length - 2; // original range size
  const seen = new Array(n).fill(false); // boolean marks
  const res = [];
  for (const x of nums) {
    if (seen[x]) {
      res.push(x); // duplicate found
      if (res.length === 2) break; // stop early if both found
    } else {
      seen[x] = true; // mark as seen
    }
  }
  return res;
};
