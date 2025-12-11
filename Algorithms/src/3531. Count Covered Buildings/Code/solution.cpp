#include <bits/stdc++.h>
using namespace std;

class Solution
{
public:
    int countCoveredBuildings(int n, vector<vector<int>> &buildings)
    {
        int m = buildings.size();
        unordered_map<int, vector<int>> row; // x -> list of y
        unordered_map<int, vector<int>> col; // y -> list of x
        row.reserve(m * 2);
        col.reserve(m * 2);

        // Build maps
        for (auto &b : buildings)
        {
            int x = b[0], y = b[1];
            row[x].push_back(y);
            col[y].push_back(x);
        }

        // Sort each row and column
        for (auto &kv : row)
        {
            auto &v = kv.second;
            sort(v.begin(), v.end());
        }
        for (auto &kv : col)
        {
            auto &v = kv.second;
            sort(v.begin(), v.end());
        }

        // For quick membership and to check positions, use unordered_map of unordered_set?
        // Instead, we will check by binary search for position in sorted vectors.
        int ans = 0;
        for (auto &b : buildings)
        {
            int x = b[0], y = b[1];
            auto &r = row[x]; // sorted y's in this row
            auto &c = col[y]; // sorted x's in this column

            // building must not be first or last in row (so has both left and right)
            bool insideRow = false, insideCol = false;
            // find position of y in r
            auto itR = lower_bound(r.begin(), r.end(), y);
            if (itR != r.begin() && (itR + 1) != r.end())
                insideRow = true;
            // find position of x in c
            auto itC = lower_bound(c.begin(), c.end(), x);
            if (itC != c.begin() && (itC + 1) != c.end())
                insideCol = true;

            if (insideRow && insideCol)
                ++ans;
        }
        return ans;
    }
};
