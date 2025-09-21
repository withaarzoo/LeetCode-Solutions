import heapq
from collections import defaultdict
from typing import List, Tuple

class MovieRentingSystem:
    def __init__(self, n: int, entries: List[List[int]]):
        # price[(shop,movie)] = price
        self.price = {}
        # version[(shop,movie)] increments each time the (shop,movie) state changes
        self.version = {}
        # rentedState[(shop,movie)] = True if currently rented, False otherwise
        self.rented_state = {}

        # available heaps per movie: movie -> list of (price, shop, version)
        self.avail = defaultdict(list)
        # global rented heap: list of (price, shop, movie, version)
        self.rented = []

        for shop, movie, p in entries:
            key = (shop, movie)
            self.price[key] = p
            self.version[key] = 0
            self.rented_state[key] = False
            heapq.heappush(self.avail[movie], (p, shop, 0))

    def search(self, movie: int) -> List[int]:
        res = []
        tmp = []
        heap = self.avail.get(movie, [])
        # collect up to 5 valid shops
        while len(res) < 5 and heap:
            p, shop, ver = heapq.heappop(heap)
            key = (shop, movie)
            # valid if version matches and not rented
            if self.version.get(key, 0) == ver and not self.rented_state.get(key, False):
                res.append(shop)
                tmp.append((p, shop, ver))
            # else stale, skip
        # push back valid popped items
        for item in tmp:
            heapq.heappush(heap, item)
        return res

    def rent(self, shop: int, movie: int) -> None:
        key = (shop, movie)
        # bump version so prior avail entries become stale
        self.version[key] = self.version.get(key, 0) + 1
        self.rented_state[key] = True
        p = self.price[key]
        heapq.heappush(self.rented, (p, shop, movie, self.version[key]))

    def drop(self, shop: int, movie: int) -> None:
        key = (shop, movie)
        self.version[key] = self.version.get(key, 0) + 1
        self.rented_state[key] = False
        p = self.price[key]
        heapq.heappush(self.avail[movie], (p, shop, self.version[key]))

    def report(self) -> List[List[int]]:
        res = []
        tmp = []
        while len(res) < 5 and self.rented:
            p, shop, movie, ver = heapq.heappop(self.rented)
            key = (shop, movie)
            # valid if the version matches and it is currently rented
            if self.version.get(key, 0) == ver and self.rented_state.get(key, False):
                res.append([shop, movie])
                tmp.append((p, shop, movie, ver))
            # else stale
        for item in tmp:
            heapq.heappush(self.rented, item)
        return res

# Usage:
# obj = MovieRentingSystem(n, entries)
# param_1 = obj.search(movie)
# obj.rent(shop, movie)
# obj.drop(shop, movie)
# param_4 = obj.report()
