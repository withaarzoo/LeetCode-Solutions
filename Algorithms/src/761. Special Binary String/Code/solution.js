/**
 * @param {string} s
 * @return {string}
 */
var makeLargestSpecial = function (s) {
  let parts = [];
  let count = 0;
  let start = 0;

  for (let i = 0; i < s.length; i++) {
    if (s[i] === "1") count++;
    else count--;

    if (count === 0) {
      let inner = makeLargestSpecial(s.substring(start + 1, i));
      parts.push("1" + inner + "0");
      start = i + 1;
    }
  }

  // Sort descending
  parts.sort((a, b) => b.localeCompare(a));

  return parts.join("");
};
