class Solution:
    def shortestBeautifulSubstring(self, s: str, k: int) -> str:
        answer = ""  # I store the best beautiful substring found so far.
        left = 0     # I keep the left boundary of the sliding window.
        ones = 0     # I count how many '1' characters are inside the window.

        # I expand the window by moving the right pointer through the string.
        for right in range(len(s)):
            # I update the count when the newly added character is '1'.
            if s[right] == '1':
                ones += 1

            # If I have too many ones, I shrink the window from the left
            # until it contains at most k ones again.
            while ones > k:
                if s[left] == '1':
                    ones -= 1
                left += 1

            # I remove leading zeros because they do not affect the number
            # of ones and only make the current valid substring longer.
            while ones == k and s[left] == '0':
                left += 1

            # The current window is beautiful when it contains exactly k ones.
            if ones == k:
                # I create the current shortest candidate ending at right.
                candidate = s[left:right + 1]

                # I update the answer when this candidate is shorter, or when
                # equal-length candidates need lexicographical comparison.
                if (
                    not answer
                    or len(candidate) < len(answer)
                    or (len(candidate) == len(answer) and candidate < answer)
                ):
                    answer = candidate

        # I return the best substring, or an empty string if none was found.
        return answer