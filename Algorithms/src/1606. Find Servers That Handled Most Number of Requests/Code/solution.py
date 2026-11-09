from sortedcontainers import SortedList
import heapq

class Solution:
    def busiestServers(self, k: int, arrival: list[int], load: list[int]) -> list[int]:
        free = SortedList(range(k))
        busy = []
        cnt = [0] * k
        n = len(arrival)
        for i in range(n):
            t = arrival[i]
            while busy and busy[0][0] <= t:
                free.add(heapq.heappop(busy)[1])
            if not free:
                continue
            pref = i % k
            idx = free.bisect_left(pref)
            if idx == len(free):
                idx = 0
            s = free[idx]
            free.pop(idx)
            cnt[s] += 1
            heapq.heappush(busy, (t + load[i], s))
        mx = max(cnt)
        return [i for i in range(k) if cnt[i] == mx]