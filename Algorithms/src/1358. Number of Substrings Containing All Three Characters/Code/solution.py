class Solution:
    def numberOfSubstrings(self, s: str) -> int:

        # Store the frequency of 'a', 'b', and 'c'
        freq = [0] * 3

        left = 0
        ans = 0
        n = len(s)

        # Expand the window
        for right in range(n):

            # Add the current character
            freq[ord(s[right]) - ord('a')] += 1

            # Shrink while all three characters exist
            while freq[0] > 0 and freq[1] > 0 and freq[2] > 0:

                # Count every valid substring starting at 'left'
                ans += n - right

                # Remove the leftmost character
                freq[ord(s[left]) - ord('a')] -= 1

                # Move left forward
                left += 1

        # Return the total count
        return ans