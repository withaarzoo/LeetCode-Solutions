/**
 * @param {string} word
 * @return {number}
 */
var numberOfSpecialChars = function (word) {
  // Store all characters inside a Set
  const set = new Set(word);

  // Variable to store answer
  let count = 0;

  // Loop through all lowercase letters
  for (let i = 0; i < 26; i++) {
    // Current lowercase character
    let lower = String.fromCharCode(97 + i);

    // Corresponding uppercase character
    let upper = String.fromCharCode(65 + i);

    // If both exist, increase answer
    if (set.has(lower) && set.has(upper)) {
      count++;
    }
  }

  // Return total special characters
  return count;
};
