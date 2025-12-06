#include <bits/stdc++.h>
using namespace std;

class Solution
{
public:
    int countPartitions(vector<int> &nums, int k)
    {
        const int MOD = 1'000'000'007;
        int n = (int)nums.size();

        // dp[i] = ways to partition prefix of length i (nums[0..i-1])
        vector<long long> dp(n + 1, 0), pref(n + 1, 0);
        dp[0] = 1;
        pref[0] = 1;

        deque<int> maxdq, mindq; // store indices
        int l = 0;

        for (int r = 0; r < n; ++r)
        {
            int x = nums[r];

            // maintain decreasing deque for max
            while (!maxdq.empty() && nums[maxdq.back()] <= x)
            {
                maxdq.pop_back();
            }
            maxdq.push_back(r);

            // maintain increasing deque for min
            while (!mindq.empty() && nums[mindq.back()] >= x)
            {
                mindq.pop_back();
            }
            mindq.push_back(r);

            // shrink left until window is valid
            while (!maxdq.empty() && !mindq.empty() &&
                   (long long)nums[maxdq.front()] - nums[mindq.front()] > k)
            {
                if (maxdq.front() == l)
                    maxdq.pop_front();
                if (mindq.front() == l)
                    mindq.pop_front();
                ++l;
            }

            int L = l;     // minimal valid start index for segment ending at r
            int i = r + 1; // prefix length

            long long ways = pref[i - 1]; // sum dp[0..i-1]
            if (L > 0)
            {
                ways -= pref[L - 1]; // remove dp[0..L-1]
            }
            ways %= MOD;
            if (ways < 0)
                ways += MOD;

            dp[i] = ways;
            pref[i] = (pref[i - 1] + dp[i]) % MOD;
        }

        return (int)dp[n];
    }
};
