class Solution {
public:
    vector<int> pathsWithMaxScore(vector<string>& board) {
        const int MOD = 1000000007;
        int n = board.size();

        // These arrays store the DP values for the row below.
        // A score of -1 means the cell cannot reach S.
        vector<int> nextScore(n + 1, -1);
        vector<int> nextWays(n + 1, 0);

        // I process rows from bottom to top.
        for (int i = n - 1; i >= 0; --i) {
            // Fresh arrays are needed for the current row.
            vector<int> currScore(n + 1, -1);
            vector<int> currWays(n + 1, 0);

            // I process right to left so the right cell is already solved.
            for (int j = n - 1; j >= 0; --j) {
                char cell = board[i][j];

                // An obstacle can never be part of a valid path.
                if (cell == 'X') {
                    continue;
                }

                // S is the starting point of the reversed DP.
                if (cell == 'S') {
                    currScore[j] = 0;
                    currWays[j] = 1;
                    continue;
                }

                // Find the best score among down, right, and diagonal.
                int best = max({
                    nextScore[j],
                    currScore[j + 1],
                    nextScore[j + 1]
                });

                // If all three cells are unreachable, this cell is unreachable too.
                if (best == -1) {
                    continue;
                }

                long long ways = 0;

                // Add counts only from cells that achieve the best score.
                if (nextScore[j] == best) {
                    ways += nextWays[j];
                }
                if (currScore[j + 1] == best) {
                    ways += currWays[j + 1];
                }
                if (nextScore[j + 1] == best) {
                    ways += nextWays[j + 1];
                }

                // E adds no points; a digit adds its numeric value.
                int value = (cell == 'E') ? 0 : cell - '0';

                currScore[j] = best + value;
                currWays[j] = ways % MOD;
            }

            // The current row becomes the row below for the next iteration.
            nextScore = move(currScore);
            nextWays = move(currWays);
        }

        // If E cannot reach S, the required answer is [0, 0].
        if (nextScore[0] == -1) {
            return {0, 0};
        }

        return {nextScore[0], nextWays[0]};
    }
};