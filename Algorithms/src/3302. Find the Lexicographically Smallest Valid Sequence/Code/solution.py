from typing import List

class Solution:
    def validSequence(self, word1: str, word2: str) -> List[int]:
        n = len(word1)  # I store the length of word1 for the two linear scans.
        m = len(word2)  # I store the length of word2 because the answer has m indices.

        last = [-1] * m  # I store a position that can match every suffix of word2.

        i = n - 1  # I start from the end of word1.
        j = m - 1  # I start matching from the end of word2.

        # I greedily match word2 from right to left.
        # This gives me positions that can handle the suffix after a mismatch.
        while i >= 0 and j >= 0:
            # If the current characters match, this index can represent word2[j].
            if word1[i] == word2[j]:
                last[j] = i  # I remember this position for the suffix.
                j -= 1       # I now need to match the previous character.

            i -= 1  # I continue searching toward the beginning of word1.

        ans = []  # I store the final lexicographically smallest sequence.
        can_skip = True  # I have not used the one allowed mismatch yet.
        j = 0             # I start matching word2 from its first character.

        # I scan from left to right so that I always choose the earliest possible index.
        for i in range(n):
            # Once every character of word2 is matched, I do not need more indices.
            if j == m:
                break

            # A matching character can be selected without using the mismatch.
            if word1[i] == word2[j]:
                ans.append(i)  # I take this earliest valid index.
                j += 1         # I move to the next character of word2.

            # Otherwise, I try to spend the one allowed mismatch.
            elif can_skip and (j == m - 1 or i < last[j + 1]):
                can_skip = False  # I use the only allowed mismatch.
                ans.append(i)     # I choose this earliest possible index.
                j += 1            # I move to the next character of word2.

        # I return the sequence only if every character of word2 was matched.
        if j == m:
            return ans

        # If I could not match all of word2, no valid sequence exists.
        return []