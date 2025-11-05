from typing import List
import heapq

class Solution:
    def findXSum(self, nums: List[int], k: int, x: int) -> List[int]:
        n = len(nums)
        ans = [0] * (n - k + 1)

        cnt = {}                 # value -> frequency in current window
        chosen = set()           # values currently in TOP

        # hot: min-heap of (freq, value) — worst among TOP at top
        hot: list[tuple[int, int]] = []
        # pool: max-heap of (freq, value) — best candidate at top (store as negatives)
        pool: list[tuple[int, int]] = []

        total = 0                # sum over TOP of value * freq

        def clean():
            # Drop stale entries sitting at the heap tops
            while hot and (hot[0][1] not in chosen or cnt.get(hot[0][1], 0) != hot[0][0]):
                heapq.heappop(hot)
            while pool and ((-pool[0][1]) in chosen or cnt.get(-pool[0][1], 0) != -pool[0][0] or -pool[0][0] == 0):
                heapq.heappop(pool)

        def demote_if_chosen(v: int):
            nonlocal total
            if v in chosen:
                # remove its current contribution using OLD freq
                chosen.remove(v)
                total -= v * cnt.get(v, 0)

        def promote_if_needed():
            nonlocal total
            clean()
            while len(chosen) < x and pool:
                f, v = -pool[0][0], -pool[0][1]
                if cnt.get(v, 0) != f or v in chosen or f == 0:
                    heapq.heappop(pool)
                    continue
                heapq.heappop(pool)
                chosen.add(v)
                total += v * f
                heapq.heappush(hot, (f, v))
            clean()

        def add_one(v: int):
            nonlocal total
            demote_if_chosen(v)                 # remove old contribution if it was in TOP
            f = cnt.get(v, 0) + 1
            cnt[v] = f
            heapq.heappush(pool, (-f, -v))      # new candidate goes to pool

            if len(chosen) < x:
                promote_if_needed()
            else:
                clean()
                if pool and hot:
                    bf, bv = -pool[0][0], -pool[0][1]  # best from REST
                    wf, wv = hot[0]                    # worst in TOP
                    if bf > wf or (bf == wf and bv > wv):
                        # promote best
                        heapq.heappop(pool)
                        chosen.add(bv)
                        total += bv * bf
                        heapq.heappush(hot, (bf, bv))

                        # demote worst
                        heapq.heappop(hot)
                        if wv in chosen:
                            chosen.remove(wv)
                            total -= wv * wf
                        heapq.heappush(pool, (-wf, -wv))
                clean()

        def remove_one(v: int):
            nonlocal total
            demote_if_chosen(v)                 # remove old contribution if it was in TOP
            f = cnt.get(v, 0) - 1
            if f <= 0:
                cnt.pop(v, None)
            else:
                cnt[v] = f
                heapq.heappush(pool, (-f, -v))  # rank got worse → pool
            promote_if_needed()

        # build first window
        for i in range(k):
            add_one(nums[i])
        ans[0] = total

        # slide
        for i in range(k, n):
            remove_one(nums[i - k])
            add_one(nums[i])
            ans[i - k + 1] = total

        return ans
