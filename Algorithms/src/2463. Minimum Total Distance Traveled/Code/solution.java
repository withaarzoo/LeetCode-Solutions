class Solution {
    long[][] dp;
    long INF = (long) 1e18;

    private long solve(int i, int j, List<Integer> robot, int[][] factory) {
        int n = robot.size();
        int m = factory.length;

        // All robots repaired
        if (i == n)
            return 0;

        // No factories left
        if (j == m)
            return INF;

        if (dp[i][j] != -1)
            return dp[i][j];

        // Skip current factory
        long ans = solve(i, j + 1, robot, factory);

        long distance = 0;
        int pos = factory[j][0];
        int limit = factory[j][1];

        // Use current factory for next k robots
        for (int k = 0; k < limit && i + k < n; k++) {
            distance += Math.abs(robot.get(i + k) - pos);

            long next = solve(i + k + 1, j + 1, robot, factory);

            if (next != INF) {
                ans = Math.min(ans, distance + next);
            }
        }

        return dp[i][j] = ans;
    }

    public long minimumTotalDistance(List<Integer> robot, int[][] factory) {
        Collections.sort(robot);

        Arrays.sort(factory, (a, b) -> Integer.compare(a[0], b[0]));

        int n = robot.size();
        int m = factory.length;

        dp = new long[n + 1][m + 1];

        for (long[] row : dp) {
            Arrays.fill(row, -1);
        }

        return solve(0, 0, robot, factory);
    }
}