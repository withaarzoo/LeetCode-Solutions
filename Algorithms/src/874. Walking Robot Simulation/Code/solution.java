class Solution {
    public int robotSim(int[] commands, int[][] obstacles) {
        // Store obstacles in a HashSet
        Set<String> obstacleSet = new HashSet<>();

        for (int[] obs : obstacles) {
            obstacleSet.add(obs[0] + "," + obs[1]);
        }

        // Directions: North, East, South, West
        int[] dx = { 0, 1, 0, -1 };
        int[] dy = { 1, 0, -1, 0 };

        int dir = 0; // North
        int x = 0, y = 0;
        int maxDistance = 0;

        for (int command : commands) {
            // Turn right
            if (command == -1) {
                dir = (dir + 1) % 4;
            }
            // Turn left
            else if (command == -2) {
                dir = (dir + 3) % 4;
            }
            // Move forward
            else {
                for (int step = 0; step < command; step++) {
                    int nextX = x + dx[dir];
                    int nextY = y + dy[dir];

                    String nextPos = nextX + "," + nextY;

                    // Stop if obstacle exists
                    if (obstacleSet.contains(nextPos)) {
                        break;
                    }

                    x = nextX;
                    y = nextY;

                    maxDistance = Math.max(maxDistance, x * x + y * y);
                }
            }
        }

        return maxDistance;
    }
}