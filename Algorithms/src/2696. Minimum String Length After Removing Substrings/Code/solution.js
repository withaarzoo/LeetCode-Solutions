/**
 * @param {string} s
 * @return {number}
 */
var minLength = function (s) {
  let stack = [];

  // Traverse through each character in the string
  for (let ch of s) {
    // Check if top of the stack forms "AB" or "CD" with current character
    if (
      stack.length > 0 &&
      ((stack[stack.length - 1] === "A" && ch === "B") ||
        (stack[stack.length - 1] === "C" && ch === "D"))
    ) {
      stack.pop(); // Pop to remove the pair
    } else {
      stack.push(ch); // Push current character onto the stack if no match
    }
  }

  return stack.length; // Remaining stack size is the minimum length of the string
};
