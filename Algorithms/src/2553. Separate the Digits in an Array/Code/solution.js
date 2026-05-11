/**
 * @param {number[]} nums
 * @return {number[]}
 */
var separateDigits = function (nums) {
  // Final array to store separated digits
  let result = [];

  // Traverse every number in nums
  for (let num of nums) {
    // Convert number into string
    let str = num.toString();

    // Traverse every character
    for (let ch of str) {
      // Convert character into number and store it
      result.push(Number(ch));
    }
  }

  // Return final array
  return result;
};
