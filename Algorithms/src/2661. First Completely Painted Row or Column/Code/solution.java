class Solution {
    public int firstCompleteIndex(int[] arr, int[][] mat) {
        int m = mat.length, n = mat[0].length;
        Map<Integer, int[]> position = new HashMap<>();
        int[] rowCount = new int[m];
        int[] colCount = new int[n];

        // Map matrix values to their positions
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                position.put(mat[i][j], new int[] { i, j });
            }
        }

        // Iterate through the array and simulate painting
        for (int i = 0; i < arr.length; i++) {
            int[] pos = position.get(arr[i]);
            int row = pos[0], col = pos[1];
            rowCount[row]++;
            colCount[col]++;

            if (rowCount[row] == n || colCount[col] == m) {
                return i;
            }
        }
        return -1;
    }
}
