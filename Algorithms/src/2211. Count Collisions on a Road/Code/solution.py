class Solution:
    def countCollisions(self, directions: str) -> int:
        n = len(directions)
        i, j = 0, n - 1
        
        # Skip leading 'L' cars (safe)
        while i < n and directions[i] == 'L':
            i += 1
        
        # Skip trailing 'R' cars (safe)
        while j >= 0 and directions[j] == 'R':
            j -= 1
        
        collisions = 0
        # Count all moving cars ('L' or 'R') in the remaining middle part
        for k in range(i, j + 1):
            if directions[k] != 'S':
                collisions += 1
        
        return collisions
