class Solution:
    def maximumJumps(self, nums: List[int], target: int) -> int:
        
        n = len(nums)

        # dp[i] = maximum jumps needed to reach index i
        dp = [-1] * n

        # Starting index requires 0 jumps
        dp[0] = 0

        # Try every current index
        for i in range(n):

            # Skip unreachable indices
            if dp[i] == -1:
                continue

            # Try every next index
            for j in range(i + 1, n):

                # Difference between values
                diff = nums[j] - nums[i]

                # Check whether jump is valid
                if -target <= diff <= target:

                    # Update maximum jumps
                    dp[j] = max(dp[j], dp[i] + 1)

        # Final answer
        return dp[n - 1]