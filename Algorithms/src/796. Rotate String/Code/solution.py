class Solution:
    def rotateString(self, s: str, goal: str) -> bool:
        # If lengths are different, rotation is impossible
        if len(s) != len(goal):
            return False

        # Concatenate s with itself
        doubled = s + s

        # Check if goal is inside doubled string
        return goal in doubled