class Solution:
    def longestCommonPrefix(self, arr1: List[int], arr2: List[int]) -> int:
        
        # Set to store all prefixes from arr1
        prefixes = set()

        # Generate prefixes from arr1
        for num in arr1:

            x = num

            # Keep removing last digit
            while x > 0:

                # Store current prefix
                prefixes.add(x)

                # Remove last digit
                x //= 10

        ans = 0

        # Process arr2
        for num in arr2:

            x = num

            # Try all prefixes
            while x > 0:

                # Prefix exists
                if x in prefixes:

                    # Update answer with digit length
                    ans = max(ans, len(str(x)))

                    # Stop because longer prefix already found
                    break

                # Remove last digit
                x //= 10

        return ans