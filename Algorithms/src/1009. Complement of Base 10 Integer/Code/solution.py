class Solution:
    def bitwiseComplement(self, n: int) -> int:
        
        # Edge case
        if n == 0:
            return 1

        mask = 0

        # Create mask with all bits set to 1
        while mask < n:
            mask = (mask << 1) | 1

        # XOR flips the bits
        return mask ^ n