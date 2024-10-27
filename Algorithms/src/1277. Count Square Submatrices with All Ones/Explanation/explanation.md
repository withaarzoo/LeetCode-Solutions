# Count Square Submatrices with All Ones - Step-by-Step Solution

In this guide, we break down the solution to the problem "Count Square Submatrices with All Ones" in five programming languages: C++, Java, JavaScript, Python, and Go. The goal is to provide a clear understanding of each step, showing how we build an optimized dynamic programming (DP) solution.

## Problem Summary

Given a matrix of `1`s and `0`s, count the number of square submatrices that contain only `1`s. Each square submatrix should be entirely made up of `1`s, and the solution should efficiently handle large matrices.

## Solution Overview

1. **Initialize DP Array**: We create a DP array where each entry `dp[i][j]` represents the side length of the largest square submatrix that ends at position `(i, j)`.
  
2. **Set Up Base Conditions**: For cells in the first row or first column, the value of `dp[i][j]` is simply the same as `matrix[i][j]`, since squares in these positions can only be `1x1` squares.

3. **Fill DP Array**:
   - For other cells, if `matrix[i][j]` is `1`, we calculate `dp[i][j]` as `min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1`. This allows us to extend the size of squares based on neighboring cells.
   - If `matrix[i][j]` is `0`, `dp[i][j]` is `0`.

4. **Count Total Squares**: By summing up all values in `dp`, we obtain the total number of square submatrices with all `1`s.

## Language-Specific Steps

### C++ Code

1. **Define a DP Array**: Create a 2D DP array of the same size as the input matrix.
2. **Loop Through Matrix Cells**:
   - If the current cell is `1` and not on the first row or column, set `dp[i][j]` based on the minimum of the three neighboring cells (top, left, and top-left).
   - If the cell is on the first row or column and is `1`, set `dp[i][j]` to `1`.
3. **Sum the DP Array**: Each entry in `dp` contributes to the final count, so sum them up to get the total number of square submatrices.

### Java Code

1. **Initialize DP Array**: Define a 2D integer array `dp` to store the maximum square side length at each position.
2. **Iterate Through the Matrix**:
   - For each cell, if it contains `1`, check if it’s in the first row or column. If so, set `dp[i][j]` to `1`.
   - Otherwise, set `dp[i][j]` to `min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1` if it’s in a square of `1`s.
3. **Calculate Total Squares**: Add up all entries in `dp` to get the result.

### JavaScript Code

1. **Create DP Array**: Use a 2D array initialized to `0` for tracking square sizes at each cell.
2. **Loop Over the Matrix**:
   - For cells with `1`, if they’re in the first row or column, set `dp[i][j]` to `1`.
   - For other cells with `1`, compute the side length by taking `min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1`.
3. **Sum All Squares**: After processing the matrix, sum all values in `dp` to get the final count of square submatrices.

### Python Code

1. **Initialize a DP Matrix**: Create a 2D list, `dp`, that mirrors the input matrix to store square sizes at each cell.
2. **Traverse the Matrix**:
   - For each cell containing `1`, if it’s in the first row or column, set `dp[i][j]` to `1`.
   - For other cells containing `1`, calculate `dp[i][j]` as `min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1`.
3. **Aggregate Square Counts**: Sum all values in `dp` for the total count of square submatrices.

### Go Code

1. **Setup the DP Array**: Create a 2D array of integers to track the side lengths of squares ending at each cell.
2. **Loop Through Each Cell**:
   - For cells in the first row or column containing `1`, set `dp[i][j]` to `1`.
   - For other cells containing `1`, set `dp[i][j]` to `min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1`.
3. **Sum the Result**: Loop over `dp` to sum up the entries and get the total count of square submatrices.
