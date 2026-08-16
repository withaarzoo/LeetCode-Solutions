from typing import List

class Solution:
    def stoneGameIX(self, stones: List[int]) -> bool:
        # cnt[r] stores how many stones have remainder r modulo 3.
        cnt = [0, 0, 0]

        # Count the stones in each remainder group.
        for stone in stones:
            # Only the remainder matters for deciding the winner.
            cnt[stone % 3] += 1

        # With an even number of remainder-0 stones,
        # Alice needs both a remainder-1 and a remainder-2 stone.
        if cnt[0] % 2 == 0:
            return cnt[1] > 0 and cnt[2] > 0

        # With an odd number of remainder-0 stones,
        # Alice wins when the two useful groups differ by more than 2.
        return abs(cnt[1] - cnt[2]) > 2