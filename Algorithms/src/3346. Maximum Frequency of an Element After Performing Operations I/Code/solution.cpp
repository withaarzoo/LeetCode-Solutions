#include <bits/stdc++.h>
using namespace std;

class Solution
{
public:
    int maxFrequency(vector<int> &nums, int k, int numOperations)
    {
        if (nums.empty())
            return 0;
        int mx = *max_element(nums.begin(), nums.end());
        int size = mx + k + 2; // safe buffer for R = t + k
        vector<int> count(size, 0);

        // fill frequency
        for (int v : nums)
            count[v]++;

        // prefix sums
        for (int i = 1; i < size; ++i)
            count[i] += count[i - 1];

        int ans = 0;
        for (int t = 0; t < size; ++t)
        {
            int L = max(0, t - k);
            int R = min(size - 1, t + k);
            int total = count[R] - (L > 0 ? count[L - 1] : 0);         // how many can be converted to t
            int freq_t = (t > 0 ? count[t] - count[t - 1] : count[t]); // how many are already t
            int canConvert = total - freq_t;                           // non-equal available
            int take = min(numOperations, canConvert);
            ans = max(ans, freq_t + take);
        }
        return ans;
    }
};
