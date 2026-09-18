class Solution:
    def maxNumOfSubstrings(self, s: str) -> list[str]:
        n = len(s)
        # first and last occurrence of each letter
        left = [-1] * 26
        right = [-1] * 26
        for i, ch in enumerate(s):
            c = ord(ch) - 97
            if left[c] == -1:
                left[c] = i          # record first time we see it
            right[c] = i             # always update last time

        # collect every valid closed interval
        intervals = []
        for i in range(26):
            if left[i] == -1:
                continue             # letter never appeared
            L, R = left[i], right[i]
            valid = True
            # expand right end while scanning the current range
            j = L
            while j <= R:
                c = ord(s[j]) - 97
                if left[c] < L:      # needs to start earlier -> not closed
                    valid = False
                    break
                R = max(R, right[c]) # push right end if needed
                j += 1
            if valid:
                intervals.append((L, R))

        # sort by ending position so greedy can pick earliest-ending first
        intervals.sort(key=lambda x: x[1])

        # greedy selection of non-overlapping intervals
        ans = []
        last_end = -1
        for L, R in intervals:
            if L > last_end:         # completely after previous piece
                ans.append(s[L:R+1])
                last_end = R
        return ans