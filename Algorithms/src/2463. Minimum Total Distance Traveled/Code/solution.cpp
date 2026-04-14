class Solution
{
public:
    long long dp[101][101];
    const long long INF = 1e18;

    long long solve(int i, int j, vector<int> &robot, vector<vector<int>> &factory)
    {
        int n = robot.size();
        int m = factory.size();

        // All robots are repaired
        if (i == n)
            return 0;

        // No factories left but robots still remain
        if (j == m)
            return INF;

        // Already computed
        if (dp[i][j] != -1)
            return dp[i][j];

        // Option 1: Skip current factory
        long long ans = solve(i, j + 1, robot, factory);

        long long distance = 0;
        int pos = factory[j][0];
        int limit = factory[j][1];

        // Option 2: Use current factory for next k robots
        for (int k = 0; k < limit && i + k < n; k++)
        {
            distance += abs(robot[i + k] - pos);

            long long next = solve(i + k + 1, j + 1, robot, factory);

            if (next != INF)
            {
                ans = min(ans, distance + next);
            }
        }

        return dp[i][j] = ans;
    }

    long long minimumTotalDistance(vector<int> &robot, vector<vector<int>> &factory)
    {
        sort(robot.begin(), robot.end());
        sort(factory.begin(), factory.end());

        memset(dp, -1, sizeof(dp));

        return solve(0, 0, robot, factory);
    }
};