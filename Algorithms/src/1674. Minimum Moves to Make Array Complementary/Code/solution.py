class Solution:
    def minMoves(self, nums: List[int], limit: int) -> int:
        
        n = len(nums)

        # Difference array
        diff = [0] * (2 * limit + 2)

        # Process every pair
        for i in range(n // 2):

            a = min(nums[i], nums[n - 1 - i])
            b = max(nums[i], nums[n - 1 - i])

            # Range where only 1 move is needed
            diff[a + 1] -= 1
            diff[b + limit + 1] += 1

            # Exact sum where 0 moves are needed
            diff[a + b] -= 1
            diff[a + b + 1] += 1

        pairs = n // 2

        # Initially every pair needs 2 moves
        current = pairs * 2

        answer = float('inf')

        # Prefix sum traversal
        for target_sum in range(2, 2 * limit + 1):

            current += diff[target_sum]

            answer = min(answer, current)

        return answer