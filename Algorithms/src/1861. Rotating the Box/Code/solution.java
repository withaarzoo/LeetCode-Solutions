class Solution {
    public char[][] rotateTheBox(char[][] boxGrid) {

        int m = boxGrid.length;
        int n = boxGrid[0].length;

        // Process every row
        for (int row = 0; row < m; row++) {

            // Rightmost empty position
            int emptyCol = n - 1;

            // Traverse row from right to left
            for (int col = n - 1; col >= 0; col--) {

                // Obstacle found
                if (boxGrid[row][col] == '*') {

                    // Reset valid falling position
                    emptyCol = col - 1;
                }

                // Stone found
                else if (boxGrid[row][col] == '#') {

                    // Remove stone
                    boxGrid[row][col] = '.';

                    // Move stone to valid position
                    boxGrid[row][emptyCol] = '#';

                    // Update next empty spot
                    emptyCol--;
                }
            }
        }

        // Rotated matrix
        char[][] rotated = new char[n][m];

        // Rotate clockwise
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {

                rotated[j][m - 1 - i] = boxGrid[i][j];
            }
        }

        return rotated;
    }
}