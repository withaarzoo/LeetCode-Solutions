/**
 * @param {number} n
 * @return {number}
 */
var binaryGap = function (n) {
  let lastPosition = -1; // last seen 1 index
  let maxDistance = 0; // max gap
  let currentPosition = 0; // bit index

  while (n > 0) {
    // Check last bit
    if ((n & 1) === 1) {
      if (lastPosition !== -1) {
        maxDistance = Math.max(maxDistance, currentPosition - lastPosition);
      }
      lastPosition = currentPosition;
    }

    n = n >> 1; // shift right
    currentPosition++;
  }

  return maxDistance;
};
