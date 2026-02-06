class Solution:
    def minimumDeletions(self, s: str) -> int:
        countB = 0
        deletions = 0

        for ch in s:
            if ch == 'b':
                countB += 1
            else:
                deletions = min(deletions + 1, countB)

        return deletions
