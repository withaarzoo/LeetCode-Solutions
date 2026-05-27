class Solution:
    def numberOfSpecialChars(self, word: str) -> int:
        
        # Store last occurrence of lowercase letters
        lower = [-1] * 26

        # Store first occurrence of uppercase letters
        upper = [-1] * 26

        # Traverse the string
        for i, ch in enumerate(word):

            # If lowercase letter
            if 'a' <= ch <= 'z':

                # Update last occurrence
                lower[ord(ch) - ord('a')] = i

            else:
                idx = ord(ch) - ord('A')

                # Store only first occurrence
                if upper[idx] == -1:
                    upper[idx] = i

        ans = 0

        # Check all 26 letters
        for i in range(26):

            # Both lowercase and uppercase must exist
            if lower[i] != -1 and upper[i] != -1:

                # Lowercase must come before uppercase
                if lower[i] < upper[i]:
                    ans += 1

        return ans