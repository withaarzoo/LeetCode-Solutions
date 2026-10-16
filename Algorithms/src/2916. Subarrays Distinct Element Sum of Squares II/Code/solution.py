class Solution:
    def sumCounts(self, nums: List[int]) -> int:
        MOD = 10**9 + 7
        n = len(nums)
        sum_ = [0] * (4 * n)
        sumsq = [0] * (4 * n)
        lazy = [0] * (4 * n)

        def push(v, tl, tr):
            if lazy[v] != 0 and tl != tr:
                mid = (tl + tr) // 2
                llen = mid - tl + 1
                rlen = tr - mid
                add = lazy[v]

                sumsq[v * 2] = (sumsq[v * 2] + 2 * add * sum_[v * 2] + add * add % MOD * llen) % MOD
                sum_[v * 2] = (sum_[v * 2] + add * llen) % MOD
                lazy[v * 2] = (lazy[v * 2] + add) % MOD

                sumsq[v * 2 + 1] = (sumsq[v * 2 + 1] + 2 * add * sum_[v * 2 + 1] + add * add % MOD * rlen) % MOD
                sum_[v * 2 + 1] = (sum_[v * 2 + 1] + add * rlen) % MOD
                lazy[v * 2 + 1] = (lazy[v * 2 + 1] + add) % MOD

                lazy[v] = 0

        def update(v, tl, tr, l, r, add):
            if l > r:
                return
            if l == tl and r == tr:
                length = tr - tl + 1
                sumsq[v] = (sumsq[v] + 2 * add * sum_[v] + add * add % MOD * length) % MOD
                sum_[v] = (sum_[v] + add * length) % MOD
                lazy[v] = (lazy[v] + add) % MOD
                return
            push(v, tl, tr)
            mid = (tl + tr) // 2
            update(v * 2, tl, mid, l, min(r, mid), add)
            update(v * 2 + 1, mid + 1, tr, max(l, mid + 1), r, add)
            sum_[v] = (sum_[v * 2] + sum_[v * 2 + 1]) % MOD
            sumsq[v] = (sumsq[v * 2] + sumsq[v * 2 + 1]) % MOD

        def query(v, tl, tr, l, r):
            if l > r:
                return 0
            if l == tl and r == tr:
                return sumsq[v]
            push(v, tl, tr)
            mid = (tl + tr) // 2
            return (query(v * 2, tl, mid, l, min(r, mid)) +
                    query(v * 2 + 1, mid + 1, tr, max(l, mid + 1), r)) % MOD

        last = [-1] * 100001
        ans = 0
        for r in range(n):
            prv = last[nums[r]]
            if prv + 1 <= r - 1:
                update(1, 0, n - 1, prv + 1, r - 1, 1)
            update(1, 0, n - 1, r, r, 1)
            ans = (ans + query(1, 0, n - 1, 0, r)) % MOD
            last[nums[r]] = r
        return ans