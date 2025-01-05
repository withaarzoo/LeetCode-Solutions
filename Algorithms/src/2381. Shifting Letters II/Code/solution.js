/**
 * @param {string} s
 * @param {number[][]} shifts
 * @return {string}
 */
var shiftingLetters = function (s, shifts) {
  const n = s.length;
  const diff = Array(n + 1).fill(0);

  // Build the difference array
  for (const [start, end, direction] of shifts) {
    const delta = direction === 1 ? 1 : -1;
    diff[start] += delta;
    if (end + 1 < n) diff[end + 1] -= delta;
  }

  // Calculate cumulative shifts
  let shift = 0;
  const result = s.split("");
  for (let i = 0; i < n; i++) {
    shift += diff[i];
    shift = ((shift % 26) + 26) % 26; // Normalize shift to [0, 25]
    result[i] = String.fromCharCode(
      "a".charCodeAt(0) +
        ((result[i].charCodeAt(0) - "a".charCodeAt(0) + shift) % 26)
    );
  }

  return result.join("");
};
