class Solution:
    def findSubtreeSizes(self, parent: List[int], s: str) -> List[int]:
        n = len(parent)
        g = [[] for _ in range(n)]
        for i in range(1, n):
            g[parent[i]].append(i)
        np = parent[:]
        last = [-1] * 26
        def dfs(u):
            c = ord(s[u]) - 97
            prev = last[c]
            if prev != -1:
                np[u] = prev
            last[c] = u
            for v in g[u]:
                dfs(v)
            last[c] = prev
        dfs(0)
        ng = [[] for _ in range(n)]
        for i in range(1, n):
            ng[np[i]].append(i)
        ans = [0] * n
        def size(u):
            sz = 1
            for v in ng[u]:
                sz += size(v)
            ans[u] = sz
            return sz
        size(0)
        return ans