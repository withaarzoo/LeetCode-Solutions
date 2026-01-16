class Solution:
    def maximizeSquareArea(self, m: int, n: int, hFences: list[int], vFences: list[int]) -> int:
        MOD = 10**9 + 7

        hFences.append(1)
        hFences.append(m)
        vFences.append(1)
        vFences.append(n)

        hFences.sort()
        vFences.sort()

        horizontal = set()
        vertical = set()

        for i in range(len(hFences)):
            for j in range(i + 1, len(hFences)):
                horizontal.add(hFences[j] - hFences[i])

        for i in range(len(vFences)):
            for j in range(i + 1, len(vFences)):
                vertical.add(vFences[j] - vFences[i])

        maxSide = 0
        for d in horizontal:
            if d in vertical:
                maxSide = max(maxSide, d)

        if maxSide == 0:
            return -1

        return (maxSide * maxSide) % MOD
