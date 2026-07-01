/**
 * @param {string} s
 * @return {number}
 */
var numberOfSubstrings = function (s) {
  // Store the frequency of 'a', 'b', and 'c'
  const freq = [0, 0, 0];

  let left = 0;
  let ans = 0;
  const n = s.length;

  // Expand the window
  for (let right = 0; right < n; right++) {
    // Add the current character
    freq[s.charCodeAt(right) - 97]++;

    // Shrink while all characters are present
    while (freq[0] > 0 && freq[1] > 0 && freq[2] > 0) {
      // Every larger ending index is also valid
      ans += n - right;

      // Remove the leftmost character
      freq[s.charCodeAt(left) - 97]--;

      // Move left forward
      left++;
    }
  }

  // Return the answer
  return ans;
};
