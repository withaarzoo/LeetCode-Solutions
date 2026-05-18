class Solution:
    def minJumps(self, arr: List[int]) -> int:

        n = len(arr)

        # No jump needed
        if n == 1:
            return 0

        from collections import defaultdict, deque

        # Store all indices for every value
        mp = defaultdict(list)

        for i, val in enumerate(arr):
            mp[val].append(i)

        # BFS queue
        q = deque([0])

        # Visited array
        visited = [False] * n

        visited[0] = True

        steps = 0

        while q:

            size = len(q)

            # Process one BFS level
            for _ in range(size):

                idx = q.popleft()

                # Last index reached
                if idx == n - 1:
                    return steps

                # Move left
                if idx - 1 >= 0 and not visited[idx - 1]:
                    visited[idx - 1] = True
                    q.append(idx - 1)

                # Move right
                if idx + 1 < n and not visited[idx + 1]:
                    visited[idx + 1] = True
                    q.append(idx + 1)

                # Move to same-value indices
                for next_idx in mp[arr[idx]]:

                    if not visited[next_idx]:
                        visited[next_idx] = True
                        q.append(next_idx)

                # Clear processed group
                mp[arr[idx]].clear()

            # Next BFS level
            steps += 1

        return -1