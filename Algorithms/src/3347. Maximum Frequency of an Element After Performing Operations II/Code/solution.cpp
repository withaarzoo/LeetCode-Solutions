#include <bits/stdc++.h>
using namespace std;

class Solution
{
public:
    int maxFrequency(vector<int> &nums, int k, int numOperations)
    {
        int n = nums.size();
        if (n == 0)
            return 0;
        sort(nums.begin(), nums.end());

        // frequency map for existing values
        unordered_map<long long, int> freq;
        freq.reserve(n * 2);
        for (int x : nums)
            freq[x]++;

        int ans = 1;

        // Case A: existing values as target
        // For each distinct value v, count how many nums are in [v-k, v+k]
        for (auto &p : freq)
        {
            long long v = p.first;
            int already = p.second;
            // find number of elements in [v-k, v+k]
            long long lowVal = v - k;
            long long highVal = v + k;
            auto L = lower_bound(nums.begin(), nums.end(), (int)lowVal);
            auto R = upper_bound(nums.begin(), nums.end(), (int)highVal);
            int totalInRange = int(R - L);
            int need = totalInRange - already; // those that require operations
            int canFix = min(need, numOperations);
            ans = max(ans, already + canFix);
        }

        // Case B: non-existing target (elements inside a 2*k window can meet)
        // two pointers to find longest subarray with nums[j] - nums[i] <= 2*k
        int l = 0;
        for (int r = 0; r < n; ++r)
        {
            while (l <= r && (long long)nums[r] - nums[l] > 2LL * k)
                ++l;
            int w = r - l + 1;
            // we can make at most numOperations elements equal (if none already equal),
            // but window size gives how many can meet to some common new value
            ans = max(ans, min(w, numOperations));
            // Note: This line ensures cases like [5,64], k=42, numOperations=2 -> window w=2 -> ans=2
            // Also combination with Case A covers windows that already have duplicates.
        }

        return ans;
    }
};
