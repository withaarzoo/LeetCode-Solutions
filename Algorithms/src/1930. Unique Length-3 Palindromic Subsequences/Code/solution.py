class Solution:
    def countPalindromicSubsequence(self, s: str) -> int:
        first = [-1] * 26
        last = [-1] * 26
        
        for i, char in enumerate(s):
            index = ord(char) - ord('a')
            if first[index] == -1:
                first[index] = i
            last[index] = i
        
        result = 0
        for i in range(26):
            if first[i] != -1 and last[i] > first[i]:
                middle_chars = set(s[first[i] + 1:last[i]])
                result += len(middle_chars)
        
        return result
