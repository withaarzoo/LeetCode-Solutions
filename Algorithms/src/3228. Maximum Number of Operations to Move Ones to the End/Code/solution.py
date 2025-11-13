class Solution:
    def maxOperations(self, s: str) -> int:
        # ans and cnt use Python int (arbitrary precision)
        ans = 0
        cnt = 0  # number of '1's seen so far
        n = len(s)
        for i, c in enumerate(s):
            if c == '1':
                cnt += 1
            else:  # c == '0'
                # If this zero is immediately after a '1', it will be "used"
                # by each previously seen '1' across the optimal sequence.
                if i > 0 and s[i-1] == '1':
                    ans += cnt
        return ans
