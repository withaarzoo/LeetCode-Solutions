/**
 * @param {string[]} words
 * @param {number[]} weights
 * @return {string}
 */
var mapWordWeights = function (words, weights) {
  let result = "";

  // Process every word
  for (const word of words) {
    let sumWeight = 0;

    // Add character weights
    for (const ch of word) {
      sumWeight += weights[ch.charCodeAt(0) - 97];
    }

    // Reduce into range [0, 25]
    const value = sumWeight % 26;

    // Reverse alphabet mapping
    result += String.fromCharCode("z".charCodeAt(0) - value);
  }

  return result;
};
