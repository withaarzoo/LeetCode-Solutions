class Solution:
    def pivotArray(self, nums: List[int], pivot: int) -> List[int]:

        # Store elements smaller than pivot
        smaller = []

        # Store elements equal to pivot
        equal = []

        # Store elements greater than pivot
        greater = []

        # Classify each element
        for num in nums:
            if num < pivot:
                smaller.append(num)
            elif num == pivot:
                equal.append(num)
            else:
                greater.append(num)

        # Return all groups in required order
        return smaller + equal + greater