class Solution:
    def maximumProduct(self, nums: List[int]) -> int:

        # Store the three largest numbers
        max1 = float("-inf")
        max2 = float("-inf")
        max3 = float("-inf")

        # Store the two smallest numbers
        min1 = float("inf")
        min2 = float("inf")

        # Traverse the array once
        for num in nums:

            # Update the three largest numbers
            if num >= max1:
                max3 = max2
                max2 = max1
                max1 = num
            elif num >= max2:
                max3 = max2
                max2 = num
            elif num >= max3:
                max3 = num

            # Update the two smallest numbers
            if num <= min1:
                min2 = min1
                min1 = num
            elif num <= min2:
                min2 = num

        # Product of the three largest numbers
        product1 = max1 * max2 * max3

        # Product of the two smallest numbers and the largest number
        product2 = min1 * min2 * max1

        # Return the larger product
        return max(product1, product2)