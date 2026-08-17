from typing import List

class Solution:
    def stoneGameV(self, stoneValue: List[int]) -> int:
        n = len(stoneValue)

        # prefix[i] stores the sum of the first i stones.
        # This makes every interval sum available in O(1).
        prefix = [0] * (n + 1)

        for i in range(n):
            # Build the prefix sum array.
            prefix[i + 1] = prefix[i] + stoneValue[i]

        # dp[l][r] stores Alice's maximum score for interval [l, r].
        dp = [[0] * n for _ in range(n)]

        # left_best[l][r] stores the best value of:
        # dp[l][k] + sum(l, k).
        left_best = [[0] * n for _ in range(n)]

        # right_best[l][r] stores the best value of:
        # dp[k][r] + sum(k, r).
        right_best = [[0] * n for _ in range(n)]

        # left_ptr[l] is the last split where leftSum <= rightSum.
        left_ptr = [0] * n

        # right_ptr[l] is the first split where leftSum >= rightSum.
        right_ptr = list(range(n))

        for i in range(n):
            # A single stone cannot be split, so its score is 0.
            # But the helper tables must contain the stone value.
            left_best[i][i] = stoneValue[i]
            right_best[i][i] = stoneValue[i]

            # No valid split exists before index i.
            left_ptr[i] = i - 1

            # The first possible split starts at i.
            right_ptr[i] = i

        # Process intervals from shorter to longer intervals.
        # This ensures every smaller dp state is already calculated.
        for length in range(2, n + 1):
            for l in range(n - length + 1):
                r = l + length - 1

                # Calculate the total sum of the current interval.
                total = prefix[r + 1] - prefix[l]

                # Move left_ptr while leftSum <= rightSum.
                while left_ptr[l] + 1 <= r - 1:
                    k = left_ptr[l] + 1
                    left_sum = prefix[k + 1] - prefix[l]

                    # 2 * left_sum <= total means left_sum <= right_sum.
                    if 2 * left_sum > total:
                        break

                    left_ptr[l] += 1

                # Move right_ptr until leftSum >= rightSum.
                while right_ptr[l] <= r - 1:
                    k = right_ptr[l]
                    left_sum = prefix[k + 1] - prefix[l]

                    # Stop at the first split where the left side
                    # becomes at least as large as the right side.
                    if 2 * left_sum >= total:
                        break

                    right_ptr[l] += 1

                best = 0

                # If the left side is smaller or equal, Alice keeps it.
                if left_ptr[l] >= l:
                    best = left_best[l][left_ptr[l]]

                # If the right side is smaller or equal, Alice keeps it.
                if right_ptr[l] <= r - 1:
                    best = max(
                        best,
                        right_best[right_ptr[l] + 1][r]
                    )

                # Store Alice's maximum score for [l, r].
                dp[l][r] = best

                # Use [l, r] as a possible left side in a larger interval.
                left_best[l][r] = max(
                    left_best[l][r - 1],
                    dp[l][r] + total
                )

                # Use [l, r] as a possible right side in a larger interval.
                right_best[l][r] = max(
                    right_best[l + 1][r],
                    dp[l][r] + total
                )

        # Return the result for the complete array.
        return dp[0][n - 1]