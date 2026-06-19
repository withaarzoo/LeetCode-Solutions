/**
 * @param {number[]} gain
 * @return {number}
 */
var largestAltitude = function (gain) {
  // Current altitude starts at 0
  let currentAltitude = 0;

  // Highest altitude seen so far
  let maxAltitude = 0;

  // Process every gain value
  for (const change of gain) {
    // Apply altitude change
    currentAltitude += change;

    // Update highest altitude if current altitude is larger
    maxAltitude = Math.max(maxAltitude, currentAltitude);
  }

  // Return the highest altitude reached
  return maxAltitude;
};
