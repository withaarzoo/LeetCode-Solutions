class Solution
{
public:
    // Each DP state stores the maximum score and the sorted original indices.
    struct Node
    {
        long long score; // Total weight of the selected intervals.
        vector<int> ids; // Original indices, kept in sorted order.
        bool valid;      // Tells me whether this state can be formed.

        Node() : score(0), valid(false) {}
        Node(long long s, vector<int> v) : score(s), ids(std::move(v)), valid(true) {}
    };

    // Returns true when 'a' is better than 'b'.
    bool better(const Node &a, const Node &b)
    {
        // A valid candidate is always better than an invalid state.
        if (!a.valid)
            return false;
        if (!b.valid)
            return true;

        // A larger score is the primary requirement.
        if (a.score != b.score)
            return a.score > b.score;

        // If scores tie, I need the lexicographically smaller index array.
        return a.ids < b.ids;
    }

    vector<int> maximumWeight(vector<vector<int>> &intervals)
    {
        int n = intervals.size(); // Number of intervals.
        const int K = 4;          // I can choose at most four intervals.

        // I store [left, right, weight, originalIndex].
        vector<array<long long, 4>> a(n);

        // Attach each interval's original index before sorting.
        for (int i = 0; i < n; ++i)
        {
            a[i] = {intervals[i][0], intervals[i][1], intervals[i][2], i};
        }

        // Sorting by right endpoint makes compatible intervals form a prefix.
        sort(a.begin(), a.end(), [](const auto &x, const auto &y)
             { return x[1] < y[1]; });

        // Store all right endpoints so binary search becomes easy.
        vector<long long> ends(n);
        for (int i = 0; i < n; ++i)
        {
            ends[i] = a[i][1];
        }

        // dp[k][i] = best result using exactly k intervals
        // from the first i sorted intervals.
        vector<vector<Node>> dp(K + 1, vector<Node>(n + 1));

        // Choosing zero intervals always gives score 0 and an empty index list.
        for (int i = 0; i <= n; ++i)
        {
            dp[0][i] = Node(0, {});
        }

        // Process the sorted intervals one by one.
        for (int i = 1; i <= n; ++i)
        {
            long long l = a[i - 1][0];  // Current interval's left endpoint.
            long long w = a[i - 1][2];  // Current interval's weight.
            int idx = (int)a[i - 1][3]; // Current interval's original index.

            // Find the first ending point >= l.
            // Every interval before it has end < l and is compatible.
            int p = lower_bound(ends.begin(), ends.begin() + (i - 1), l) - ends.begin();

            // For every possible number of selected intervals...
            for (int k = 1; k <= K; ++k)
            {
                // Option 1: skip the current interval.
                dp[k][i] = dp[k][i - 1];

                // Option 2 is possible only when dp[k - 1][p] exists.
                if (dp[k - 1][p].valid)
                {
                    // Copy the previous selected indices.
                    vector<int> ids = dp[k - 1][p].ids;

                    // Add the current original index.
                    ids.push_back(idx);

                    // Keep indices sorted because the final answer
                    // has to be compared lexicographically.
                    sort(ids.begin(), ids.end());

                    // Build the candidate obtained by taking this interval.
                    Node take(dp[k - 1][p].score + w, std::move(ids));

                    // Keep whichever of skip/take is better.
                    if (better(take, dp[k][i]))
                    {
                        dp[k][i] = std::move(take);
                    }
                }
            }
        }

        // The answer can contain 1, 2, 3, or 4 intervals.
        Node ans;

        // Compare all valid final states.
        for (int k = 1; k <= K; ++k)
        {
            if (better(dp[k][n], ans))
            {
                ans = dp[k][n];
            }
        }

        // Return the original indices of the best set.
        return ans.ids;
    }
};