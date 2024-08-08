import java.util.ArrayList;
import java.util.List;

class Solution {
    public int[][] spiralMatrixIII(int rows, int cols, int rStart, int cStart) {
        // List to store the coordinates in the order we visit them
        List<int[]> result = new ArrayList<>();

        // Directions for movement: right, down, left, up
        int[][] directions = { { 0, 1 }, { 1, 0 }, { 0, -1 }, { -1, 0 } };

        // Initial step count
        int steps = 1;

        // Direction index (0: right, 1: down, 2: left, 3: up)
        int d = 0;

        // Starting coordinates
        int r = rStart, c = cStart;

        // Add the starting position to the result
        result.add(new int[] { r, c });

        // Continue until we've added all positions in the matrix
        while (result.size() < rows * cols) {
            // Repeat twice for each direction pair
            for (int i = 0; i < 2; ++i) {
                // Move 'steps' times in the current direction
                for (int j = 0; j < steps; ++j) {
                    // Update the current position
                    r += directions[d][0];
                    c += directions[d][1];

                    // Check if the new position is within bounds
                    if (r >= 0 && r < rows && c >= 0 && c < cols) {
                        // If within bounds, add the position to the result
                        result.add(new int[] { r, c });
                    }
                }

                // Change direction (right -> down -> left -> up)
                d = (d + 1) % 4;
            }

            // After every two directions, increment the step count
            ++steps;
        }

        // Convert the result list to an array and return it
        return result.toArray(new int[result.size()][]);
    }
}
