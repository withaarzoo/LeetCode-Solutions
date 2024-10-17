/**
 * @param {number} num
 * @return {number}
 */
var maximumSwap = function (num) {
  // Convert the number to a string array
  let numArr = num.toString().split("");

  // Track the last occurrence of each digit
  let last = new Array(10).fill(-1);
  for (let i = 0; i < numArr.length; i++) {
    last[parseInt(numArr[i])] = i;
  }

  // Traverse the digits from left to right
  for (let i = 0; i < numArr.length; i++) {
    // Check if a larger digit can be swapped later
    for (let d = 9; d > numArr[i]; d--) {
      if (last[d] > i) {
        // Swap the digits and return the new number
        [numArr[i], numArr[last[d]]] = [numArr[last[d]], numArr[i]];
        return parseInt(numArr.join(""));
      }
    }
  }

  // Return the original number if no swap was performed
  return num;
};
