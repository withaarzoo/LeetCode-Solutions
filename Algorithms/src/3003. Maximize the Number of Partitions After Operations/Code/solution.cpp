#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    string s;
    int k;
    unordered_map<long long, int> memo;

    int dp(int i, long long mask, bool canChange) {
        if (i == s.size()) return 0;
        long long key = ((long long)i << 30) | (mask << 1) | canChange;
        if (memo.count(key)) return memo[key];

        int bit = s[i] - 'a';
        long long newMask = mask | (1LL << bit);
        int res = 0;

        // If current partition exceeds k distinct chars, we start new
        if (__builtin_popcountll(newMask) > k)
            res = 1 + dp(i + 1, 1LL << bit, canChange);
        else
            res = dp(i + 1, newMask, canChange);

        // Try changing this character (if still available)
        if (canChange) {
            for (int j = 0; j < 26; ++j) {
                long long changeMask = mask | (1LL << j);
                if (__builtin_popcountll(changeMask) > k)
                    res = max(res, 1 + dp(i + 1, 1LL << j, false));
                else
                    res = max(res, dp(i + 1, changeMask, false));
            }
        }

        return memo[key] = res;
    }

    int maxPartitionsAfterOperations(string s, int k) {
        this->s = s;
        this->k = k;
        memo.clear();
        return dp(0, 0, true) + 1;
    }
};
