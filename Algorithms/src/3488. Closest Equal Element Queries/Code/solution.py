class Solution:
    def solveQueries(self, nums: List[int], queries: List[int]) -> List[int]:
        n = len(nums)

        # Store all indices for every value
        positions = {}

        for i, num in enumerate(nums):
            if num not in positions:
                positions[num] = []
            positions[num].append(i)

        # answer[i] = minimum circular distance for index i
        answer = [-1] * n

        # Process each group of equal values
        for pos in positions.values():
            m = len(pos)

            # If value appears once, answer stays -1
            if m == 1:
                continue

            for i in range(m):
                curr = pos[i]

                # Previous and next occurrence in circular order
                prev_idx = pos[(i - 1 + m) % m]
                next_idx = pos[(i + 1) % m]

                # Distance to previous occurrence
                dist_prev = abs(curr - prev_idx)
                dist_prev = min(dist_prev, n - dist_prev)

                # Distance to next occurrence
                dist_next = abs(curr - next_idx)
                dist_next = min(dist_next, n - dist_next)

                # Best answer for current index
                answer[curr] = min(dist_prev, dist_next)

        # Return answers for queries
        return [answer[idx] for idx in queries]