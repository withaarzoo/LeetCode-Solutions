class Solution:
    def removeCoveredIntervals(self, intervals: List[List[int]]) -> int:
        # Sort by start in ascending order.
        # For equal starts, -end puts the larger end first,
        # so the covering interval is processed before the covered one.
        intervals.sort(key=lambda interval: (interval[0], -interval[1]))

        # Every valid end is greater than 0, so -1 is a safe initial value.
        max_end = -1

        # This stores how many intervals are not covered.
        remaining = 0

        # Check every interval in the sorted order.
        for start, end in intervals:
            # A larger end means no earlier interval can cover this one.
            if end > max_end:
                remaining += 1

                # Remember the farthest ending point seen so far.
                max_end = end

            # Otherwise, end <= max_end, so this interval is covered.

        # Return the number of intervals that remain.
        return remaining