class Solution:
    def maxTotalValue(self, nums: List[int], k: int) -> int:
        n = len(nums)

        # floor(log2(i))
        lg = [0] * (n + 1)
        for i in range(2, n + 1):
            lg[i] = lg[i // 2] + 1

        K = lg[n] + 1

        # Sparse table for maximums
        mx = [[0] * n for _ in range(K)]

        # Sparse table for minimums
        mn = [[0] * n for _ in range(K)]

        for i in range(n):
            mx[0][i] = nums[i]
            mn[0][i] = nums[i]

        # Build sparse tables
        for j in range(1, K):
            length = 1 << j

            for i in range(n - length + 1):
                mx[j][i] = max(
                    mx[j - 1][i],
                    mx[j - 1][i + (length >> 1)]
                )

                mn[j][i] = min(
                    mn[j - 1][i],
                    mn[j - 1][i + (length >> 1)]
                )

        # O(1) range value query
        def get_value(l, r):
            length = r - l + 1
            p = lg[length]

            mx_val = max(
                mx[p][l],
                mx[p][r - (1 << p) + 1]
            )

            mn_val = min(
                mn[p][l],
                mn[p][r - (1 << p) + 1]
            )

            return mx_val - mn_val

        import heapq

        # Python heap is min-heap, so store negative values
        pq = []

        for l in range(n):
            heapq.heappush(
                pq,
                (-get_value(l, n - 1), l, n - 1)
            )

        ans = 0

        for _ in range(k):
            neg_val, l, r = heapq.heappop(pq)

            val = -neg_val
            ans += val

            if r > l:
                heapq.heappush(
                    pq,
                    (-get_value(l, r - 1), l, r - 1)
                )

        return ans