class Solution
{
public:
    int maxBuilding(int n, vector<vector<int>> &restrictions)
    {
        // Building 1 must have height 0
        restrictions.push_back({1, 0});

        // Building n can never exceed n - 1
        restrictions.push_back({n, n - 1});

        // Sort restrictions by building index
        sort(restrictions.begin(), restrictions.end());

        int m = restrictions.size();

        // Left to right pass
        // Make sure every restriction is reachable from the left
        for (int i = 1; i < m; i++)
        {
            int dist = restrictions[i][0] - restrictions[i - 1][0];

            restrictions[i][1] = min(
                restrictions[i][1],
                restrictions[i - 1][1] + dist);
        }

        // Right to left pass
        // Make sure every restriction is reachable from the right
        for (int i = m - 2; i >= 0; i--)
        {
            int dist = restrictions[i + 1][0] - restrictions[i][0];

            restrictions[i][1] = min(
                restrictions[i][1],
                restrictions[i + 1][1] + dist);
        }

        long long ans = 0;

        // Compute highest peak inside every interval
        for (int i = 1; i < m; i++)
        {
            long long x1 = restrictions[i - 1][0];
            long long h1 = restrictions[i - 1][1];

            long long x2 = restrictions[i][0];
            long long h2 = restrictions[i][1];

            long long dist = x2 - x1;

            // Highest achievable height in this segment
            long long peak =
                max(h1, h2) +
                (dist - llabs(h1 - h2)) / 2;

            ans = max(ans, peak);
        }

        return (int)ans;
    }
};