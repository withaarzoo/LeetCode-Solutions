/**
 * @param {number[][]} points - A 2D array where points[i][j] represents the score at row i and column j.
 * @return {number} - The maximum points that can be obtained.
 */
var maxPoints = function (points) {
  // m represents the number of rows in the points matrix
  const m = points.length;
  // n represents the number of columns in the points matrix
  const n = points[0].length;

  // Initialize dp array to store the maximum points for the first row
  let dp = Array(n).fill(0);

  // Fill the dp array with values from the first row of points
  for (let j = 0; j < n; ++j) {
    dp[j] = points[0][j];
  }

  // Traverse through each subsequent row to calculate the maximum points
  for (let i = 1; i < m; ++i) {
    // Arrays to store the maximum values when moving from left to right and right to left
    const leftMax = Array(n).fill(0);
    const rightMax = Array(n).fill(0);
    // Array to store the new dp values for the current row
    const newDp = Array(n).fill(0);

    // Calculate the maximum points for each column when moving from left to right
    leftMax[0] = dp[0]; // The first element remains the same
    for (let j = 1; j < n; ++j) {
      // Maximum of the previous leftMax or current dp value adjusted by column index
      leftMax[j] = Math.max(leftMax[j - 1], dp[j] + j);
    }

    // Calculate the maximum points for each column when moving from right to left
    rightMax[n - 1] = dp[n - 1] - (n - 1); // The last element is adjusted by its index
    for (let j = n - 2; j >= 0; --j) {
      // Maximum of the next rightMax or current dp value adjusted by column index
      rightMax[j] = Math.max(rightMax[j + 1], dp[j] - j);
    }

    // Calculate the new dp values for the current row based on leftMax and rightMax
    for (let j = 0; j < n; ++j) {
      // The value for each dp[j] is the maximum of leftMax and rightMax adjusted by index
      // plus the current points value
      newDp[j] = Math.max(leftMax[j] - j, rightMax[j] + j) + points[i][j];
    }

    // Update the dp array to the newly calculated values for the next iteration
    dp = newDp;
  }

  // Return the maximum value in the dp array which represents the maximum points
  return Math.max(...dp);
};
