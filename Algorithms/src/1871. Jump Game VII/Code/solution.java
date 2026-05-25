class Solution {
    public boolean canReach(String s, int minJump, int maxJump) {

        int n = s.length();

        // Queue for BFS traversal
        Queue<Integer> queue = new LinkedList<>();

        // Visited array
        boolean[] visited = new boolean[n];

        // Start from index 0
        queue.offer(0);
        visited[0] = true;

        // Farthest processed index
        int far = 0;

        while (!queue.isEmpty()) {

            int i = queue.poll();

            // Reached destination
            if (i == n - 1) {
                return true;
            }

            // Valid jump range
            int start = Math.max(i + minJump, far + 1);
            int end = Math.min(i + maxJump, n - 1);

            // Check every possible next position
            for (int j = start; j <= end; j++) {

                // Only move to '0'
                if (s.charAt(j) == '0' && !visited[j]) {
                    visited[j] = true;
                    queue.offer(j);
                }
            }

            // Update processed boundary
            far = Math.max(far, end);
        }

        return false;
    }
}