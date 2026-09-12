from typing import List


class Solution:
    def maximumWeight(self, intervals: List[List[int]]) -> List[int]:
        n = len(intervals)  # Number of intervals.
        K = 4               # I can choose at most four intervals.

        # Add the original index because sorting changes the order.
        intervals = [
            [l, r, w, i]
            for i, (l, r, w) in enumerate(intervals)
        ]

        # Sort by right endpoint so compatible intervals form a prefix.
        intervals.sort(key=lambda x: x[1])

        # Store right endpoints for binary search.
        ends = [x[1] for x in intervals]

        # A DP state is [score, sorted original indices].
        # None means that the state cannot be formed.
        dp = [[None] * (n + 1) for _ in range(K + 1)]

        # Selecting zero intervals always gives score 0.
        for i in range(n + 1):
            dp[0][i] = (0, [])

        # Compares two states and returns the better one.
        def better(a, b):
            # A valid state beats an unreachable state.
            if a is None:
                return False
            if b is None:
                return True

            # Maximum score is the primary condition.
            if a[0] != b[0]:
                return a[0] > b[0]

            # Equal scores require the lexicographically smallest indices.
            return a[1] < b[1]

        # Finds the first endpoint >= target.
        def lower_bound(length, target):
            lo = 0
            hi = length

            # Standard binary search.
            while lo < hi:
                mid = lo + (hi - lo) // 2

                # end >= target is incompatible because boundaries overlap.
                if ends[mid] >= target:
                    hi = mid
                else:
                    # end < target is compatible.
                    lo = mid + 1

            # This is the number of compatible previous intervals.
            return lo

        # Process intervals in sorted order.
        for i in range(1, n + 1):
            l, _, w, original_index = intervals[i - 1]

            # Find the compatible prefix before the current interval.
            p = lower_bound(i - 1, l)

            # Try choosing exactly k intervals.
            for k in range(1, K + 1):
                # Option 1: skip the current interval.
                dp[k][i] = dp[k][i - 1]

                # Option 2: take the current interval.
                prev = dp[k - 1][p]

                if prev is not None:
                    # Copy previous indices so this DP state stays independent.
                    ids = prev[1] + [original_index]

                    # Sort indices because the answer is compared lexicographically.
                    ids.sort()

                    # Build the candidate produced by taking this interval.
                    take = (prev[0] + w, ids)

                    # Keep whichever option is better.
                    if better(take, dp[k][i]):
                        dp[k][i] = take

        # The answer can contain 1, 2, 3, or 4 intervals.
        answer = None

        # Compare every possible number of selected intervals.
        for k in range(1, K + 1):
            if better(dp[k][n], answer):
                answer = dp[k][n]

        # Return only the selected original indices.
        return answer[1]