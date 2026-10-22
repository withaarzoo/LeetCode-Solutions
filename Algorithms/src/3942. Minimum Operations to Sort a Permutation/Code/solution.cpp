class Solution {
public:
    int minOperations(vector<int>& nums) {
        int n = nums.size();
        vector<pair<int,int>> starts;
        bool inc = true, dec = true;
        for (int i = 0; i < n; ++i) {
            if (nums[(i + 1) % n] != (nums[i] + 1) % n) inc = false;
            if (nums[(i + 1) % n] != (nums[i] - 1 + n) % n) dec = false;
        }
        if (inc) starts.emplace_back(0, nums[0]);
        if (dec) starts.emplace_back(1, (n - 1 - nums[0] + n) % n);
        if (starts.empty()) return -1;

        vector<vector<int>> dist(2, vector<int>(n, -1));
        queue<pair<int,int>> q;
        for (auto [t, s] : starts) {
            dist[t][s] = 0;
            q.emplace(t, s);
        }
        while (!q.empty()) {
            auto [t, s] = q.front(); q.pop();
            if (t == 0 && s == 0) return dist[t][s];
            int d = dist[t][s];
            int nt = t, ns = (s + 1) % n;
            if (dist[nt][ns] == -1) {
                dist[nt][ns] = d + 1;
                q.emplace(nt, ns);
            }
            nt = 1 - t;
            ns = (n - s) % n;
            if (dist[nt][ns] == -1) {
                dist[nt][ns] = d + 1;
                q.emplace(nt, ns);
            }
        }
        return -1;
    }
};