class Solution:
    def pathsWithMaxScore(self, board: List[str]) -> List[int]:
        MOD = 10**9 + 7
        n = len(board)

        # These arrays store DP values for the row below.
        # A score of -1 means the cell cannot reach S.
        next_score = [-1] * (n + 1)
        next_ways = [0] * (n + 1)

        # I process rows from bottom to top.
        for i in range(n - 1, -1, -1):
            # Fresh arrays store the current row.
            curr_score = [-1] * (n + 1)
            curr_ways = [0] * (n + 1)

            # I process right to left so the right cell is already solved.
            for j in range(n - 1, -1, -1):
                cell = board[i][j]

                # Obstacles cannot be part of any valid path.
                if cell == 'X':
                    continue

                # S is the base case for the reversed DP.
                if cell == 'S':
                    curr_score[j] = 0
                    curr_ways[j] = 1
                    continue

                # Check down, right, and bottom-right diagonal.
                best = max(
                    next_score[j],
                    curr_score[j + 1],
                    next_score[j + 1]
                )

                # If no next cell can reach S, this cell cannot either.
                if best == -1:
                    continue

                ways = 0

                # Add counts only from next cells with the best score.
                if next_score[j] == best:
                    ways += next_ways[j]
                if curr_score[j + 1] == best:
                    ways += curr_ways[j + 1]
                if next_score[j + 1] == best:
                    ways += next_ways[j + 1]

                # E contributes 0; a digit contributes its integer value.
                value = 0 if cell == 'E' else int(cell)

                curr_score[j] = best + value
                curr_ways[j] = ways % MOD

            # The completed row becomes the row below.
            next_score = curr_score
            next_ways = curr_ways

        # If E cannot reach S, there is no valid path.
        if next_score[0] == -1:
            return [0, 0]

        return [next_score[0], next_ways[0]]