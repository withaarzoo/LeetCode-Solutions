#include <bits/stdc++.h>
using namespace std;

class Solution
{
public:
    int minCost(string colors, vector<int> &neededTime)
    {
        long long ans = 0;       // result (use long long to be safe)
        long long block_sum = 0; // sum of times in current same-color block
        int block_max = 0;       // maximum time in current block
        int n = colors.size();

        for (int i = 0; i < n; ++i)
        {
            if (i > 0 && colors[i] != colors[i - 1])
            {
                // end of previous block, add cost to remove all but the max
                ans += block_sum - block_max;
                block_sum = 0;
                block_max = 0;
            }
            block_sum += neededTime[i];
            block_max = max(block_max, neededTime[i]);
        }
        // handle last block
        ans += block_sum - block_max;
        return (int)ans;
    }
};
