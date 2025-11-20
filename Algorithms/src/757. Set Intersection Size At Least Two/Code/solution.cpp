#include <bits/stdc++.h>
using namespace std;

class Solution
{
public:
    int intersectionSizeTwo(vector<vector<int>> &intervals)
    {
        // sort by end ascending, if tie then start descending
        sort(intervals.begin(), intervals.end(), [](const vector<int> &A, const vector<int> &B)
             {
            if (A[1] != B[1]) return A[1] < B[1];
            return A[0] > B[0]; });

        int a = -1e9, b = -1e9; // last two chosen integers (a < b)
        int ans = 0;
        for (auto &iv : intervals)
        {
            int l = iv[0], r = iv[1];
            if (l > b)
            {
                // none of a,b in [l,r]; pick r-1 and r
                ans += 2;
                a = r - 1;
                b = r;
            }
            else if (l > a)
            {
                // only b is in [l,r]; pick r
                ans += 1;
                a = b;
                b = r;
            }
            else
            {
                // both a and b already in [l,r]; do nothing
            }
        }
        return ans;
    }
};
