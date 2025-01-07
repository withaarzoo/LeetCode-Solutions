/**
 * @param {string[]} words
 * @return {string[]}
 */
var stringMatching = function (words) {
  // Sort words by length
  words.sort((a, b) => a.length - b.length);

  const result = [];

  // Check for substrings
  for (let i = 0; i < words.length; i++) {
    for (let j = i + 1; j < words.length; j++) {
      if (words[j].includes(words[i])) {
        result.push(words[i]);
        break;
      }
    }
  }

  return result;
};
