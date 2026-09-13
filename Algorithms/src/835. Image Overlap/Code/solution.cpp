class Solution {
public:
    int largestOverlap(vector<vector<int>>& img1, vector<vector<int>>& img2) {
        int n = img1.size();
        // collect every coordinate that holds a 1
        vector<pair<int,int>> A, B;
        for (int i = 0; i < n; ++i) {
            for (int j = 0; j < n; ++j) {
                if (img1[i][j] == 1) A.emplace_back(i, j);
                if (img2[i][j] == 1) B.emplace_back(i, j);
            }
        }
        // count how many times each possible shift appears
        // shifts range from -(n-1) to +(n-1), so offset by n
        vector<vector<int>> cnt(2 * n, vector<int>(2 * n, 0));
        int best = 0;
        for (auto& a : A) {
            for (auto& b : B) {
                int dx = b.first - a.first + n;   // make index non-negative
                int dy = b.second - a.second + n;
                best = max(best, ++cnt[dx][dy]);
            }
        }
        return best;
    }
};