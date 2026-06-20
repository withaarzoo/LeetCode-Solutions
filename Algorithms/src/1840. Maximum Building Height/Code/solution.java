class Solution {
    public int maxBuilding(int n, int[][] restrictions) {
        int m = restrictions.length;

        // Create new array including building 1 and building n
        int[][] arr = new int[m + 2][2];

        for (int i = 0; i < m; i++) {
            arr[i] = restrictions[i];
        }

        // Building 1 has fixed height 0
        arr[m] = new int[] { 1, 0 };

        // Building n can be at most n - 1
        arr[m + 1] = new int[] { n, n - 1 };

        // Sort by building index
        Arrays.sort(arr, (a, b) -> a[0] - b[0]);

        int size = arr.length;

        // Left to right pass
        for (int i = 1; i < size; i++) {
            int dist = arr[i][0] - arr[i - 1][0];

            arr[i][1] = Math.min(
                    arr[i][1],
                    arr[i - 1][1] + dist);
        }

        // Right to left pass
        for (int i = size - 2; i >= 0; i--) {
            int dist = arr[i + 1][0] - arr[i][0];

            arr[i][1] = Math.min(
                    arr[i][1],
                    arr[i + 1][1] + dist);
        }

        long ans = 0;

        // Calculate peak for every interval
        for (int i = 1; i < size; i++) {
            long x1 = arr[i - 1][0];
            long h1 = arr[i - 1][1];

            long x2 = arr[i][0];
            long h2 = arr[i][1];

            long dist = x2 - x1;

            long peak = Math.max(h1, h2) +
                    (dist - Math.abs(h1 - h2)) / 2;

            ans = Math.max(ans, peak);
        }

        return (int) ans;
    }
}