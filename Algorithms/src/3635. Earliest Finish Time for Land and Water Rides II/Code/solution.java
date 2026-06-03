class Solution {

    // Computes the best answer when category A is taken first
    private long solve(int[] startA, int[] durA,
            int[] startB, int[] durB) {

        int m = startB.length;

        int[][] rides = new int[m][2];

        // Store (start, duration)
        for (int i = 0; i < m; i++) {
            rides[i][0] = startB[i];
            rides[i][1] = durB[i];
        }

        // Sort by start time
        java.util.Arrays.sort(rides, (a, b) -> Integer.compare(a[0], b[0]));

        int[] starts = new int[m];
        long[] prefixMinDur = new long[m];
        long[] suffixMinFinish = new long[m];

        for (int i = 0; i < m; i++) {
            starts[i] = rides[i][0];

            if (i == 0)
                prefixMinDur[i] = rides[i][1];
            else
                prefixMinDur[i] = Math.min(prefixMinDur[i - 1], rides[i][1]);
        }

        for (int i = m - 1; i >= 0; i--) {
            long finish = (long) rides[i][0] + rides[i][1];

            if (i == m - 1)
                suffixMinFinish[i] = finish;
            else
                suffixMinFinish[i] = Math.min(suffixMinFinish[i + 1], finish);
        }

        long ans = Long.MAX_VALUE;

        for (int i = 0; i < startA.length; i++) {

            long finish1 = (long) startA[i] + durA[i];

            int pos = upperBound(starts, finish1);

            if (pos > 0) {
                ans = Math.min(ans, finish1 + prefixMinDur[pos - 1]);
            }

            if (pos < m) {
                ans = Math.min(ans, suffixMinFinish[pos]);
            }
        }

        return ans;
    }

    // First index with value > target
    private int upperBound(int[] arr, long target) {
        int left = 0;
        int right = arr.length;

        while (left < right) {
            int mid = left + (right - left) / 2;

            if (arr[mid] <= target)
                left = mid + 1;
            else
                right = mid;
        }

        return left;
    }

    public int earliestFinishTime(int[] landStartTime, int[] landDuration, int[] waterStartTime, int[] waterDuration) {

        long ans1 = solve(
                landStartTime, landDuration,
                waterStartTime, waterDuration);

        long ans2 = solve(
                waterStartTime, waterDuration,
                landStartTime, landDuration);

        return (int) Math.min(ans1, ans2);
    }
}