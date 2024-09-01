class Solution:
    def construct2DArray(self, original: List[int], m: int, n: int) -> List[List[int]]:
        # Step 1: Check if the reshaping is possible.
        # The total number of elements in the original list should match the total number of elements
        # in the desired 2D array, which is m * n. If not, return an empty list.
        if len(original) != m * n:
            return []  # Reshaping is not possible, so return an empty list.

        # Step 2: Initialize the 2D array.
        # We will construct the 2D array by slicing the original list.
        # For each row in the 2D array (there are 'm' rows), we slice 'n' elements from the original list.
        result = []  # This will hold the 2D array.
        
        # Step 3: Fill the 2D array row by row.
        for i in range(m):
            # Calculate the start and end indices for the slice.
            start_index = i * n  # The starting index for the slice.
            end_index = (i + 1) * n  # The ending index for the slice (exclusive).
            
            # Slice the original list to get a row of 'n' elements and add it to the result.
            row = original[start_index:end_index]
            result.append(row)  # Add the row to the result.
        
        # Step 4: Return the constructed 2D array.
        return result
