class Solution:
    def maximumLengthSubstring(self, s: str) -> int:
        # Store the number of times each lowercase letter appears
        # inside the current sliding window.
        freq = [0] * 26

        # left marks the start of the current window.
        left = 0

        # ans stores the maximum valid window length found so far.
        ans = 0

        # Expand the window one character at a time.
        for right in range(len(s)):
            # Convert the current character into an index from 0 to 25
            # and increase its frequency in the window.
            index = ord(s[right]) - ord('a')
            freq[index] += 1

            # If this character appears more than two times,
            # shrink the window until it appears at most twice.
            while freq[index] > 2:
                # Remove the character leaving the window.
                freq[ord(s[left]) - ord('a')] -= 1

                # Move the left pointer forward.
                left += 1

            # The current window is valid, so update the maximum length.
            ans = max(ans, right - left + 1)

        # Return the longest valid substring length.
        return ans