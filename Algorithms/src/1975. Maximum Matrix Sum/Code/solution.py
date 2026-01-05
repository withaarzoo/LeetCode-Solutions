class Solution:
    def maxMatrixSum(self, matrix):
        total_sum = 0
        negative_count = 0
        min_abs = float('inf')

        for row in matrix:
            for val in row:
                total_sum += abs(val)      # add absolute value
                if val < 0:
                    negative_count += 1
                min_abs = min(min_abs, abs(val))

        if negative_count % 2 == 1:
            total_sum -= 2 * min_abs

        return total_sum
