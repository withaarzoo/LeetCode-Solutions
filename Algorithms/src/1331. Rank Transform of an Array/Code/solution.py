class Solution:
    def arrayRankTransform(self, arr: List[int]) -> List[int]:
        if not arr:
            return []
        
        # Step 1: Create a sorted copy of the array
        sorted_arr = sorted(arr)
        
        # Step 2: Create a dictionary to assign ranks
        rank_map = {}
        rank = 1
        
        # Step 3: Assign ranks to sorted elements
        for num in sorted_arr:
            if num not in rank_map:
                rank_map[num] = rank
                rank += 1
        
        # Step 4: Replace each element in the original array with its rank
        return [rank_map[num] for num in arr]