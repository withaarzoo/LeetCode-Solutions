class Solution:
    def pathExistenceQueries(self, n: int, nums: List[int], maxDiff: int, queries: List[List[int]]) -> List[int]:
        # Sort original node indices according to their nums values.
        order = sorted(range(n), key=lambda i: nums[i])

        # pos[node] gives the sorted position of an original node.
        pos = [0] * n

        # values stores nums in sorted order.
        values = [0] * n

        for i in range(n):
            values[i] = nums[order[i]]
            pos[order[i]] = i

        # bit_length gives enough levels to represent every possible jump count.
        LOG = max(1, n.bit_length())

        # jump[p][i] is the farthest position after at most 2^p greedy jumps.
        jump = [[0] * n for _ in range(LOG)]

        # Use two pointers to find the farthest one-edge reach from every position.
        r = 0

        for i in range(n):
            # Keep r at or after the current position.
            if r < i:
                r = i

            # Extend r while a direct edge from i is still allowed.
            while r + 1 < n and values[r + 1] - values[i] <= maxDiff:
                r += 1

            # Save the farthest position reachable in one jump.
            jump[0][i] = r

        # Build larger jumps by combining two smaller jumps.
        for p in range(1, LOG):
            for i in range(n):
                # Applying level p - 1 twice gives 2^p greedy jumps.
                jump[p][i] = jump[p - 1][jump[p - 1][i]]

        # Store the answer for every query.
        answer = []

        for u, v in queries:
            # Convert original nodes into sorted positions.
            left = pos[u]
            right = pos[v]

            # The graph is undirected, so always process left to right.
            if left > right:
                left, right = right, left

            # The distance from a node to itself is zero.
            if left == right:
                answer.append(0)
                continue

            current = left
            distance = 0

            # Take the largest groups of jumps that still stop before the target.
            for p in range(LOG - 1, -1, -1):
                if jump[p][current] < right:
                    # Skip 2^p greedy jumps at once.
                    current = jump[p][current]
                    distance += 1 << p

            # Check whether one final edge reaches the target.
            if jump[0][current] >= right:
                answer.append(distance + 1)
            else:
                # No path exists if the greedy jump cannot move toward the target.
                answer.append(-1)

        return answer