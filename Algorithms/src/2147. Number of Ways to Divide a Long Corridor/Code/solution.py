class Solution:
    def numberOfWays(self, corridor: str) -> int:
        MOD = 10**9 + 7
        seats = []

        for i, c in enumerate(corridor):
            if c == 'S':
                seats.append(i)

        if len(seats) == 0 or len(seats) % 2 != 0:
            return 0

        ways = 1
        for i in range(2, len(seats), 2):
            ways = (ways * (seats[i] - seats[i - 1])) % MOD

        return ways
