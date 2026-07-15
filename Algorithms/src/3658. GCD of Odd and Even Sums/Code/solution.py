class Solution:
    def gcdOfOddEvenSums(self, n: int) -> int:
        # From the mathematical proof:
        # GCD(n^2, n*(n+1)) = n
        return n