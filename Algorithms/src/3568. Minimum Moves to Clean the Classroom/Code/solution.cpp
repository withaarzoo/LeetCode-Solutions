class Solution
{
public:
    int minMoves(vector<string> &classroom, int energy)
    {
        int m = classroom.size();
        int n = classroom[0].size();

        // Give every litter cell a unique bit position.
        vector<vector<int>> id(m, vector<int>(n, -1));

        int k = 0;
        int sr = 0, sc = 0;

        // Find the starting position and assign IDs to all litter cells.
        for (int r = 0; r < m; r++)
        {
            for (int c = 0; c < n; c++)
            {
                if (classroom[r][c] == 'S')
                {
                    sr = r;
                    sc = c;
                }
                else if (classroom[r][c] == 'L')
                {
                    id[r][c] = k++;
                }
            }
        }

        // If there is no litter, the student is already done.
        if (k == 0)
            return 0;

        int totalMask = (1 << k) - 1;

        // bestEnergy[r][c][mask] stores the maximum energy seen
        // at this position after collecting exactly this mask.
        vector<vector<vector<int>>> best(
            m, vector<vector<int>>(n, vector<int>(1 << k, -1)));

        // Each BFS state contains position, collected litter mask,
        // remaining energy, and number of moves used.
        struct State
        {
            int r, c, mask, e, moves;
        };

        queue<State> q;

        // Initially, no litter is collected and full energy is available.
        best[sr][sc][0] = energy;
        q.push({sr, sc, 0, energy, 0});

        // Four possible movement directions.
        int dr[] = {-1, 1, 0, 0};
        int dc[] = {0, 0, -1, 1};

        while (!q.empty())
        {
            State cur = q.front();
            q.pop();

            // Try moving in all four directions.
            for (int d = 0; d < 4; d++)
            {
                int nr = cur.r + dr[d];
                int nc = cur.c + dc[d];

                // Ignore positions outside the classroom.
                if (nr < 0 || nr >= m || nc < 0 || nc >= n)
                    continue;

                // The student cannot move through obstacles.
                if (classroom[nr][nc] == 'X')
                    continue;

                // Every movement consumes one unit of energy.
                int ne = cur.e - 1;

                // The student cannot make a move without energy.
                if (ne < 0)
                    continue;

                int nmask = cur.mask;

                // Reset the energy immediately after entering an R cell.
                if (classroom[nr][nc] == 'R')
                {
                    ne = energy;
                }

                // If this cell contains litter, mark its bit as collected.
                if (classroom[nr][nc] == 'L')
                {
                    nmask |= (1 << id[nr][nc]);
                }

                // All litter has been collected, so this is the answer.
                if (nmask == totalMask)
                {
                    return cur.moves + 1;
                }

                // If we already reached this position with the same
                // mask and at least this much energy, this state is useless.
                if (ne <= best[nr][nc][nmask])
                    continue;

                // Keep only the strongest energy value for this state.
                best[nr][nc][nmask] = ne;

                // Add the improved state to BFS.
                q.push({nr, nc, nmask, ne, cur.moves + 1});
            }
        }

        // BFS finished without collecting all litter.
        return -1;
    }
};