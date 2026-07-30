/**
 * @param {string} word
 * @return {number}
 */
var minimumPushes = function (word) {
  // Store the answer
  let pushes = 0;

  // Visit every character
  for (let i = 0; i < word.length; i++) {
    // Cost increases after every 8 letters
    pushes += Math.floor(i / 8) + 1;
  }

  // Return the minimum pushes
  return pushes;
};
