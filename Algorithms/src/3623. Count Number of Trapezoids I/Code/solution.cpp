#include <bits/stdc++.h>
using namespace std;

class Solution
{
public:
    int countTrapezoids(vector<vector<int>> &points)
    {
        const long long MOD = 1000000007LL;
        const long long INV2 = (MOD + 1) / 2; // modular inverse of 2

        // 1. Count points per y-coordinate
        unordered_map<long long, int> freq;
        freq.reserve(points.size() * 2);
        for (auto &p : points)
        {
            long long y = p[1];
            ++freq[y];
        }

        long long sumF = 0;  // S = sum of f(y)
        long long sumF2 = 0; // SQ = sum of f(y)^2

        // 2. For each y, compute number of pairs C(c,2),
        //    then update S and SQ
        for (auto &kv : freq)
        {
            long long c = kv.second;
            if (c >= 2)
            {
                long long f = c * (c - 1) / 2 % MOD; // C(c,2) mod MOD
                sumF = (sumF + f) % MOD;
                sumF2 = (sumF2 + f * f % MOD) % MOD;
            }
        }

        // 3. Compute ((S^2 - SQ) / 2) mod MOD
        long long ans = (sumF * sumF) % MOD;
        ans = (ans - sumF2 + MOD) % MOD; // ensure non-negative
        ans = ans * INV2 % MOD;

        return (int)ans;
    }
};
