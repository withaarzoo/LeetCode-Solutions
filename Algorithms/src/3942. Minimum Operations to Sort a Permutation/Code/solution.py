class Solution:
    def minOperations(self, nums: List[int]) -> int:
        n = len(nums)
        starts = []
        inc = dec = True
        for i in range(n):
            if nums[(i + 1) % n] != (nums[i] + 1) % n:
                inc = False
            if nums[(i + 1) % n] != (nums[i] - 1 + n) % n:
                dec = False
        if inc:
            starts.append((0, nums[0]))
        if dec:
            starts.append((1, (n - 1 - nums[0]) % n))
        if not starts:
            return -1

        dist = [[-1] * n for _ in range(2)]
        q = deque()
        for t, s in starts:
            dist[t][s] = 0
            q.append((t, s))
        while q:
            t, s = q.popleft()
            if t == 0 and s == 0:
                return dist[t][s]
            d = dist[t][s]
            nt, ns = t, (s + 1) % n
            if dist[nt][ns] == -1:
                dist[nt][ns] = d + 1
                q.append((nt, ns))
            nt, ns = 1 - t, (n - s) % n
            if dist[nt][ns] == -1:
                dist[nt][ns] = d + 1
                q.append((nt, ns))
        return -1