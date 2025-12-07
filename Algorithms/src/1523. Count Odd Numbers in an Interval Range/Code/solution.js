/**
 * @param {number} low
 * @param {number} high
 * @return {number}
 */
var countOdds = function (low, high) {
  // Helper: count of odd numbers from 1 to x
  const oddsUpTo = (x) => {
    // Math.floor because JS division returns float
    return Math.floor((x + 1) / 2);
  };

  // Odds in [low, high]
  return oddsUpTo(high) - oddsUpTo(low - 1);
};
