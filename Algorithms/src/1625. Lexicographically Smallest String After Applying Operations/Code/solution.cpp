#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    string findLexSmallestString(string s, int a, int b) {
        int n = s.size();
        unordered_set<string> seen;
        queue<string> q;
        seen.insert(s);
        q.push(s);
        string ans = s;

        while (!q.empty()) {
            string cur = q.front(); q.pop();
            if (cur < ans) ans = cur;

            // Operation 1: add a to all odd indices (0-indexed => positions 1,3,5,...)
            string addOp = cur;
            for (int i = 1; i < n; i += 2) {
                int d = (addOp[i] - '0' + a) % 10;
                addOp[i] = char('0' + d);
            }
            if (!seen.count(addOp)) {
                seen.insert(addOp);
                q.push(addOp);
            }

            // Operation 2: rotate right by b -> equivalent to rotate left by n-b
            string rotOp = cur.substr(n - b) + cur.substr(0, n - b);
            if (!seen.count(rotOp)) {
                seen.insert(rotOp);
                q.push(rotOp);
            }
        }
        return ans;
    }
};
