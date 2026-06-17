/**
 * @param {string} s
 * @param {number} k
 * @return {character}
 */
var processStr = function (s, k) {
  const n = s.length;

  // Use BigInt because lengths can reach 1e15
  const len = new Array(n);
  let curLen = 0n;

  for (let i = 0; i < n; i++) {
    const c = s[i];

    if (c >= "a" && c <= "z") {
      // Append character
      curLen++;
    } else if (c === "*") {
      // Remove last character if it exists
      if (curLen > 0n) curLen--;
    } else if (c === "#") {
      // Duplicate string
      curLen *= 2n;
    } else {
      // '%' does not change length
    }

    len[i] = curLen;
  }

  let idx = BigInt(k);

  // Out of bounds
  if (idx >= curLen) return ".";

  // Undo operations
  for (let i = n - 1; i >= 0; i--) {
    const c = s[i];
    const before = i === 0 ? 0n : len[i - 1];

    if (c >= "a" && c <= "z") {
      // Letter was appended at position "before"
      if (idx === before) return c;
    } else if (c === "#") {
      // Undo duplication
      if (before > 0n) idx %= before;
    } else if (c === "%") {
      // Undo reverse
      idx = before - 1n - idx;
    } else {
      // '*' keeps surviving indices unchanged
    }
  }

  return ".";
};
