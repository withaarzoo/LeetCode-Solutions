class Solution:
    def makeFancyString(self, s: str) -> str:
        result = []
        for c in s:
            if len(result) < 2 or not (result[-1] == c and result[-2] == c):
                result.append(c)
        return ''.join(result)