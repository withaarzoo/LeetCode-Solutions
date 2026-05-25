class Solution:
    def canReach(self, s: str, minJump: int, maxJump: int) -> bool:

        n = len(s)

        # Queue for BFS traversal
        from collections import deque
        q = deque([0])

        # Visited array
        visited = [False] * n
        visited[0] = True

        # Farthest processed position
        far = 0

        while q:

            i = q.popleft()

            # If last index is reached
            if i == n - 1:
                return True

            # Valid jump range
            start = max(i + minJump, far + 1)
            end = min(i + maxJump, n - 1)

            # Explore all possible next positions
            for j in range(start, end + 1):

                # Only move to positions containing '0'
                if s[j] == '0' and not visited[j]:
                    visited[j] = True
                    q.append(j)

            # Update processed boundary
            far = max(far, end)

        return False