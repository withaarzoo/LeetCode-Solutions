class Solution
{
public:
    long long maxProfit(vector<int> &prices, vector<int> &strategy, int k)
    {
        int n = prices.size();

        // Step 1: Base profit
        long long base = 0;
        for (int i = 0; i < n; i++)
        {
            base += 1LL * strategy[i] * prices[i];
        }

        // Step 2: Prefix sums
        vector<long long> prefPrice(n + 1, 0), prefProfit(n + 1, 0);
        for (int i = 0; i < n; i++)
        {
            prefPrice[i + 1] = prefPrice[i] + prices[i];
            prefProfit[i + 1] = prefProfit[i] + 1LL * strategy[i] * prices[i];
        }

        long long bestDelta = 0;
        int half = k / 2;

        // Step 3: Sliding window
        for (int l = 0; l + k <= n; l++)
        {
            int m = l + half;
            int r = l + k;

            long long oldProfit = prefProfit[r] - prefProfit[l];
            long long newProfit = prefPrice[r] - prefPrice[m]; // only sells

            bestDelta = max(bestDelta, newProfit - oldProfit);
        }

        return base + bestDelta;
    }
};
