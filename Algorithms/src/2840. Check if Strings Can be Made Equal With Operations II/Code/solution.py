class Solution:
    def checkStrings(self, s1: str, s2: str) -> bool:
        # Frequency arrays for even and odd positions
        even = [0] * 26
        odd = [0] * 26

        for i in range(len(s1)):
            if i % 2 == 0:
                # Count characters at even indexes
                even[ord(s1[i]) - ord('a')] += 1
                even[ord(s2[i]) - ord('a')] -= 1
            else:
                # Count characters at odd indexes
                odd[ord(s1[i]) - ord('a')] += 1
                odd[ord(s2[i]) - ord('a')] -= 1

        # Check if all frequencies become zero
        for i in range(26):
            if even[i] != 0 or odd[i] != 0:
                return False

        return True