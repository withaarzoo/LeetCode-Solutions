class Solution:
    def maxBottlesDrunk(self, numBottles: int, numExchange: int) -> int:
        full = numBottles    # current full bottles
        empty = 0            # current empty bottles
        ans = 0              # total bottles drunk
        curEx = numExchange  # current exchange requirement

        while full > 0:
            # drink all full bottles I have
            ans += full
            empty += full
            full = 0

            # exchange empties for one full bottle at a time,
            # because curEx increases after each exchange
            while empty >= curEx:
                empty -= curEx
                full += 1
                curEx += 1

        return ans
