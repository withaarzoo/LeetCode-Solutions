class Solution {
    public int[] pathsWithMaxScore(List<String> board) {
        final int MOD = 1_000_000_007;
        int n = board.size();

        // These arrays store DP values for the row below.
        int[] nextScore = new int[n + 1];
        int[] nextWays = new int[n + 1];

        // -1 marks every score as unreachable at the beginning.
        Arrays.fill(nextScore, -1);

        // I process the board from bottom to top.
        for (int i = n - 1; i >= 0; i--) {
            // These arrays store values for the current row.
            int[] currScore = new int[n + 1];
            int[] currWays = new int[n + 1];

            // Every current-row cell starts as unreachable.
            Arrays.fill(currScore, -1);

            // I move right to left so the right cell is already solved.
            for (int j = n - 1; j >= 0; j--) {
                char cell = board.get(i).charAt(j);

                // Obstacles cannot be used.
                if (cell == 'X') {
                    continue;
                }

                // S starts with score 0 and exactly one path.
                if (cell == 'S') {
                    currScore[j] = 0;
                    currWays[j] = 1;
                    continue;
                }

                // Check down, right, and bottom-right diagonal.
                int best = Math.max(
                    nextScore[j],
                    Math.max(currScore[j + 1], nextScore[j + 1])
                );

                // No reachable next cell means this cell is also unreachable.
                if (best == -1) {
                    continue;
                }

                long ways = 0;

                // Count every next cell that gives the maximum score.
                if (nextScore[j] == best) {
                    ways += nextWays[j];
                }
                if (currScore[j + 1] == best) {
                    ways += currWays[j + 1];
                }
                if (nextScore[j + 1] == best) {
                    ways += nextWays[j + 1];
                }

                // E contributes 0; digit cells contribute their digit.
                int value = (cell == 'E') ? 0 : cell - '0';

                currScore[j] = best + value;
                currWays[j] = (int) (ways % MOD);
            }

            // Reuse the completed current row as the next row.
            nextScore = currScore;
            nextWays = currWays;
        }

        // An unreachable E means no valid path exists.
        if (nextScore[0] == -1) {
            return new int[]{0, 0};
        }

        return new int[]{nextScore[0], nextWays[0]};
    }
}