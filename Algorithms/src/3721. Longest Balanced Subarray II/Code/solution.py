# Python3 (concise and commented)
from typing import List
class SegTree:
    def __init__(self, n):
        self.n = n
        self.mn = [0] * (4*n)
        self.mx = [0] * (4*n)
        self.lazy = [0] * (4*n)
    def apply(self, idx, v):
        self.mn[idx] += v
        self.mx[idx] += v
        self.lazy[idx] += v
    def push(self, idx):
        z = self.lazy[idx]
        if z:
            self.apply(idx<<1, z)
            self.apply(idx<<1|1, z)
            self.lazy[idx] = 0
    def pull(self, idx):
        self.mn[idx] = min(self.mn[idx<<1], self.mn[idx<<1|1])
        self.mx[idx] = max(self.mx[idx<<1], self.mx[idx<<1|1])
    def add_range(self, idx, l, r, ql, qr, val):
        if ql > qr: return
        if ql <= l and r <= qr:
            self.apply(idx, val); return
        self.push(idx)
        mid = (l + r) >> 1
        if ql <= mid: self.add_range(idx<<1, l, mid, ql, min(qr, mid), val)
        if qr > mid:  self.add_range(idx<<1|1, mid+1, r, max(ql, mid+1), qr, val)
        self.pull(idx)
    def add(self, l, r, v):
        if l > r: return
        self.add_range(1, 0, self.n-1, l, r, v)
    def find_rightmost_zero(self, idx, l, r, ql, qr):
        if ql > qr or qr < l or ql > r: return -1
        if self.mn[idx] > 0 or self.mx[idx] < 0: return -1
        if l == r:
            return l if self.mn[idx] == 0 else -1
        self.push(idx)
        mid = (l + r) >> 1
        if qr > mid:
            res = self.find_rightmost_zero(idx<<1|1, mid+1, r, max(ql, mid+1), qr)
            if res != -1: return res
        if ql <= mid:
            return self.find_rightmost_zero(idx<<1, l, mid, ql, min(qr, mid))
        return -1
    def find(self, l, r):
        if l > r: return -1
        return self.find_rightmost_zero(1, 0, self.n-1, l, r)

class Solution:
    def longestBalanced(self, nums: List[int]) -> int:
        n = len(nums)
        pos = {}
        for i, v in enumerate(nums):
            pos.setdefault(v, []).append(i)
        st = SegTree(n)
        for v, lst in pos.items():
            sign = 1 if (v & 1) else -1
            st.add(lst[0], n-1, sign)
        ptr = {v:0 for v in pos}
        ans = 0
        for l in range(n):
            r = st.find(l, n-1)
            if r != -1:
                ans = max(ans, r - l + 1)
            x = nums[l]
            pIndex = ptr[x]; ptr[x] = pIndex + 1
            lst = pos[x]
            nextPos = lst[ptr[x]] if ptr[x] < len(lst) else n
            sign = 1 if (x & 1) else -1
            L, R = l, nextPos - 1
            if L <= R:
                st.add(L, R, -sign)
        return ans
