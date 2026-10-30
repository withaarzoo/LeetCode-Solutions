class Solution {
public:
    int shortestPath(int n, vector<vector<int>>& edges, string labels, int k) {
        vector<vector<pair<int,int>>> g(n);
        for (auto& e : edges) {
            g[e[0]].emplace_back(e[1], e[2]);
        }
        const int INF = 1e9;
        vector<vector<int>> dist(n, vector<int>(k + 1, INF));
        priority_queue<tuple<int,int,int>, vector<tuple<int,int,int>>, greater<>> pq;
        dist[0][1] = 0;
        pq.emplace(0, 0, 1);
        while (!pq.empty()) {
            auto [d, u, cnt] = pq.top();
            pq.pop();
            if (d > dist[u][cnt]) continue;
            if (u == n - 1) return d;
            for (auto [v, w] : g[u]) {
                int ncnt = (labels[u] == labels[v]) ? cnt + 1 : 1;
                if (ncnt > k) continue;
                if (d + w < dist[v][ncnt]) {
                    dist[v][ncnt] = d + w;
                    pq.emplace(dist[v][ncnt], v, ncnt);
                }
            }
        }
        return -1;
    }
};