/**
 * @param {character[][]} boxGrid
 * @return {character[][]}
 */
var rotateTheBox = function (boxGrid) {
  const m = boxGrid.length;
  const n = boxGrid[0].length;

  // Process each row
  for (let row = 0; row < m; row++) {
    // Rightmost valid empty position
    let emptyCol = n - 1;

    // Traverse from right to left
    for (let col = n - 1; col >= 0; col--) {
      // Obstacle found
      if (boxGrid[row][col] === "*") {
        // Reset empty position
        emptyCol = col - 1;
      }

      // Stone found
      else if (boxGrid[row][col] === "#") {
        // Remove current stone
        boxGrid[row][col] = ".";

        // Move stone
        boxGrid[row][emptyCol] = "#";

        // Update next empty spot
        emptyCol--;
      }
    }
  }

  // Create rotated matrix
  const rotated = Array.from({ length: n }, () => Array(m).fill("."));

  // Rotate clockwise
  for (let i = 0; i < m; i++) {
    for (let j = 0; j < n; j++) {
      rotated[j][m - 1 - i] = boxGrid[i][j];
    }
  }

  return rotated;
};
