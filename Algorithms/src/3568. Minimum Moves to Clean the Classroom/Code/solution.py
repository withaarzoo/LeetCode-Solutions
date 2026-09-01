from typing import List
from collections import deque

class Solution:
    def minMoves(self, classroom: List[str], energy: int) -> int:
        m = len(classroom)
        n = len(classroom[0])

        # id[r][c] stores the bit assigned to a litter cell.
        id = [[-1] * n for _ in range(m)]

        k = 0
        sr = 0
        sc = 0

        # Find the starting position and assign bits to all litter cells.
        for r in range(m):
            for c in range(n):
                if classroom[r][c] == 'S':
                    sr = r
                    sc = c
                elif classroom[r][c] == 'L':
                    id[r][c] = k
                    k += 1

        # If there is no litter, the task is already complete.
        if k == 0:
            return 0

        # This mask has all k litter bits turned on.
        total_mask = (1 << k) - 1

        # best[r][c][mask] stores the maximum energy seen
        # for this position and collected-litter mask.
        best = [
            [
                [-1] * (1 << k)
                for _ in range(n)
            ]
            for _ in range(m)
        ]

        # BFS state is (row, column, mask, remaining energy, moves).
        queue = deque()

        # Start at S with no collected litter and full energy.
        best[sr][sc][0] = energy
        queue.append((sr, sc, 0, energy, 0))

        # Four possible movement directions.
        directions = [(-1, 0), (1, 0), (0, -1), (0, 1)]

        while queue:
            # Get the next state in BFS order.
            r, c, mask, e, moves = queue.popleft()

            # Try moving to all four neighboring cells.
            for dr, dc in directions:
                nr = r + dr
                nc = c + dc

                # Ignore positions outside the classroom.
                if nr < 0 or nr >= m or nc < 0 or nc >= n:
                    continue

                # Obstacles cannot be entered.
                if classroom[nr][nc] == 'X':
                    continue

                # Every movement uses one unit of energy.
                ne = e - 1

                # We cannot move if energy becomes negative.
                if ne < 0:
                    continue

                nmask = mask

                # Reset energy when entering an R cell.
                if classroom[nr][nc] == 'R':
                    ne = energy

                # Mark the litter as collected by setting its bit.
                if classroom[nr][nc] == 'L':
                    nmask |= 1 << id[nr][nc]

                # All litter has been collected, so return the move count.
                if nmask == total_mask:
                    return moves + 1

                # If an equal or stronger state already exists,
                # this new state cannot give us a better path.
                if ne <= best[nr][nc][nmask]:
                    continue

                # Keep the maximum energy for this position and mask.
                best[nr][nc][nmask] = ne

                # Add the improved state to the BFS queue.
                queue.append((nr, nc, nmask, ne, moves + 1))

        # BFS could not find a path that collects all litter.
        return -1