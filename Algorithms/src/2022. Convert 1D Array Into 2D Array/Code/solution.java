class Solution {
    public int[][] construct2DArray(int[] original, int m, int n) {
        // Step 1: Check if reshaping the 1D array into a 2D array is possible
        // The reshaping is only possible if the total number of elements in the 1D
        // array
        // matches the total number of elements required in the 2D array (m * n).
        if (original.length != m * n) {
            // If the total number of elements does not match, return an empty 2D array.
            return new int[0][0];
        }

        // Step 2: Initialize the result 2D array with dimensions m (rows) x n (columns)
        int[][] result = new int[m][n];

        // Step 3: Fill the 2D array with elements from the 1D array
        // Iterate through each element in the 1D array.
        for (int i = 0; i < original.length; i++) {
            // Calculate the row index in the 2D array using integer division (i / n).
            // Calculate the column index in the 2D array using the modulus operation (i %
            // n).
            // Assign the value from the original 1D array to the corresponding position in
            // the 2D array.
            result[i / n][i % n] = original[i];
        }

        // Step 4: Return the constructed 2D array as the final result
        return result;
    }
}
