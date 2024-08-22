class Solution:
    def findComplement(self, num: int) -> int:
        # Initialize the mask as 0. This mask will be used to create a number 
        # with all bits set to 1 that is the same length as the binary representation of `num`.
        mask = 0
        
        # Create a temporary variable `temp` initialized to `num`.
        # This will help us determine the length of the binary representation of `num`.
        temp = num
        
        # Loop until `temp` becomes 0. Each iteration will shift `mask` to the left by one bit
        # and set the least significant bit of `mask` to 1.
        while temp != 0:
            # Left shift `mask` by 1 to make space for the next bit and then set the last bit to 1.
            mask = (mask << 1) | 1
            # Right shift `temp` by 1 to move to the next bit.
            temp >>= 1
        
        # XOR `num` with `mask` to flip all the bits of `num`.
        # This operation effectively finds the complement of the binary representation of `num`.
        return num ^ mask
