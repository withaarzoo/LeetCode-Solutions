class Solution:
    def uniqueXorTriplets(self, nums: List[int]) -> int:
        # Length of the permutation
        n = len(nums)

        # Small cases
        if n <= 2:
            return n

        # bit_length() gives the number of bits needed to represent n
        return 1 << n.bit_length()