/**
 * @param {string[]} words1
 * @param {string[]} words2
 * @return {string[]}
 */
var wordSubsets = function (words1, words2) {
  let maxFreq = new Array(26).fill(0);

  // Precompute the maximum frequency for each character in words2
  for (let word of words2) {
    let freq = new Array(26).fill(0);
    for (let char of word) freq[char.charCodeAt(0) - 97]++;
    for (let i = 0; i < 26; i++) {
      maxFreq[i] = Math.max(maxFreq[i], freq[i]);
    }
  }

  let result = [];
  // Check each word in words1
  for (let word of words1) {
    let freq = new Array(26).fill(0);
    for (let char of word) freq[char.charCodeAt(0) - 97]++;
    let isUniversal = true;
    for (let i = 0; i < 26; i++) {
      if (freq[i] < maxFreq[i]) {
        isUniversal = false;
        break;
      }
    }
    if (isUniversal) result.push(word);
  }

  return result;
};
