class Solution:
    def resultArray(self, nums: List[int], k: int, queries: List[List[int]]) -> List[int]:
        class Node:
            __slots__ = ('prod', 'freq')
            def __init__(self):
                self.prod = 1
                self.freq = [0] * k
        
        n = len(nums)
        tree = [Node() for _ in range(4 * n)]
        
        def merge(L: Node, R: Node) -> Node:
            res = Node()
            res.prod = (L.prod * R.prod) % k
            res.freq = L.freq[:]
            for r in range(k):
                if R.freq[r]:
                    nr = (L.prod * r) % k
                    res.freq[nr] += R.freq[r]
            return res
        
        def build(v: int, tl: int, tr: int) -> None:
            if tl == tr:
                tree[v].prod = nums[tl] % k
                tree[v].freq[tree[v].prod] = 1
                return
            tm = (tl + tr) // 2
            build(v * 2, tl, tm)
            build(v * 2 + 1, tm + 1, tr)
            tree[v] = merge(tree[v * 2], tree[v * 2 + 1])
        
        def update(v: int, tl: int, tr: int, pos: int, val: int) -> None:
            if tl == tr:
                tree[v].prod = val % k
                tree[v].freq = [0] * k
                tree[v].freq[tree[v].prod] = 1
                return
            tm = (tl + tr) // 2
            if pos <= tm:
                update(v * 2, tl, tm, pos, val)
            else:
                update(v * 2 + 1, tm + 1, tr, pos, val)
            tree[v] = merge(tree[v * 2], tree[v * 2 + 1])
        
        def query(v: int, tl: int, tr: int, l: int, r: int) -> Node:
            if l > r:
                return Node()
            if l == tl and r == tr:
                return tree[v]
            tm = (tl + tr) // 2
            return merge(query(v * 2, tl, tm, l, min(r, tm)),
                         query(v * 2 + 1, tm + 1, tr, max(l, tm + 1), r))
        
        build(1, 0, n - 1)
        ans = []
        for idx, val, start, x in queries:
            update(1, 0, n - 1, idx, val)
            res = query(1, 0, n - 1, start, n - 1)
            ans.append(res.freq[x])
        return ans