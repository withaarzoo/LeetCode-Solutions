from typing import List

class Solution:
    def sortByBits(self, arr: List[int]) -> List[int]:
        
        # Sort using custom key
        return sorted(arr, key=lambda x: (bin(x).count('1'), x))