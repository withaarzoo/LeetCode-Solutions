/**
 * @param {string} s
 * @return {string}
 */
var smallestPalindrome = function (s) {
  // Store frequency of every lowercase letter
  const freq = new Array(26).fill(0);

  // Count character frequencies
  for (const ch of s) freq[ch.charCodeAt(0) - 97]++;

  let left = "";
  let middle = "";

  // Build the left half and find the middle character
  for (let i = 0; i < 26; i++) {
    // Add half of the occurrences to the left side
    left += String.fromCharCode(97 + i).repeat(Math.floor(freq[i] / 2));

    // Odd frequency character becomes the center
    if (freq[i] % 2) middle = String.fromCharCode(97 + i);
  }

  // Right half is the reverse of the left half
  const right = left.split("").reverse().join("");

  // Return the complete palindrome
  return left + middle + right;
};
