class Solution {
    private int largestRectangleArea(int[] heights) {
        Stack<Integer> stack = new Stack<>();
        int maxArea = 0;
        int n = heights.length;
        int[] arr = new int[n + 1];
        System.arraycopy(heights, 0, arr, 0, n);

        for (int i = 0; i <= n; i++) {
            while (!stack.isEmpty() && arr[stack.peek()] > arr[i]) {
                int h = arr[stack.pop()];
                int w = stack.isEmpty() ? i : i - stack.peek() - 1;
                maxArea = Math.max(maxArea, h * w);
            }
            stack.push(i);
        }
        return maxArea;
    }

    public int maximalRectangle(char[][] matrix) {
        if (matrix.length == 0)
            return 0;

        int cols = matrix[0].length;
        int[] heights = new int[cols];
        int ans = 0;

        for (char[] row : matrix) {
            for (int j = 0; j < cols; j++) {
                heights[j] = row[j] == '1' ? heights[j] + 1 : 0;
            }
            ans = Math.max(ans, largestRectangleArea(heights));
        }
        return ans;
    }
}
