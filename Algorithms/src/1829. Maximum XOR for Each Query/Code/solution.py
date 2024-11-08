class Solution:
    def getMaximumXor(self, nums: List[int], maximumBit: int) -> List[int]:
        n = len(nums)
        answer = [0] * n
        XORed = 0
        
        # Calculate the cumulative XOR of the entire nums array
        for num in nums:
            XORed ^= num
        
        # max_k is 2^maximumBit - 1
        max_k = (1 << maximumBit) - 1
        
        # Process each query in reverse
        for i in range(n):
            # Calculate the k that maximizes XOR
            answer[i] = XORed ^ max_k
            
            # Update XORed by removing the effect of the last element
            XORed ^= nums[n - 1 - i]
        
        return answer
