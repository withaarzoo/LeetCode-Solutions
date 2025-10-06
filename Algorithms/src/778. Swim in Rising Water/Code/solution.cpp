#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int swimInWater(vector<vector<int>>& grid) {
        int n = grid.size();
        vector<vector<bool>> vis(n, vector<bool>(n, false));
        // min-heap of tuples (time, r, c)
        using T = tuple<int,int,int>;
        priority_queue<T, vector<T>, greater<T>> pq;
        pq.emplace(grid[0][0], 0, 0);

        int dirs[4][2] = {{1,0},{-1,0},{0,1},{0,-1}};
        while (!pq.empty()) {
            auto [t, r, c] = pq.top(); pq.pop();
            if (vis[r][c]) continue;
            vis[r][c] = true;
            // if we reached bottom-right, t is the minimum required time
            if (r == n-1 && c == n-1) return t;
            for (auto &d : dirs) {
                int nr = r + d[0], nc = c + d[1];
                if (nr >= 0 && nr < n && nc >= 0 && nc < n && !vis[nr][nc]) {
                    int nt = max(t, grid[nr][nc]); // time needed to step into neighbor
                    pq.emplace(nt, nr, nc);
                }
            }
        }
        return -1; // unreachable (shouldn't happen under given constraints)
    }
};
