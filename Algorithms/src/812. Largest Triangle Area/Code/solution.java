class Solution {
    public double largestTriangleArea(int[][] points) {
        int n = points.length;
        double maxArea = 0.0;
        // check every triple i < j < k
        for (int i = 0; i < n - 2; i++) {
            for (int j = i + 1; j < n - 1; j++) {
                for (int k = j + 1; k < n; k++) {
                    int x1 = points[i][0], y1 = points[i][1];
                    int x2 = points[j][0], y2 = points[j][1];
                    int x3 = points[k][0], y3 = points[k][1];
                    // Compute doubled area using cross-product formula
                    double doubled = Math.abs(
                        (double)x1 * (y2 - y3) +
                        (double)x2 * (y3 - y1) +
                        (double)x3 * (y1 - y2)
                    );
                    double area = doubled * 0.5;
                    if (area > maxArea) maxArea = area;
                }
            }
        }
        return maxArea;
    }
}
