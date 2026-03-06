class Solution:
    def checkOnesSegment(self, s: str) -> bool:
        # If "01" exists, it means another segment of 1s started
        return "01" not in s