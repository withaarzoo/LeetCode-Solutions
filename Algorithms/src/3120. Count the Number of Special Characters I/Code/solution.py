class Solution:
    def numberOfSpecialChars(self, word: str) -> int:
        
        # Store all characters from the string
        st = set(word)

        # Variable to store answer
        count = 0

        # Check every lowercase letter
        for i in range(26):

            # Current lowercase character
            lower = chr(ord('a') + i)

            # Corresponding uppercase character
            upper = chr(ord('A') + i)

            # If both exist, it is a special character
            if lower in st and upper in st:
                count += 1

        # Return final answer
        return count