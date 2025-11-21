class Solution:
    def countPalindromicSubsequence(self, s: str) -> int:
        n = len(s)
        A = 26
        first = [10**9] * A
        last = [-1] * A
        # record first and last occurrence for each letter
        for i, ch in enumerate(s):
            idx = ord(ch) - ord('a')
            if i < first[idx]:
                first[idx] = i
            if i > last[idx]:
                last[idx] = i

        ans = 0
        # for each outer letter, count distinct middle letters between first and last
        for c in range(A):
            if first[c] < last[c]:
                seen = [False] * A
                for i in range(first[c] + 1, last[c]):
                    seen[ord(s[i]) - ord('a')] = True
                ans += sum(1 for x in seen if x)
        return ans
