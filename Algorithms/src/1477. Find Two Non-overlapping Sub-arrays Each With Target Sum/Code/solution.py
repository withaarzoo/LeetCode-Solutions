from typing import List


class Solution:
    def minSumOfLengths(self, arr: List[int], target: int) -> int:
        n = len(arr)

        # best[i] stores the shortest target-sum sub-array
        # found completely inside arr[0..i].
        best = [float("inf")] * n

        left = 0              # Left boundary of the sliding window.
        total = 0             # Sum of the current sliding window.
        answer = float("inf") # Stores the minimum combined length.

        for right in range(n):
            # Expand the window by adding arr[right].
            total += arr[right]

            # All elements are positive, so moving left forward
            # always decreases the window sum.
            while total > target:
                total -= arr[left]
                left += 1

            # The current window is a valid target-sum sub-array.
            if total == target:
                current_length = right - left + 1

                # A previous sub-array must end before left,
                # so best[left - 1] guarantees no overlap.
                if left > 0 and best[left - 1] != float("inf"):
                    answer = min(
                        answer,
                        current_length + best[left - 1]
                    )

                # Save the current valid sub-array length
                # for the prefix ending at right.
                best[right] = current_length

            # Keep the shortest valid sub-array found anywhere
            # in the prefix arr[0..right].
            if right > 0:
                best[right] = min(best[right], best[right - 1])

        # If two non-overlapping sub-arrays were not found, return -1.
        return -1 if answer == float("inf") else answer