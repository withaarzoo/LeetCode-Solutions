class Solution:
    def arrayRankTransform(self, arr: List[int]) -> List[int]:
        # Create a sorted copy of the array
        sorted_arr = sorted(arr)

        # Store value -> rank
        rank = {}
        current_rank = 1

        # Assign ranks only to unique values
        for num in sorted_arr:
            if num not in rank:
                rank[num] = current_rank
                current_rank += 1

        # Replace every element with its assigned rank
        return [rank[num] for num in arr]