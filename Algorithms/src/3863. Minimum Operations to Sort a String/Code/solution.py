class Solution:
    def minOperations(self, s: str) -> int:
        n = len(s)
        if all(s[i] <= s[i + 1] for i in range(n - 1)):
            return 0
        if n == 2:
            return -1
        mn = min(s)
        mx = max(s)
        if s[0] == mn or s[-1] == mx:
            return 1
        if any(c == mn or c == mx for c in s[1:-1]):
            return 2
        return 3