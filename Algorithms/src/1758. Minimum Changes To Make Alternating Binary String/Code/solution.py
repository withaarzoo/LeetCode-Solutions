class Solution:
    def minOperations(self, s: str) -> int:
        startWith0 = 0
        startWith1 = 0

        for i in range(len(s)):

            expected0 = '0' if i % 2 == 0 else '1'
            expected1 = '1' if i % 2 == 0 else '0'

            if s[i] != expected0:
                startWith0 += 1

            if s[i] != expected1:
                startWith1 += 1

        return min(startWith0, startWith1)