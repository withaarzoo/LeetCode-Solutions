/**
 * @param {string} s
 * @param {string} locked
 * @return {boolean}
 */
var canBeValid = function (s, locked) {
  if (s.length % 2 !== 0) return false; // Odd length can't be balanced

  let open = 0,
    flexible = 0;
  // Left-to-right pass
  for (let i = 0; i < s.length; i++) {
    if (locked[i] === "1") {
      open += s[i] === "(" ? 1 : -1;
    } else {
      flexible++;
    }
    if (open + flexible < 0) return false;
  }

  open = 0;
  flexible = 0;
  // Right-to-left pass
  for (let i = s.length - 1; i >= 0; i--) {
    if (locked[i] === "1") {
      open += s[i] === ")" ? 1 : -1;
    } else {
      flexible++;
    }
    if (open + flexible < 0) return false;
  }

  return true;
};
