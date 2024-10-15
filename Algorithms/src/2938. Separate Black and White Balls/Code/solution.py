class Solution:
    def minimumSteps(self, s: str) -> int:
        ans = 0
        blackCount = 0 # Tracks the number of black balls (1s)

        # Traverse through the string
        for ch in s:
            if ch == '0':
                # White ball encountered, add the number of black balls on its left
                ans += blackCount
            else:
                # Black ball encountered, increment the black ball count
                blackCount += 1

        return ans
