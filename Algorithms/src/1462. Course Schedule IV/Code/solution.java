class Solution {
    public List<Boolean> checkIfPrerequisite(int numCourses, int[][] prerequisites, int[][] queries) {
        // Initialize the graph
        boolean[][] graph = new boolean[numCourses][numCourses];

        // Build the direct edges from prerequisites
        for (int[] edge : prerequisites) {
            graph[edge[0]][edge[1]] = true;
        }

        // Floyd-Warshall to compute transitive closure
        for (int k = 0; k < numCourses; k++) {
            for (int i = 0; i < numCourses; i++) {
                for (int j = 0; j < numCourses; j++) {
                    if (graph[i][k] && graph[k][j]) {
                        graph[i][j] = true;
                    }
                }
            }
        }

        // Answer the queries
        List<Boolean> result = new ArrayList<>();
        for (int[] query : queries) {
            result.add(graph[query[0]][query[1]]);
        }

        return result;
    }
}
