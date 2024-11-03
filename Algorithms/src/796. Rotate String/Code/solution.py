class Solution:
    def rotateString(self, s: str, goal: str) -> bool:
        # If lengths differ, s cannot be rotated to match goal
        if len(s) != len(goal):
            return False
        
        # Concatenate s with itself
        doubled = s + s
        
        # Check if goal is a substring of doubled
        return goal in doubled
