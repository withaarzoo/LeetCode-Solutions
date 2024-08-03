from typing import List

class Solution:
    def canBeEqual(self, target: List[int], arr: List[int]) -> bool:
        # Step 1: Sort the target array
        target.sort()
        
        # Step 2: Sort the arr array
        arr.sort()
        
        # Step 3: Compare the sorted target and arr arrays
        # If they are equal, it means that one can be rearranged to match the other
        return target == arr
