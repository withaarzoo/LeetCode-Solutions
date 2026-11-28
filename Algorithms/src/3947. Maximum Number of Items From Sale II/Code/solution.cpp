class Solution {
public:
    int maximumSaleItems(vector<vector<int>>& items, int budget) {
        int n = items.size();
        vector<int> cnt(n + 1, 0);
        for (auto& it : items) cnt[it[0]]++;
        vector<int> mult(n + 1, 0);
        for (int f = 1; f <= n; f++) {
            for (int m = f; m <= n; m += f) mult[f] += cnt[m];
        }
        int minP = INT_MAX;
        vector<pair<int, int>> cand;
        for (auto& it : items) {
            int f = it[0], p = it[1];
            minP = min(minP, p);
            int g = mult[f] - 1;
            if (g > 0) cand.emplace_back(p, g);
        }
        sort(cand.begin(), cand.end());
        long long spent = 0, total = 0, B = budget;
        for (auto& [p, g] : cand) {
            if (p > 2LL * minP) break;
            long long can = min((long long)g, (B - spent) / p);
            total += can * 2;
            spent += can * p;
        }
        total += (B - spent) / minP;
        return (int)total;
    }
};