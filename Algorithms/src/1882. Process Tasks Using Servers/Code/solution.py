class Solution:
    def assignTasks(self, servers: list[int], tasks: list[int]) -> list[int]:
        import heapq
        free = [(w, i) for i, w in enumerate(servers)]
        heapq.heapify(free)
        busy = []
        ans = [0] * len(tasks)
        for j, t in enumerate(tasks):
            while busy and busy[0][0] <= j:
                _, w, i = heapq.heappop(busy)
                heapq.heappush(free, (w, i))
            if free:
                w, i = heapq.heappop(free)
                ans[j] = i
                heapq.heappush(busy, (j + t, w, i))
            else:
                ft, w, i = heapq.heappop(busy)
                ans[j] = i
                heapq.heappush(busy, (ft + t, w, i))
        return ans