class Solution:
    def findKthBit(self, n: int, k: int) -> str:
        # Base case: When n = 1, the binary string is "0"
        if n == 1:
            return '0'
        
        # Find the length of the current string Sn, which is 2^n - 1
        length = (1 << n) - 1
        
        # Find the middle position
        mid = length // 2 + 1
        
        # If k is the middle position, return '1'
        if k == mid:
            return '1'
        
        # If k is in the first half, find the bit in Sn-1
        if k < mid:
            return self.findKthBit(n - 1, k)
        
        # If k is in the second half, find the bit in Sn-1 and invert it
        return '1' if self.findKthBit(n - 1, length - k + 1) == '0' else '0'
