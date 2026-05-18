class Solution {
    public int minJumps(int[] arr) {

        int n = arr.length;

        // No jump needed if only one element exists
        if (n == 1)
            return 0;

        // Store all indices for every value
        HashMap<Integer, List<Integer>> map = new HashMap<>();

        for (int i = 0; i < n; i++) {
            map.computeIfAbsent(arr[i], k -> new ArrayList<>()).add(i);
        }

        // Queue for BFS
        Queue<Integer> queue = new LinkedList<>();

        // Visited array
        boolean[] visited = new boolean[n];

        queue.offer(0);
        visited[0] = true;

        int steps = 0;

        while (!queue.isEmpty()) {

            int size = queue.size();

            // Process current BFS level
            while (size-- > 0) {

                int idx = queue.poll();

                // Last index reached
                if (idx == n - 1) {
                    return steps;
                }

                // Move left
                if (idx - 1 >= 0 && !visited[idx - 1]) {
                    visited[idx - 1] = true;
                    queue.offer(idx - 1);
                }

                // Move right
                if (idx + 1 < n && !visited[idx + 1]) {
                    visited[idx + 1] = true;
                    queue.offer(idx + 1);
                }

                // Move to same-value indices
                for (int nextIdx : map.get(arr[idx])) {

                    if (!visited[nextIdx]) {
                        visited[nextIdx] = true;
                        queue.offer(nextIdx);
                    }
                }

                // Remove repeated processing
                map.get(arr[idx]).clear();
            }

            // One level completed
            steps++;
        }

        return -1;
    }
}