class Solution {
    public int maxSideLength(int[][] mat, int threshold) {
        int m = mat.length, n = mat[0].length;
        int[][] pre = new int[m + 1][n + 1];

        // Prefix sum
        for (int i = 1; i <= m; i++) {
            for (int j = 1; j <= n; j++) {
                pre[i][j] = mat[i - 1][j - 1]
                        + pre[i - 1][j]
                        + pre[i][j - 1]
                        - pre[i - 1][j - 1];
            }
        }

        int left = 0, right = Math.min(m, n), ans = 0;

        while (left <= right) {
            int mid = (left + right) / 2;
            boolean found = false;

            for (int i = mid; i <= m && !found; i++) {
                for (int j = mid; j <= n; j++) {
                    int sum = pre[i][j]
                            - pre[i - mid][j]
                            - pre[i][j - mid]
                            + pre[i - mid][j - mid];

                    if (sum <= threshold) {
                        found = true;
                        break;
                    }
                }
            }

            if (found) {
                ans = mid;
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }

        return ans;
    }
}
