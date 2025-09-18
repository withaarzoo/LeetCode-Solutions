import heapq
from typing import List

class TaskManager:

    def __init__(self, tasks: List[List[int]]):
        self.mp = {}           # taskId -> (userId, priority)
        self.heap = []         # holds (-priority, -taskId, taskId)
        for t in tasks:
            if len(t) < 3:
                continue
            user, task, pr = t[0], t[1], t[2]
            self.mp[task] = (user, pr)
            heapq.heappush(self.heap, (-pr, -task, task))

    def add(self, userId: int, taskId: int, priority: int) -> None:
        self.mp[taskId] = (userId, priority)
        heapq.heappush(self.heap, (-priority, -taskId, taskId))

    def edit(self, taskId: int, newPriority: int) -> None:
        # guaranteed taskId exists
        user = self.mp[taskId][0]
        self.mp[taskId] = (user, newPriority)
        heapq.heappush(self.heap, (-newPriority, -taskId, taskId))

    def rmv(self, taskId: int) -> None:
        # guaranteed taskId exists
        if taskId in self.mp:
            del self.mp[taskId]  # heap entry becomes stale

    def execTop(self) -> int:
        while self.heap:
            neg_pr, neg_tid, tid = heapq.heappop(self.heap)
            cur = self.mp.get(tid)
            if cur is None:
                continue  # removed
            if cur[1] != -neg_pr:
                continue  # stale priority
            userId = cur[0]
            del self.mp[tid]
            return userId
        return -1

# Your TaskManager object will be instantiated and called as such:
# obj = TaskManager(tasks)
# obj.add(userId,taskId,priority)
# obj.edit(taskId,newPriority)
# obj.rmv(taskId)
# param_4 = obj.execTop()
