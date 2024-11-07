class Solution:
    def largestCombination(self, candidates: List[int]) -> int:
        bit_count = [0] * 31  # Array to count '1's at each bit position
        
        # Count '1's in each bit position across all numbers
        for num in candidates:
            for i in range(31):
                if num & (1 << i):
                    bit_count[i] += 1
                    
        # Find the maximum count in any bit position
        return max(bit_count)
