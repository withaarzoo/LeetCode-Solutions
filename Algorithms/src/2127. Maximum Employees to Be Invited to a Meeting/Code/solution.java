class Solution {
    public int maximumInvitations(int[] favorite) {
        int n = favorite.length;
        int[] inDegree = new int[n];
        int[] chainLengths = new int[n];
        boolean[] visited = new boolean[n];

        for (int fav : favorite) {
            inDegree[fav]++;
        }

        Queue<Integer> queue = new LinkedList<>();
        for (int i = 0; i < n; ++i) {
            if (inDegree[i] == 0) {
                queue.offer(i);
            }
        }

        while (!queue.isEmpty()) {
            int node = queue.poll();
            visited[node] = true;

            int next = favorite[node];
            chainLengths[next] = chainLengths[node] + 1;
            if (--inDegree[next] == 0) {
                queue.offer(next);
            }
        }

        int maxCycle = 0, totalChains = 0;
        for (int i = 0; i < n; ++i) {
            if (!visited[i]) {
                int current = i, cycleLength = 0;
                while (!visited[current]) {
                    visited[current] = true;
                    current = favorite[current];
                    cycleLength++;
                }

                if (cycleLength == 2) {
                    totalChains += 2 + chainLengths[i] + chainLengths[favorite[i]];
                } else {
                    maxCycle = Math.max(maxCycle, cycleLength);
                }
            }
        }

        return Math.max(maxCycle, totalChains);
    }
}
