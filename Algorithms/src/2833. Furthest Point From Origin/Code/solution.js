/**
 * @param {string} moves
 * @return {number}
 */
var furthestDistanceFromOrigin = function (moves) {
  let left = 0,
    right = 0,
    blank = 0;

  for (let c of moves) {
    if (c === "L") left++;
    else if (c === "R") right++;
    else blank++;
  }

  let position = right - left;
  return Math.abs(position) + blank;
};
