/**
 * @param {string} s
 * @return {string}
 */
var processStr = function (s) {
  // Stores the current result being built
  let result = "";

  for (const c of s) {
    // Lowercase letter -> append to result
    if (c >= "a" && c <= "z") {
      result += c;
    }
    // Remove last character if it exists
    else if (c === "*") {
      if (result.length > 0) {
        result = result.slice(0, -1);
      }
    }
    // Duplicate current result
    else if (c === "#") {
      result += result;
    }
    // Reverse current result
    else if (c === "%") {
      result = result.split("").reverse().join("");
    }
  }

  return result;
};
