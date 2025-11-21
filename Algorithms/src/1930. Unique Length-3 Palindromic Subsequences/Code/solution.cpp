#include <bits/stdc++.h>
using namespace std;

class Solution
{
public:
    int countPalindromicSubsequence(string s)
    {
        int n = s.size();
        const int A = 26;
        vector<int> first(A, INT_MAX), last(A, -1);
        // record first and last occurrence for every letter
        for (int i = 0; i < n; ++i)
        {
            int c = s[i] - 'a';
            first[c] = min(first[c], i);
            last[c] = max(last[c], i);
        }

        int ans = 0;
        // for each outer letter, count distinct middle letters between first and last
        for (int c = 0; c < A; ++c)
        {
            if (first[c] < last[c])
            {
                vector<bool> seen(A, false);
                for (int i = first[c] + 1; i < last[c]; ++i)
                {
                    seen[s[i] - 'a'] = true;
                }
                for (int j = 0; j < A; ++j)
                    if (seen[j])
                        ++ans;
            }
        }
        return ans;
    }
};
