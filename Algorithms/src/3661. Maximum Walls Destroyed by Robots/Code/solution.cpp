class Solution
{
public:
    int maxWalls(vector<int> &robots, vector<int> &distance, vector<int> &walls)
    {
        int n = robots.size();

        vector<pair<int, int>> arr;
        for (int i = 0; i < n; i++)
        {
            arr.push_back({robots[i], distance[i]});
        }

        sort(arr.begin(), arr.end());
        sort(walls.begin(), walls.end());

        // Dummy robot to avoid checking bounds for next robot
        arr.push_back({(int)1e9, 0});

        auto countWalls = [&](int left, int right)
        {
            if (left > right)
                return 0;

            int r = upper_bound(walls.begin(), walls.end(), right) - walls.begin();
            int l = lower_bound(walls.begin(), walls.end(), left) - walls.begin();

            return r - l;
        };

        vector<vector<int>> dp(n, vector<int>(2, 0));

        // First robot shoots left
        dp[0][0] = countWalls(arr[0].first - arr[0].second, arr[0].first);

        // First robot shoots right
        int rightLimit = (n == 1)
                             ? arr[0].first + arr[0].second
                             : min(arr[0].first + arr[0].second, arr[1].first - 1);

        dp[0][1] = countWalls(arr[0].first, rightLimit);

        for (int i = 1; i < n; i++)
        {
            int pos = arr[i].first;
            int dist = arr[i].second;

            // Current robot shoots right
            int reachRight = min(pos + dist, arr[i + 1].first - 1);
            int rightWalls = countWalls(pos, reachRight);

            dp[i][1] = max(dp[i - 1][0], dp[i - 1][1]) + rightWalls;

            // Current robot shoots left
            int leftStart = max(pos - dist, arr[i - 1].first + 1);
            int leftWalls = countWalls(leftStart, pos);

            // Previous robot also shot left
            dp[i][0] = dp[i - 1][0] + leftWalls;

            // Previous robot shot right -> overlap may happen
            int prevRightEnd = min(arr[i - 1].first + arr[i - 1].second, pos - 1);

            int overlapLeft = leftStart;
            int overlapRight = min(prevRightEnd, pos - 1);

            int overlapWalls = countWalls(overlapLeft, overlapRight);

            dp[i][0] = max(dp[i][0], dp[i - 1][1] + leftWalls - overlapWalls);
        }

        return max(dp[n - 1][0], dp[n - 1][1]);
    }
};