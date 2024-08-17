class Solution:
    def maxPoints(self, points: List[List[int]]) -> int:
        # Get the number of rows (m) and columns (n) from the points matrix
        m, n = len(points), len(points[0])
        
        # Initialize dp array to store maximum points for each column in the current row
        dp = [0] * n

        # Initialize dp with the first row's values
        for j in range(n):
            dp[j] = points[0][j]

        # Iterate through each row starting from the second row
        for i in range(1, m):
            # Arrays to store maximum values from left and right directions
            leftMax = [0] * n
            rightMax = [0] * n
            # Temporary array to store the new dp values for the current row
            newDp = [0] * n

            # Calculate leftMax array
            # leftMax[j] stores the maximum value we can get when moving left to right
            leftMax[0] = dp[0]  # The first element is the same as dp[0]
            for j in range(1, n):
                # Maximum of the current leftMax[j-1] or the value at dp[j] adjusted by column index
                leftMax[j] = max(leftMax[j - 1], dp[j] + j)

            # Calculate rightMax array
            # rightMax[j] stores the maximum value we can get when moving right to left
            rightMax[n - 1] = dp[n - 1] - (n - 1)  # The last element adjusted by its column index
            for j in range(n - 2, -1, -1):
                # Maximum of the current rightMax[j+1] or the value at dp[j] adjusted by column index
                rightMax[j] = max(rightMax[j + 1], dp[j] - j)

            # Calculate new dp values for the current row
            for j in range(n):
                # The new dp[j] is the maximum value we can obtain either from leftMax or rightMax adjusted by column index
                newDp[j] = max(leftMax[j] - j, rightMax[j] + j) + points[i][j]

            # Update dp to the newly calculated dp values
            dp = newDp

        # Return the maximum value from the dp array, which represents the maximum points possible
        return max(dp)
