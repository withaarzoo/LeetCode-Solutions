/**
 * @param {string} directions
 * @return {number}
 */
var countCollisions = function (directions) {
  const n = directions.length;
  let i = 0,
    j = n - 1;

  // Skip leading 'L' cars - they never collide
  while (i < n && directions[i] === "L") {
    i++;
  }

  // Skip trailing 'R' cars - they never collide
  while (j >= 0 && directions[j] === "R") {
    j--;
  }

  let collisions = 0;
  // Every non-'S' in the middle will collide exactly once
  for (let k = i; k <= j; k++) {
    if (directions[k] !== "S") {
      collisions++;
    }
  }

  return collisions;
};
