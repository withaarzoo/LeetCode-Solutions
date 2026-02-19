class Solution:
    def countBinarySubstrings(self, s: str) -> int:
        
        prevGroup = 0
        currGroup = 1
        result = 0
        
        for i in range(1, len(s)):
            
            if s[i] == s[i - 1]:
                currGroup += 1
            else:
                result += min(prevGroup, currGroup)
                prevGroup = currGroup
                currGroup = 1
        
        result += min(prevGroup, currGroup)
        
        return result
