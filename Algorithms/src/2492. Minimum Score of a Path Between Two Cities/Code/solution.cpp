class Solution {
public:
    int minScore(int n, vector<vector<int>>& roads) {
        // graph[city] stores {neighbor, road distance}.
        vector<vector<pair<int, int>>> graph(n + 1);

        // Build the adjacency list.
        // Each road is stored in both directions because the graph is undirected.
        for (const auto& road : roads) {
            int a = road[0];
            int b = road[1];
            int distance = road[2];

            graph[a].push_back({b, distance});
            graph[b].push_back({a, distance});
        }

        // visited prevents processing the same city again and again.
        vector<bool> visited(n + 1, false);

        // I use a queue to perform BFS from city 1.
        queue<int> q;
        q.push(1);
        visited[1] = true;

        // Start with the largest possible integer value.
        int answer = INT_MAX;

        // Visit every city connected to city 1.
        while (!q.empty()) {
            int city = q.front();
            q.pop();

            // Check every road leaving the current city.
            for (const auto& edge : graph[city]) {
                int nextCity = edge.first;
                int distance = edge.second;

                // Every edge in this component can be part of a valid path.
                answer = min(answer, distance);

                // Add an unvisited city so its roads are also checked.
                if (!visited[nextCity]) {
                    visited[nextCity] = true;
                    q.push(nextCity);
                }
            }
        }

        // This is the smallest road distance in city 1's component.
        return answer;
    }
};