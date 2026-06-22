class Solution:
    def maxNumberOfBalloons(self, text: str) -> int:
        # Store frequency of all lowercase letters
        freq = [0] * 26

        # Count each character
        for ch in text:
            freq[ord(ch) - ord('a')] += 1

        # Return the smallest possible complete balloon count
        return min(
            freq[ord('b') - ord('a')],      # Need 1 'b'
            freq[ord('a') - ord('a')],      # Need 1 'a'
            freq[ord('l') - ord('a')] // 2, # Need 2 'l'
            freq[ord('o') - ord('a')] // 2, # Need 2 'o'
            freq[ord('n') - ord('a')]       # Need 1 'n'
        )