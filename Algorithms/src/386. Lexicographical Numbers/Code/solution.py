class Solution:
    def lexicalOrder(self, n: int) -> List[int]:
        """
        This function returns the numbers from 1 to n in lexicographical order.
        
        :param n: The upper bound of the range of numbers
        :return: A list of integers from 1 to n in lexicographical order
        """
        result = []  # This list will store the result of numbers in lexicographical order
        
        # Iterate over numbers 1 to 9 as potential starting points
        # These will serve as the root for DFS exploration
        for i in range(1, 10):
            self.dfs(i, n, result)
        
        return result  # Return the final list of numbers in lexicographical order
    
    def dfs(self, curr: int, n: int, result: List[int]):
        """
        This is a helper function that performs depth-first search (DFS) to generate numbers.
        
        :param curr: The current number being processed
        :param n: The upper bound for the number range
        :param result: The list where valid numbers will be added
        """
        # Base case: If the current number exceeds n, we stop further exploration
        if curr > n:
            return  # No need to proceed as curr is greater than n
        
        # Append the current number to the result list as it is valid (<= n)
        result.append(curr)
        
        # Try to generate new numbers by appending digits 0-9 to the current number
        # We explore deeper in DFS by creating numbers like curr0, curr1, curr2, etc.
        for i in range(10):
            new_number = curr * 10 + i  # Generate a new number by adding a digit to curr
            # If the new number exceeds n, no need to proceed further in this branch
            if new_number > n:
                return  # Stop recursion for this branch
            # Recursively call DFS on the new number
            self.dfs(new_number, n, result)
