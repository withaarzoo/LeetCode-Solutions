class Solution:
    def maxTwoEvents(self, events):
        # Sort by start time
        events.sort()

        # Sort by end time
        end_sorted = sorted(events, key=lambda x: x[1])

        n = len(events)
        max_value_till = [0] * n

        max_value_till[0] = end_sorted[0][2]
        for i in range(1, n):
            max_value_till[i] = max(max_value_till[i - 1], end_sorted[i][2])

        ans = 0
        j = 0

        for start, end, value in events:
            while j < n and end_sorted[j][1] < start:
                j += 1

            ans = max(ans, value)
            if j > 0:
                ans = max(ans, value + max_value_till[j - 1])

        return ans
