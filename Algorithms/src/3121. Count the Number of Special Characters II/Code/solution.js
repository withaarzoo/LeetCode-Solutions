/**
 * @param {string} word
 * @return {number}
 */
var numberOfSpecialChars = function (word) {
  // Store last occurrence of lowercase letters
  const lower = new Array(26).fill(-1);

  // Store first occurrence of uppercase letters
  const upper = new Array(26).fill(-1);

  // Traverse the string
  for (let i = 0; i < word.length; i++) {
    const ch = word[i];

    // If lowercase letter
    if (ch >= "a" && ch <= "z") {
      // Update last occurrence
      lower[ch.charCodeAt(0) - 97] = i;
    } else {
      const idx = ch.charCodeAt(0) - 65;

      // Store only first occurrence
      if (upper[idx] === -1) {
        upper[idx] = i;
      }
    }
  }

  let ans = 0;

  // Check all letters
  for (let i = 0; i < 26; i++) {
    // Both lowercase and uppercase must exist
    if (lower[i] !== -1 && upper[i] !== -1) {
      // Lowercase must come before uppercase
      if (lower[i] < upper[i]) {
        ans++;
      }
    }
  }

  return ans;
};
