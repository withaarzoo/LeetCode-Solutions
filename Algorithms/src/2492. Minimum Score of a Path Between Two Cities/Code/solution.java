class Solution {
    public int minScore(int n, int[][] roads) {
        // graph[city] stores arrays of {neighbor, road distance}.
        List<int[]>[] graph = new ArrayList[n + 1];

        // Create an empty adjacency list for every city.
        for (int city = 1; city <= n; city++) {
            graph[city] = new ArrayList<>();
        }

        // Build the undirected graph by storing every road both ways.
        for (int[] road : roads) {
            int a = road[0];
            int b = road[1];
            int distance = road[2];

            graph[a].add(new int[] { b, distance });
            graph[b].add(new int[] { a, distance });
        }

        // visited prevents the BFS from processing one city repeatedly.
        boolean[] visited = new boolean[n + 1];

        // I use a queue to explore the whole component of city 1.
        Queue<Integer> queue = new ArrayDeque<>();
        queue.offer(1);
        visited[1] = true;

        // Start with the largest possible integer value.
        int answer = Integer.MAX_VALUE;

        // Continue until every reachable city has been processed.
        while (!queue.isEmpty()) {
            int city = queue.poll();

            // Check every road connected to the current city.
            for (int[] edge : graph[city]) {
                int nextCity = edge[0];
                int distance = edge[1];

                // Keep the smallest road found anywhere in this component.
                answer = Math.min(answer, distance);

                // Visit the neighboring city only if it is still unvisited.
                if (!visited[nextCity]) {
                    visited[nextCity] = true;
                    queue.offer(nextCity);
                }
            }
        }

        // Return the minimum road distance in city 1's component.
        return answer;
    }
}