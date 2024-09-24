class Solution:
    def longestCommonPrefix(self, arr1: list[int], arr2: list[int]) -> int:
        # Dictionary to store prefixes from arr1 and their occurrences
        prefix_map = {}

        # Step 1: Build the prefix map using arr1
        # Loop through each number in arr1
        for num in arr1:
            # Convert the number to a string to easily extract prefixes
            str_num = str(num)
            # Initialize an empty prefix string
            prefix = ""
            
            # Loop through each character in the string version of the number
            for ch in str_num:
                # Append the character to the prefix
                prefix += ch
                # Update the prefix count in the prefix_map
                # If the prefix already exists, increment its count; otherwise, set it to 1
                prefix_map[prefix] = prefix_map.get(prefix, 0) + 1
        
        # Variable to track the maximum length of a common prefix found
        max_length = 0

        # Step 2: Check for common prefixes in arr2
        # Loop through each number in arr2
        for num in arr2:
            # Convert the number to a string to extract prefixes
            str_num = str(num)
            # Initialize an empty prefix string for the current number in arr2
            prefix = ""
            
            # Loop through each character in the string version of the number
            for ch in str_num:
                # Append the character to the prefix
                prefix += ch
                # If the prefix exists in prefix_map (i.e., it was found in arr1)
                if prefix in prefix_map:
                    # Update the maximum length of the common prefix found
                    max_length = max(max_length, len(prefix))
        
        # Return the maximum length of the common prefix between arr1 and arr2
        return max_length
