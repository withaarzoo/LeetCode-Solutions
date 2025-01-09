/**
 * @param {string[]} words
 * @param {string} pref
 * @return {number}
 */
var prefixCount = function (words, pref) {
  let count = 0;
  for (let word of words) {
    // Check if the word starts with the prefix
    if (word.startsWith(pref)) {
      count++;
    }
  }
  return count;
};
