class Solution:
    def closestTarget(self, words: List[str], target: str, startIndex: int) -> int:
        n = len(words)
        ans = float('inf')

        # Traverse all indices
        for i in range(n):
            # If current word matches target
            if words[i] == target:
                diff = abs(i - startIndex)

                # Distance if we go around the circular array
                circular_dist = n - diff

                # Update minimum answer
                ans = min(ans, diff, circular_dist)

        # If target was not found
        return -1 if ans == float('inf') else ans