class Solution:
    def maximumInvitations(self, favorite: List[int]) -> int:
        n = len(favorite)
        inDegree = [0] * n
        chainLengths = [0] * n
        visited = [False] * n

        for fav in favorite:
            inDegree[fav] += 1

        queue = deque(i for i in range(n) if inDegree[i] == 0)
        while queue:
            node = queue.popleft()
            visited[node] = True

            next_node = favorite[node]
            chainLengths[next_node] = chainLengths[node] + 1
            inDegree[next_node] -= 1
            if inDegree[next_node] == 0:
                queue.append(next_node)

        maxCycle = 0
        totalChains = 0
        for i in range(n):
            if not visited[i]:
                current, cycleLength = i, 0
                while not visited[current]:
                    visited[current] = True
                    current = favorite[current]
                    cycleLength += 1

                if cycleLength == 2:
                    totalChains += 2 + chainLengths[i] + chainLengths[favorite[i]]
                else:
                    maxCycle = max(maxCycle, cycleLength)

        return max(maxCycle, totalChains)
