class Solution:
    def maxSideLength(self, mat, threshold):
        m, n = len(mat), len(mat[0])

        # Prefix sum matrix
        pre = [[0] * (n + 1) for _ in range(m + 1)]

        for i in range(1, m + 1):
            for j in range(1, n + 1):
                pre[i][j] = mat[i - 1][j - 1] \
                            + pre[i - 1][j] \
                            + pre[i][j - 1] \
                            - pre[i - 1][j - 1]

        left, right = 0, min(m, n)
        ans = 0

        while left <= right:
            mid = (left + right) // 2
            found = False

            for i in range(mid, m + 1):
                for j in range(mid, n + 1):
                    square_sum = pre[i][j] \
                               - pre[i - mid][j] \
                               - pre[i][j - mid] \
                               + pre[i - mid][j - mid]

                    if square_sum <= threshold:
                        found = True
                        break
                if found:
                    break

            if found:
                ans = mid
                left = mid + 1
            else:
                right = mid - 1

        return ans
