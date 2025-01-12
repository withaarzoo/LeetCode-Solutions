class Solution:
    def canBeValid(self, s: str, locked: str) -> bool:
        if len(s) % 2 != 0:
            return False  # Odd length can't be balanced
        
        open, flexible = 0, 0
        # Left-to-right pass
        for i in range(len(s)):
            if locked[i] == '1':
                open += 1 if s[i] == '(' else -1
            else:
                flexible += 1
            if open + flexible < 0:
                return False
        
        open, flexible = 0, 0
        # Right-to-left pass
        for i in range(len(s) - 1, -1, -1):
            if locked[i] == '1':
                open += 1 if s[i] == ')' else -1
            else:
                flexible += 1
            if open + flexible < 0:
                return False
        
        return True
