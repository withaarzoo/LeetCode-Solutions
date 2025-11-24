/**
 * @param {number[]} nums
 * @return {boolean[]}
 */
var prefixesDivBy5 = function (nums) {
  const ans = [];
  let rem = 0; // remainder of current prefix modulo 5

  for (const bit of nums) {
    // binary shift: number = number * 2 + bit
    rem = (rem * 2 + bit) % 5;

    // if remainder is 0, divisible by 5
    ans.push(rem === 0);
  }

  return ans;
};
