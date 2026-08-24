from typing import List

class Solution:
    def stoneGameVIII(self, stones: List[int]) -> int:
        # Convert stones into prefix sums directly to save extra space.
        for i in range(1, len(stones)):
            stones[i] += stones[i - 1]

        # Taking every stone gives the total sum for the first state.
        best = stones[-1]

        # Process each possible prefix from right to left.
        for i in range(len(stones) - 2, 0, -1):
            # Choose between the current best result and taking this prefix,
            # then subtracting the best score difference of the opponent.
            best = max(best, stones[i] - best)

        # Return the maximum score difference Alice can force.
        return best