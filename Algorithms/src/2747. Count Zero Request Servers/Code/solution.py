class Solution:
    def countServers(self, n: int, logs: List[List[int]], x: int, queries: List[int]) -> List[int]:
        logs.sort(key=lambda a: a[1])
        m = len(queries)
        q = [(queries[i], i) for i in range(m)]
        q.sort()
        ans = [0] * m
        freq = {}
        active = 0
        left = right = 0
        L = len(logs)
        for time, idx in q:
            while right < L and logs[right][1] <= time:
                s = logs[right][0]
                freq[s] = freq.get(s, 0) + 1
                if freq[s] == 1:
                    active += 1
                right += 1
            while left < right and logs[left][1] < time - x:
                s = logs[left][0]
                freq[s] -= 1
                if freq[s] == 0:
                    active -= 1
                left += 1
            ans[idx] = n - active
        return ans