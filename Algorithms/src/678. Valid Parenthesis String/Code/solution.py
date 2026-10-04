class Solution:
    def checkValidString(self, s: str) -> bool:
        low = high = 0
        for c in s:
            if c == '(':
                low += 1
                high += 1
            elif c == ')':
                low -= 1
                high -= 1
            else:
                low -= 1
                high += 1
            if high < 0:
                return False
            if low < 0:
                low = 0
        return low == 0