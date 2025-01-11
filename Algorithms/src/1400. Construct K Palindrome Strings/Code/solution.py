class Solution:
    def canConstruct(self, s: str, k: int) -> bool:
        if k > len(s):
            return False # More palindromes than characters
        freq = [0] * 26 # Frequency array for lowercase letters
        for char in s:
            freq[ord(char) - ord('a')] += 1
        odd_count = sum(1 for count in freq if count % 2 != 0)
        return odd_count <= k
