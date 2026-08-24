/**
 * @param {number[]} stones
 * @return {number}
 */
var stoneGameVIII = function (stones) {
  // Convert the original values into prefix sums in the same array.
  for (let i = 1; i < stones.length; i++) {
    stones[i] += stones[i - 1];
  }

  // Start with the state where the current player takes all stones.
  let best = stones[stones.length - 1];

  // Move backward through every valid prefix choice.
  for (let i = stones.length - 2; i >= 1; i--) {
    // Keep the current answer or take this prefix and
    // subtract the best result available to the opponent.
    best = Math.max(best, stones[i] - best);
  }

  // Return Alice's maximum possible score difference.
  return best;
};
