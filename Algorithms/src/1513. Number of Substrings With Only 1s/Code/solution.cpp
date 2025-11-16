#include <bits/stdc++.h>
using namespace std;

class Solution
{
public:
    int numSub(string s)
    {
        const long long MOD = 1000000007LL;
        long long res = 0;
        long long cnt = 0; // current consecutive '1's count

        for (char c : s)
        {
            if (c == '1')
            {
                cnt++;
            }
            else
            {
                // block ended, add contribution
                res = (res + (cnt * (cnt + 1) / 2) % MOD) % MOD;
                cnt = 0;
            }
        }
        // add last block if string ended with '1'
        res = (res + (cnt * (cnt + 1) / 2) % MOD) % MOD;
        return (int)res;
    }
};
