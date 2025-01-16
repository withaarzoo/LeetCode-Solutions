class Solution:
    def xorAllNums(self, nums1: List[int], nums2: List[int]) -> int:
        xor1, xor2 = 0, 0
        
        # XOR all elements in nums1
        for num in nums1:
            xor1 ^= num
        
        # XOR all elements in nums2
        for num in nums2:
            xor2 ^= num
        
        # If nums1 has odd length, include xor2
        # If nums2 has odd length, include xor1
        return ((xor2 if len(nums1) % 2 else 0) ^ 
                (xor1 if len(nums2) % 2 else 0))
