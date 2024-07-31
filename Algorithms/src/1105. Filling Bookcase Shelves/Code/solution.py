from typing import List

class Solution:
    def minHeightShelves(self, books: List[List[int]], shelfWidth: int) -> int:
        n = len(books)  # Total number of books
        dp = [float('inf')] * (n + 1)  # Initialize DP array with infinity
        dp[0] = 0  # Base case: No books, no height required
        
        # Iterate over each book
        for i in range(1, n + 1):
            width = 0  # Initialize the current shelf's width
            height = 0  # Initialize the current shelf's max height
            
            # Try placing books on the current shelf starting from book i to 1
            for j in range(i, 0, -1):
                width += books[j - 1][0]  # Add the width of the current book
                if width > shelfWidth:
                    break  # If the shelf width is exceeded, break the loop
                
                height = max(height, books[j - 1][1])  # Update the shelf height to the tallest book
                
                # Update the minimum height by comparing the current and previous configurations
                dp[i] = min(dp[i], dp[j - 1] + height)
        
        # The last element of dp array will hold the minimum height for all books
        return dp[n]
