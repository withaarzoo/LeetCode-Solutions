class Solution {
    public long maxMatrixSum(int[][] matrix) {
        long sum = 0;
        int negativeCount = 0;
        int minAbs = Integer.MAX_VALUE;

        for (int[] row : matrix) {
            for (int val : row) {
                sum += Math.abs(val); // add absolute value
                if (val < 0)
                    negativeCount++;
                minAbs = Math.min(minAbs, Math.abs(val));
            }
        }

        if (negativeCount % 2 == 1) {
            sum -= 2L * minAbs;
        }

        return sum;
    }
}
