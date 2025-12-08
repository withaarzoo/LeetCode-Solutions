import math

class Solution:
    def countTriples(self, n: int) -> int:
        count = 0
        
        # Try all possible pairs (a, b)
        for a in range(1, n + 1):
            for b in range(1, n + 1):
                sum_squares = a * a + b * b   # this should be c^2
                
                c = int(math.isqrt(sum_squares))  # integer square root (exact, no float)
                
                # Check if c is within range and forms a perfect square
                if c <= n and c * c == sum_squares:
                    count += 1  # (a, b, c) is a valid square triple
        
        return count
