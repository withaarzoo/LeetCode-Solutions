class Solution:
    def numWaterBottles(self, numBottles: int, numExchange: int) -> int:
        total = numBottles
        empties = numBottles

        while empties >= numExchange:
            new_full = empties // numExchange
            total += new_full
            empties = new_full + (empties % numExchange)

        return total
