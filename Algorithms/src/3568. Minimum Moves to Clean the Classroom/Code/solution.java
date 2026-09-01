class Solution {
    public int minMoves(String[] classroom, int energy) {
        int m = classroom.length;
        int n = classroom[0].length();

        // id[r][c] stores the bit number assigned to a litter cell.
        int[][] id = new int[m][n];

        // Initialize all IDs to -1 because non-litter cells have no bit.
        for (int r = 0; r < m; r++) {
            java.util.Arrays.fill(id[r], -1);
        }

        int k = 0;
        int sr = 0, sc = 0;

        // Find S and assign a unique bit to every L cell.
        for (int r = 0; r < m; r++) {
            for (int c = 0; c < n; c++) {
                if (classroom[r].charAt(c) == 'S') {
                    sr = r;
                    sc = c;
                } else if (classroom[r].charAt(c) == 'L') {
                    id[r][c] = k++;
                }
            }
        }

        // No litter means zero moves are needed.
        if (k == 0)
            return 0;

        int totalMask = (1 << k) - 1;

        // bestEnergy[r][c][mask] keeps the maximum energy
        // already seen for the same position and collected litter.
        int[][][] best = new int[m][n][1 << k];

        // Initialize every state as unseen.
        for (int r = 0; r < m; r++) {
            for (int c = 0; c < n; c++) {
                java.util.Arrays.fill(best[r][c], -1);
            }
        }

        // A state stores position, mask, remaining energy, and moves.
        class State {
            int r, c, mask, e, moves;

            State(int r, int c, int mask, int e, int moves) {
                this.r = r;
                this.c = c;
                this.mask = mask;
                this.e = e;
                this.moves = moves;
            }
        }

        // ArrayDeque gives efficient FIFO operations for BFS.
        java.util.ArrayDeque<State> queue = new java.util.ArrayDeque<>();

        // Start from S with no collected litter and full energy.
        best[sr][sc][0] = energy;
        queue.offer(new State(sr, sc, 0, energy, 0));

        // Four possible movement directions.
        int[] dr = { -1, 1, 0, 0 };
        int[] dc = { 0, 0, -1, 1 };

        while (!queue.isEmpty()) {
            State cur = queue.poll();

            // Try all four neighboring cells.
            for (int d = 0; d < 4; d++) {
                int nr = cur.r + dr[d];
                int nc = cur.c + dc[d];

                // Ignore cells outside the classroom.
                if (nr < 0 || nr >= m || nc < 0 || nc >= n)
                    continue;

                // Obstacles cannot be entered.
                if (classroom[nr].charAt(nc) == 'X')
                    continue;

                // Every move costs one energy.
                int ne = cur.e - 1;

                // Without enough energy, this move is impossible.
                if (ne < 0)
                    continue;

                int nmask = cur.mask;

                // R restores the student's energy to full capacity.
                if (classroom[nr].charAt(nc) == 'R') {
                    ne = energy;
                }

                // Collect the litter by setting its corresponding bit.
                if (classroom[nr].charAt(nc) == 'L') {
                    nmask |= (1 << id[nr][nc]);
                }

                // Return immediately when every litter item is collected.
                if (nmask == totalMask) {
                    return cur.moves + 1;
                }

                // A state with no better energy than an existing state
                // cannot lead to a better answer.
                if (ne <= best[nr][nc][nmask])
                    continue;

                // Store the strongest energy seen for this state.
                best[nr][nc][nmask] = ne;

                // Add the improved state to the BFS queue.
                queue.offer(new State(nr, nc, nmask, ne, cur.moves + 1));
            }
        }

        // No valid path can collect all litter.
        return -1;
    }
}