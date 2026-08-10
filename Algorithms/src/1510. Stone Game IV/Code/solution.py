class Solution:
    def winnerSquareGame(self, n: int) -> bool:
        # dp[i] tells me whether the current player can win with i stones.
        dp = [False] * (n + 1)

        # With 0 stones, the current player has no move and loses.
        dp[0] = False

        # Calculate the result for every number of stones from 1 to n.
        for i in range(1, n + 1):
            # Try every perfect square that can be removed from i.
            j = 1
            while j * j <= i:
                # If the remaining state is losing for the opponent,
                # I can make this move and force a win.
                if not dp[i - j * j]:
                    dp[i] = True

                    # One winning move is enough, so stop checking squares.
                    break

                # Move to the next possible perfect square.
                j += 1

        # Return whether Alice can force a win with n stones.
        return dp[n]