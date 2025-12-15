class Solution:
    def getDescentPeriods(self, prices):
        ans = 1   # first day
        length = 1  # current smooth descent length

        for i in range(1, len(prices)):
            if prices[i] == prices[i - 1] - 1:
                length += 1
            else:
                length = 1
            ans += length

        return ans
