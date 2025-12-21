class Solution:
    def minDeletionSize(self, strs):
        n = len(strs)
        m = len(strs[0])
        
        sorted_rows = [False] * (n - 1)
        deletions = 0

        for col in range(m):
            need_delete = False

            for row in range(n - 1):
                if not sorted_rows[row] and strs[row][col] > strs[row + 1][col]:
                    need_delete = True
                    break

            if need_delete:
                deletions += 1
                continue

            for row in range(n - 1):
                if not sorted_rows[row] and strs[row][col] < strs[row + 1][col]:
                    sorted_rows[row] = True

        return deletions
