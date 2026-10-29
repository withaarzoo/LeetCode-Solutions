class Solution:
    def palindromePath(self, n: int, edges: list[list[int]], s: str, queries: list[str]) -> list[bool]:
        g = [[] for _ in range(n)]
        for u, v in edges:
            g[u].append(v)
            g[v].append(u)
        LOG = 16
        up = [[-1] * n for _ in range(LOG)]
        depth = [0] * n
        inn = [0] * n
        out = [0] * n
        parent = [-1] * n
        mask = [1 << (ord(c) - 97) for c in s]
        timer = [1]
        def dfs(u, p):
            inn[u] = timer[0]
            timer[0] += 1
            parent[u] = p
            up[0][u] = p
            for k in range(1, LOG):
                if up[k - 1][u] != -1:
                    up[k][u] = up[k - 1][up[k - 1][u]]
            for v in g[u]:
                if v != p:
                    depth[v] = depth[u] + 1
                    dfs(v, u)
            out[u] = timer[0]
        dfs(0, -1)
        bit = [0] * (n + 2)
        def upd(i, val):
            while i <= n:
                bit[i] ^= val
                i += i & -i
        def qry(i):
            res = 0
            while i > 0:
                res ^= bit[i]
                i -= i & -i
            return res
        def range_xor(l, r, val):
            if l > r:
                return
            upd(l, val)
            upd(r + 1, val)
        for i in range(n):
            range_xor(inn[i], out[i] - 1, mask[i])
        def get_lca(u, v):
            if depth[u] < depth[v]:
                u, v = v, u
            diff = depth[u] - depth[v]
            for k in range(LOG):
                if (diff >> k) & 1:
                    u = up[k][u]
            if u == v:
                return u
            for k in range(LOG - 1, -1, -1):
                if up[k][u] != up[k][v]:
                    u = up[k][u]
                    v = up[k][v]
            return parent[u]
        def path_xor(u, v):
            l = get_lca(u, v)
            return qry(inn[u]) ^ qry(inn[v]) ^ mask[l]
        ans = []
        for q in queries:
            if q[0] == 'u':
                u = 0
                i = 7
                while q[i] != ' ':
                    u = u * 10 + (ord(q[i]) - 48)
                    i += 1
                c = q[-1]
                nm = 1 << (ord(c) - 97)
                delta = mask[u] ^ nm
                mask[u] = nm
                range_xor(inn[u], out[u] - 1, delta)
            else:
                u = 0
                v = 0
                i = 6
                while q[i] != ' ':
                    u = u * 10 + (ord(q[i]) - 48)
                    i += 1
                i += 1
                while i < len(q):
                    v = v * 10 + (ord(q[i]) - 48)
                    i += 1
                m = path_xor(u, v)
                ans.append(bin(m).count('1') <= 1)
        return ans