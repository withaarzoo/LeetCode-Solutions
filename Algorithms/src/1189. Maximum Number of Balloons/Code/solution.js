/**
 * @param {string} text
 * @return {number}
 */
var maxNumberOfBalloons = function (text) {
  // Store frequency of all lowercase letters
  const freq = new Array(26).fill(0);

  // Count each character
  for (const ch of text) {
    freq[ch.charCodeAt(0) - 97]++;
  }

  // Return the limiting character count
  return Math.min(
    freq["b".charCodeAt(0) - 97], // Need 1 'b'
    freq["a".charCodeAt(0) - 97], // Need 1 'a'
    Math.floor(freq["l".charCodeAt(0) - 97] / 2), // Need 2 'l'
    Math.floor(freq["o".charCodeAt(0) - 97] / 2), // Need 2 'o'
    freq["n".charCodeAt(0) - 97], // Need 1 'n'
  );
};
