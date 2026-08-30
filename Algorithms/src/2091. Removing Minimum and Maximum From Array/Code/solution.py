class Solution:
    def minimumDeletions(self, nums: List[int]) -> int:
        # Store the total number of elements in the array.
        n = len(nums)

        # Start by assuming the first element is both minimum and maximum.
        min_index = 0
        max_index = 0

        # Find the positions of the minimum and maximum elements.
        for i in range(1, n):
            # Update the minimum index if a smaller value is found.
            if nums[i] < nums[min_index]:
                min_index = i

            # Update the maximum index if a larger value is found.
            if nums[i] > nums[max_index]:
                max_index = i

        # Remove everything from the front up to the farther special element.
        remove_from_front = max(min_index, max_index) + 1

        # Remove everything from the back up to the farther special element.
        remove_from_back = n - min(min_index, max_index)

        # Calculate both ways of removing one element from each side.
        remove_from_both_sides = min(
            min_index + 1 + (n - max_index),
            max_index + 1 + (n - min_index)
        )

        # Return the minimum deletions among all possible strategies.
        return min(
            remove_from_front,
            remove_from_back,
            remove_from_both_sides
        )