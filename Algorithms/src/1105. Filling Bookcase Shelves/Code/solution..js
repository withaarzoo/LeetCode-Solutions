/**
 * Function to find the minimum height of shelves required to place all books
 * @param {number[][]} books - A list of books, each represented as [width, height]
 * @param {number} shelfWidth - The maximum width allowed for each shelf
 * @return {number} - The minimum height required to place all books on the shelves
 */
var minHeightShelves = function (books, shelfWidth) {
  let n = books.length; // Total number of books
  let dp = new Array(n + 1).fill(Infinity); // DP array to store minimum height at each book
  dp[0] = 0; // Base case: no books placed means height is 0

  // Iterate over each book
  for (let i = 1; i <= n; ++i) {
    let width = 0; // Current shelf width used
    let height = 0; // Current shelf maximum height

    // Check each possible end of the previous row
    for (let j = i; j > 0; --j) {
      width += books[j - 1][0]; // Add the width of the j-th book
      if (width > shelfWidth) break; // If the shelf exceeds the max width, break

      height = Math.max(height, books[j - 1][1]); // Update the shelf height to the tallest book
      // Update the minimum height to place all books up to the i-th book
      dp[i] = Math.min(dp[i], dp[j - 1] + height);
    }
  }

  return dp[n]; // Return the minimum height for all books
};
