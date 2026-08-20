class Solution:
    def resultArray(self, nums: List[int]) -> List[int]:
        # I put the first element into arr1 because the first operation is fixed.
        arr1 = [nums[0]]

        # I put the second element into arr2 because the second operation is fixed.
        arr2 = [nums[1]]

        # I process every remaining element one by one.
        for i in range(2, len(nums)):
            # If arr1 ends with a larger value, I add the current number to arr1.
            if arr1[-1] > arr2[-1]:
                arr1.append(nums[i])
            else:
                # Otherwise, I add the current number to arr2.
                arr2.append(nums[i])

        # I concatenate arr1 and arr2 because this is the required final order.
        return arr1 + arr2