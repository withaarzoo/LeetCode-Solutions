#include <vector>
#include <queue>
using namespace std;

class Solution {
public:
    vector<vector<int>> pacificAtlantic(vector<vector<int>>& heights) {
        if (heights.empty() || heights[0].empty()) return {};
        int m = heights.size(), n = heights[0].size();

        vector<vector<bool>> pac(m, vector<bool>(n, false));
        vector<vector<bool>> atl(m, vector<bool>(n, false));

        queue<pair<int,int>> q;

        // Start BFS from Pacific borders: top row and left column
        for (int j = 0; j < n; ++j) {
            q.push({0, j});
            pac[0][j] = true;
        }
        for (int i = 1; i < m; ++i) { // start at 1 to avoid pushing (0,0) twice
            q.push({i, 0});
            pac[i][0] = true;
        }
        bfs(heights, q, pac);

        // Start BFS from Atlantic borders: bottom row and right column
        queue<pair<int,int>> q2;
        for (int j = 0; j < n; ++j) {
            q2.push({m-1, j});
            atl[m-1][j] = true;
        }
        for (int i = 0; i < m-1; ++i) { // avoid pushing (m-1, n-1) twice
            q2.push({i, n-1});
            atl[i][n-1] = true;
        }
        bfs(heights, q2, atl);

        // Collect intersection
        vector<vector<int>> res;
        for (int i = 0; i < m; ++i)
            for (int j = 0; j < n; ++j)
                if (pac[i][j] && atl[i][j])
                    res.push_back({i, j});
        return res;
    }

private:
    void bfs(const vector<vector<int>>& heights, queue<pair<int,int>>& q, vector<vector<bool>>& visited) {
        int m = heights.size(), n = heights[0].size();
        const int dirs[4][2] = {{1,0},{-1,0},{0,1},{0,-1}};
        while (!q.empty()) {
            auto cur = q.front(); q.pop();
            int r = cur.first, c = cur.second;
            for (int k = 0; k < 4; ++k) {
                int nr = r + dirs[k][0], nc = c + dirs[k][1];
                if (nr < 0 || nr >= m || nc < 0 || nc >= n) continue;
                if (visited[nr][nc]) continue;                         // already known reachable
                if (heights[nr][nc] < heights[r][c]) continue;         // can't move uphill->downhill reversed rule
                visited[nr][nc] = true;
                q.push({nr, nc});
            }
        }
    }
};
