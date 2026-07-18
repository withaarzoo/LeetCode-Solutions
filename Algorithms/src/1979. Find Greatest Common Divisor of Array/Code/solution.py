class Solution:
    # Function to find GCD using the Euclidean Algorithm
    def gcd(self, a, b):
        # Continue until b becomes 0
        while b:
            a, b = b, a % b

        # a now stores the GCD
        return a

    def findGCD(self, nums: List[int]) -> int:

        # Find the smallest element
        minimum = min(nums)

        # Find the largest element
        maximum = max(nums)

        # Return the GCD of the smallest and largest values
        return self.gcd(minimum, maximum)