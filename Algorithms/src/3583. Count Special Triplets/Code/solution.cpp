#include <bits/stdc++.h>
using namespace std;

class Solution
{
public:
    int specialTriplets(vector<int> &nums)
    {
        const int MOD = 1'000'000'007;
        unordered_map<int, long long> right, left;

        // Count all elements into 'right'
        for (int x : nums)
        {
            right[x]++;
        }

        long long ans = 0;

        for (int x : nums)
        {
            // This x is now the middle element, so remove one occurrence from right
            right[x]--;

            long long target = (long long)x * 2; // value 2x

            long long cntLeft = left.count(target) ? left[target] : 0;
            long long cntRight = right.count(target) ? right[target] : 0;

            // Add all combinations of (i,k) with this middle index
            long long add = (cntLeft * cntRight) % MOD;
            ans = (ans + add) % MOD;

            // Move current x to the left map
            left[x]++;
        }

        return (int)(ans % MOD);
    }
};
