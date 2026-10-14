class Solution:
    def maxSumOfSquares(self, num: int, sum: int) -> str:
        if sum > 9 * num:
            return ""
        q, r = divmod(sum, 9)
        return '9' * q + (str(r) if r else '') + '0' * (num - q - (1 if r else 0))