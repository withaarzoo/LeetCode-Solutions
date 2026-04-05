/**
 * @param {string} moves
 * @return {boolean}
 */
var judgeCircle = function (moves) {
  // x represents left/right position
  // y represents up/down position
  let x = 0,
    y = 0;

  // Traverse every move
  for (let move of moves) {
    if (move === "U") {
      y++; // Move up
    } else if (move === "D") {
      y--; // Move down
    } else if (move === "R") {
      x++; // Move right
    } else if (move === "L") {
      x--; // Move left
    }
  }

  // Robot returns to origin only if both coordinates are 0
  return x === 0 && y === 0;
};
