from collections import deque, defaultdict
import bisect
from typing import List

class Router:
    def __init__(self, memoryLimit: int):
        self.memoryLimit = memoryLimit
        self.q = deque()                       # global FIFO queue of (s,d,t)
        self.seen = set()                      # set of (s,d,t) tuples for duplicates
        self.times = defaultdict(list)        # dest -> list of timestamps (append-only)
        self.head = defaultdict(int)          # dest -> head index (how many removed from front)

    def addPacket(self, source: int, destination: int, timestamp: int) -> bool:
        key = (source, destination, timestamp)
        if key in self.seen:
            return False

        # Evict until there is room (usually at most one)
        while len(self.q) >= self.memoryLimit:
            os, od, ot = self.q.popleft()
            self.seen.remove((os, od, ot))
            self.head[od] += 1

        # Add the new packet
        self.q.append((source, destination, timestamp))
        self.seen.add(key)
        self.times[destination].append(timestamp)
        return True

    def forwardPacket(self) -> List[int]:
        if not self.q:
            return []
        s, d, t = self.q.popleft()
        self.seen.remove((s, d, t))
        self.head[d] += 1
        return [s, d, t]

    def getCount(self, destination: int, startTime: int, endTime: int) -> int:
        arr = self.times.get(destination, [])
        h = self.head.get(destination, 0)
        # arr is sorted from appends; search only in arr[h:]
        L = bisect.bisect_left(arr, startTime, lo=h)
        R = bisect.bisect_right(arr, endTime, lo=h)
        return R - L

# Your Router object will be instantiated and called as such:
# obj = Router(memoryLimit)
# param_1 = obj.addPacket(source,destination,timestamp)
# param_2 = obj.forwardPacket()
# param_3 = obj.getCount(destination,startTime,endTime)
