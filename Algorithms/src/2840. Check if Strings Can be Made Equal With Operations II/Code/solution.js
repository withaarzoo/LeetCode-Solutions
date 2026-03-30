/**
 * @param {string} s1
 * @param {string} s2
 * @return {boolean}
 */
var checkStrings = function (s1, s2) {
  // Frequency arrays for even and odd positions
  const even = new Array(26).fill(0);
  const odd = new Array(26).fill(0);

  for (let i = 0; i < s1.length; i++) {
    if (i % 2 === 0) {
      // Count characters at even indexes
      even[s1.charCodeAt(i) - 97]++;
      even[s2.charCodeAt(i) - 97]--;
    } else {
      // Count characters at odd indexes
      odd[s1.charCodeAt(i) - 97]++;
      odd[s2.charCodeAt(i) - 97]--;
    }
  }

  // Check if all frequencies become zero
  for (let i = 0; i < 26; i++) {
    if (even[i] !== 0 || odd[i] !== 0) {
      return false;
    }
  }

  return true;
};
