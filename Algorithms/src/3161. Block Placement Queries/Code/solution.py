from typing import List

class Fenwick:
    def __init__(self, n: int) -> None:
        self.n = n
        self.bit = [0] * (n + 1)

    # Point add because I only insert obstacles.
    def add(self, idx: int, val: int) -> None:
        while idx <= self.n:
            self.bit[idx] += val
            idx += idx & -idx

    # Prefix sum because I need the number of occupied positions on the left.
    def sum(self, idx: int) -> int:
        res = 0
        while idx > 0:
            res += self.bit[idx]
            idx -= idx & -idx
        return res

    # Find the smallest index whose prefix sum is at least k.
    def kth(self, k: int) -> int:
        idx = 0
        step = 1
        while (step << 1) <= self.n:
            step <<= 1

        d = step
        while d:
            nxt = idx + d
            if nxt <= self.n and self.bit[nxt] < k:
                idx = nxt
                k -= self.bit[nxt]
            d >>= 1
        return idx + 1


class SegTree:
    def __init__(self, n: int) -> None:
        self.n = n
        self.tree = [0] * max(4, 4 * n)

    # Update one position because only one stored gap changes at a time.
    def update(self, node: int, l: int, r: int, pos: int, val: int) -> None:
        if l == r:
            self.tree[node] = val
            return
        mid = (l + r) // 2
        if pos <= mid:
            self.update(node * 2, l, mid, pos, val)
        else:
            self.update(node * 2 + 1, mid + 1, r, pos, val)
        self.tree[node] = max(self.tree[node * 2], self.tree[node * 2 + 1])

    # Query a prefix maximum because all useful gaps end at or before x.
    def query(self, node: int, l: int, r: int, ql: int, qr: int) -> int:
        if ql > r or qr < l:
            return 0
        if ql <= l and r <= qr:
            return self.tree[node]
        mid = (l + r) // 2
        return max(
            self.query(node * 2, l, mid, ql, qr),
            self.query(node * 2 + 1, mid + 1, r, ql, qr),
        )


class Solution:
    def getResults(self, queries: List[List[int]]) -> List[bool]:
        mx = 0
        for q in queries:
            mx = max(mx, q[1])

        fw = Fenwick(mx + 2)
        st = SegTree(mx + 1)

        # Sentinel obstacle at 0 makes predecessor search easy.
        fw.add(1, 1)

        ans: List[bool] = []

        for q in queries:
            t = q[0]
            x = q[1]

            if t == 1:
                # Count occupied positions strictly smaller than x.
                left_count = fw.sum(x)
                left_pos = fw.kth(left_count) - 1

                # Find the next occupied position after x if it exists.
                occupied_up_to_x = fw.sum(x + 1)
                total_occupied = fw.sum(mx + 2)
                right_pos = -1
                if occupied_up_to_x < total_occupied:
                    right_pos = fw.kth(occupied_up_to_x + 1) - 1

                # x becomes the end of a new gap.
                st.update(1, 0, mx, x, x - left_pos)

                # The next obstacle now sees a shorter gap.
                if right_pos != -1:
                    st.update(1, 0, mx, right_pos, right_pos - x)

                # Mark x as occupied.
                fw.add(x + 1, 1)
            else:
                sz = q[2]

                # The last obstacle strictly before x gives the free suffix up to x.
                left_count = fw.sum(x)
                left_pos = fw.kth(left_count) - 1

                # Prefix maximum of gaps ending at or before x.
                best_prefix = st.query(1, 0, mx, 0, x)

                ans.append((x - left_pos >= sz) or (best_prefix >= sz))

        return ans