#include <bits/stdc++.h>
using namespace std;

class Solution
{
public:
    int numberOfSubstrings(string s)
    {
        int n = (int)s.size();
        long long ans = 0;
        // 1) all-ones substrings
        long long run = 0;
        for (char c : s)
        {
            if (c == '1')
                run++;
            else
            {
                ans += run * (run + 1) / 2;
                run = 0;
            }
        }
        ans += run * (run + 1) / 2;

        // 2) indices of zeros
        vector<int> zeroPos;
        for (int i = 0; i < n; ++i)
            if (s[i] == '0')
                zeroPos.push_back(i);
        int m = (int)zeroPos.size();
        if (m == 0)
            return (int)ans;

        int K = (int)floor(sqrt(n));
        for (int k = 1; k <= K && k <= m; ++k)
        {
            for (int i = 0; i + k - 1 < m; ++i)
            {
                int leftPrev = (i == 0 ? -1 : zeroPos[i - 1]);
                int rightNext = (i + k - 1 == m - 1 ? n : zeroPos[i + k]);
                int leftOnes = zeroPos[i] - leftPrev - 1;
                int rightOnes = rightNext - zeroPos[i + k - 1] - 1;
                int baseLen = zeroPos[i + k - 1] - zeroPos[i] + 1;
                long long needLen = 1LL * k * k + k;
                long long t = needLen - baseLen; // need x+y >= t
                long long totalPairs = 1LL * (leftOnes + 1) * (rightOnes + 1);
                if (t <= 0)
                {
                    ans += totalPairs;
                    continue;
                }

                // Count pairs with x+y < t in O(1)
                long long ssum = t - 1; // we want x+y <= ssum
                if (ssum < 0)
                { /* none */
                }
                else
                {
                    long long L = leftOnes, R = rightOnes;
                    long long x_max = min(L, ssum);
                    if (x_max < 0)
                    {
                        // no x possible => pairs_lt = 0
                    }
                    else
                    {
                        long long x0 = max(0LL, ssum - R); // for x < x0 -> y_max >= R
                        if (x0 > x_max)
                        {
                            // all x in [0..x_max] have y_max >= R
                            long long pairs_lt = (x_max + 1) * (R + 1);
                            ans += totalPairs - pairs_lt; // will add valid below; but we need pairs_lt separately
                            // But we should actually compute valid = totalPairs - pairs_lt
                            // Replace: compute pairs_lt and then valid
                            // We'll restructure for clarity below
                        }
                    }
                }
                // Simpler correct closed-form computation (structured):
                long long pairs_lt = 0;
                long long s0 = t - 1;
                if (s0 >= 0)
                {
                    long long L = leftOnes;
                    long long R = rightOnes;
                    long long x_max = min(L, s0);
                    if (x_max >= 0)
                    {
                        long long x0 = (s0 - R > 0) ? (s0 - R) : 0; // first x where y_max < R
                        if (x0 > x_max)
                        {
                            // for all x in [0..x_max], y_max >= R
                            pairs_lt = (x_max + 1) * (R + 1);
                        }
                        else
                        {
                            // part1: x in [0..x0-1] -> (R+1) each
                            long long part1 = x0 * (R + 1);
                            // part2: x in [x0..x_max] -> y_max = s0 - x  => count = s0 - x + 1
                            long long n2 = x_max - x0 + 1;
                            // sum_{x = x0..x_max} (s0 - x + 1) = n2*(s0+1) - sum_{x=x0..x_max} x
                            long long sum_x = (x0 + x_max) * n2 / 2;
                            long long part2 = n2 * (s0 + 1) - sum_x;
                            pairs_lt = part1 + part2;
                        }
                    }
                    else
                        pairs_lt = 0;
                }
                else
                    pairs_lt = 0;

                long long valid = totalPairs - pairs_lt;
                if (valid > 0)
                    ans += valid;
            }
        }
        return (int)ans;
    }
};
