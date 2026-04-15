/**
 * @param {string[]} words
 * @param {string} target
 * @param {number} startIndex
 * @return {number}
 */
var closestTarget = function (words, target, startIndex) {
  const n = words.length;
  let ans = Infinity;

  // Traverse the array
  for (let i = 0; i < n; i++) {
    // If target is found
    if (words[i] === target) {
      const diff = Math.abs(i - startIndex);

      // Distance if we move around the circle
      const circularDist = n - diff;

      // Take minimum possible distance
      ans = Math.min(ans, Math.min(diff, circularDist));
    }
  }

  // If target was not found
  return ans === Infinity ? -1 : ans;
};
