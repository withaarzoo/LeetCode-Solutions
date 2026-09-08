/**
 * @param {number} n
 * @return {number}
 */
var countCommas = function (n) {
  let answer = 0; // Stores the total number of commas.
  let threshold = 1000; // First number that contains a comma.

  while (threshold <= n) {
    // Every number from threshold through n contributes
    // one comma for this comma group.
    answer += n - threshold + 1;

    // Move to the next comma level.
    threshold *= 1000;
  }

  // Return the total number of commas.
  return answer;
};
