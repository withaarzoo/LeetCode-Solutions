class Solution:
    def predictTheWinner(self, nums: List[int]) -> bool:
        n = len(nums)

        # dp[i][j] stores the maximum score difference
        # the current player can achieve for subarray [i...j].
        dp = [[0] * n for _ in range(n)]

        # Base case:
        # One number is left, so the player takes it.
        for i in range(n):
            dp[i][i] = nums[i]

        # Fill the table from smaller ranges to larger ranges.
        for length in range(2, n + 1):
            for i in range(n - length + 1):
                j = i + length - 1

                # Pick the left number.
                take_left = nums[i] - dp[i + 1][j]

                # Pick the right number.
                take_right = nums[j] - dp[i][j - 1]

                # Store the better option.
                dp[i][j] = max(take_left, take_right)

        # Non-negative means Player 1 wins or ties.
        return dp[0][n - 1] >= 0