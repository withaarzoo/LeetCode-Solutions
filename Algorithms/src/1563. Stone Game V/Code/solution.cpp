class Solution
{
public:
    int stoneGameV(vector<int> &stoneValue)
    {
        int n = stoneValue.size();

        // prefix[i] stores the sum of the first i stones.
        // It lets me calculate any interval sum in O(1).
        vector<long long> prefix(n + 1, 0);
        for (int i = 0; i < n; ++i)
        {
            prefix[i + 1] = prefix[i] + stoneValue[i];
        }

        // dp[l][r] stores Alice's maximum score for interval [l, r].
        vector<vector<int>> dp(n, vector<int>(n, 0));

        // leftBest[l][r] stores the best:
        // dp[l][k] + sum(l, k), for k in [l, r].
        vector<vector<int>> leftBest(n, vector<int>(n, 0));

        // rightBest[l][r] stores the best:
        // dp[k][r] + sum(k, r), for k in [l, r].
        vector<vector<int>> rightBest(n, vector<int>(n, 0));

        // leftPtr[l] is the last split whose left sum is <= half
        // of the current interval sum.
        vector<int> leftPtr(n);

        // rightPtr[l] is the first split whose left sum is >= half
        // of the current interval sum.
        vector<int> rightPtr(n);

        for (int i = 0; i < n; ++i)
        {
            // A single stone cannot be split, so its game score is 0.
            // For helper tables, the single stone itself is a valid side.
            leftBest[i][i] = stoneValue[i];
            rightBest[i][i] = stoneValue[i];

            // Initially there is no valid split before i.
            leftPtr[i] = i - 1;

            // The first possible split starts at i.
            rightPtr[i] = i;
        }

        // I process intervals from shorter to longer intervals.
        // This guarantees every smaller dp state is already available.
        for (int len = 2; len <= n; ++len)
        {
            for (int l = 0; l + len <= n; ++l)
            {
                int r = l + len - 1;

                // Calculate the total sum of the current interval.
                long long total = prefix[r + 1] - prefix[l];

                // Move leftPtr forward while the left side is still
                // smaller than or equal to the right side.
                while (leftPtr[l] + 1 <= r - 1)
                {
                    int k = leftPtr[l] + 1;
                    long long leftSum = prefix[k + 1] - prefix[l];

                    if (2 * leftSum > total)
                    {
                        break;
                    }

                    ++leftPtr[l];
                }

                // Move rightPtr forward until the left side becomes
                // greater than or equal to the right side.
                while (rightPtr[l] <= r - 1)
                {
                    int k = rightPtr[l];
                    long long leftSum = prefix[k + 1] - prefix[l];

                    if (2 * leftSum >= total)
                    {
                        break;
                    }

                    ++rightPtr[l];
                }

                int best = 0;

                // If this split has leftSum <= rightSum, Alice keeps
                // the left side and gets its sum plus its previous score.
                if (leftPtr[l] >= l)
                {
                    best = leftBest[l][leftPtr[l]];
                }

                // If this split has leftSum >= rightSum, Alice keeps
                // the right side and gets its sum plus its previous score.
                if (rightPtr[l] <= r - 1)
                {
                    best = max(best, rightBest[rightPtr[l] + 1][r]);
                }

                // Store the best score Alice can obtain for [l, r].
                dp[l][r] = best;

                // Add the current interval as a possible left-side interval
                // for a future larger interval.
                leftBest[l][r] = max(
                    leftBest[l][r - 1],
                    dp[l][r] + static_cast<int>(total));

                // Add the current interval as a possible right-side interval
                // for a future larger interval.
                rightBest[l][r] = max(
                    rightBest[l + 1][r],
                    dp[l][r] + static_cast<int>(total));
            }
        }

        // dp[0][n - 1] contains the answer for the complete array.
        return dp[0][n - 1];
    }
};