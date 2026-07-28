class Solution:
    def smallestPalindrome(self, s: str) -> str:

        # Store frequency of every lowercase letter
        freq = [0] * 26

        # Count character frequencies
        for ch in s:
            freq[ord(ch) - ord('a')] += 1

        left = []
        middle = ""

        # Build the left half and find the middle character
        for i in range(26):

            # Add half of the occurrences to the left side
            left.append(chr(ord('a') + i) * (freq[i] // 2))

            # Odd frequency character becomes the center
            if freq[i] % 2:
                middle = chr(ord('a') + i)

        # Convert list into string
        left = "".join(left)

        # Right half is the reverse of the left half
        right = left[::-1]

        # Return the complete palindrome
        return left + middle + right