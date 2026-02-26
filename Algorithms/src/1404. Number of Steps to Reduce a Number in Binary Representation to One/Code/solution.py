class Solution:
    def numSteps(self, s: str) -> int:
        steps = 0
        carry = 0
        
        # Traverse from right to left (ignore first bit)
        for i in range(len(s) - 1, 0, -1):
            bit = int(s[i]) + carry
            
            if bit == 1:
                # Odd case
                steps += 2
                carry = 1
            else:
                # Even case
                steps += 1
        
        return steps + carry