/**
 * @param {string[]} words
 * @return {number}
 */
var countPrefixSuffixPairs = function (words) {
  let count = 0;
  let n = words.length;

  for (let i = 0; i < n; i++) {
    for (let j = i + 1; j < n; j++) {
      let prefix = words[i];
      let word = words[j];
      let len = prefix.length;

      if (word.startsWith(prefix) && word.endsWith(prefix)) {
        count++;
      }
    }
  }

  return count;
};
