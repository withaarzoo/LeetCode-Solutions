class Solution {
    public int largestOverlap(int[][] img1, int[][] img2) {
        int n = img1.length;
        // collect every coordinate that holds a 1
        List<int[]> A = new ArrayList<>();
        List<int[]> B = new ArrayList<>();
        for (int i = 0; i < n; ++i) {
            for (int j = 0; j < n; ++j) {
                if (img1[i][j] == 1)
                    A.add(new int[] { i, j });
                if (img2[i][j] == 1)
                    B.add(new int[] { i, j });
            }
        }
        // count how many times each possible shift appears
        // shifts range from -(n-1) to +(n-1), so offset by n
        int[][] cnt = new int[2 * n][2 * n];
        int best = 0;
        for (int[] a : A) {
            for (int[] b : B) {
                int dx = b[0] - a[0] + n; // make index non-negative
                int dy = b[1] - a[1] + n;
                best = Math.max(best, ++cnt[dx][dy]);
            }
        }
        return best;
    }
}