/**
 * @param {string[]} queries
 * @param {string[]} dictionary
 * @return {string[]}
 */
var twoEditWords = function (queries, dictionary) {
  const result = [];

  // Check every query word
  for (const query of queries) {
    // Compare with every dictionary word
    for (const word of dictionary) {
      let diff = 0;

      // Count character differences
      for (let i = 0; i < query.length; i++) {
        if (query[i] !== word[i]) {
          diff++;
        }

        // Stop if more than 2 edits are needed
        if (diff > 2) {
          break;
        }
      }

      // If current word matches within 2 edits
      if (diff <= 2) {
        result.push(query);
        break;
      }
    }
  }

  return result;
};
