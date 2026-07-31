class Solution:
    def minimumPushes(self, word: str) -> int:
        # Store the frequency of every lowercase letter
        freq = [0] * 26

        # Count each character
        for ch in word:
            freq[ord(ch) - ord('a')] += 1

        # Sort frequencies from largest to smallest
        freq.sort(reverse=True)

        ans = 0

        # Assign push cost according to position
        for i in range(26):
            # Ignore letters that never appear
            if freq[i] == 0:
                break

            # Every 8 letters increase the push count
            pushes = (i // 8) + 1

            # Add contribution
            ans += freq[i] * pushes

        return ans