class Solution:
    def checkString(self, s: str) -> bool:
        seenB = False
        for c in s:
            if c == 'b':
                seenB = True
            elif seenB:
                return False
        return True