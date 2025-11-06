from typing import List
import heapq

class Solution:
    def processQueries(self, c: int, connections: List[List[int]], queries: List[List[int]]) -> List[int]:
        # DSU
        parent = list(range(c + 1))
        size = [1] * (c + 1)

        def find(x: int) -> int:
            while parent[x] != x:
                parent[x] = parent[parent[x]]
                x = parent[x]
            return x

        def unite(a: int, b: int) -> None:
            ra, rb = find(a), find(b)
            if ra == rb:
                return
            if size[ra] < size[rb]:
                ra, rb = rb, ra
            parent[rb] = ra
            size[ra] += size[rb]

        for u, v in connections:
            unite(u, v)

        # root -> min-heap of member ids
        heaps = {}
        for i in range(1, c + 1):
            r = find(i)
            if r not in heaps:
                heaps[r] = []
            heaps[r].append(i)
        for r in heaps:
            heapq.heapify(heaps[r])

        offline = [False] * (c + 1)
        ans = []

        for t, x in queries:
            if t == 2:
                offline[x] = True
            else:
                if not offline[x]:
                    ans.append(x)
                else:
                    r = find(x)
                    h = heaps.get(r, [])
                    # Lazy deletion of offline nodes
                    while h and offline[h[0]]:
                        heapq.heappop(h)
                    ans.append(h[0] if h else -1)

        return ans
