class Solution:
    def minSumSquareDiff(self, nums1: list[int], nums2: list[int], k1: int, k2: int) -> int:
        n = len(nums1)
        d = [abs(a - b) for a, b in zip(nums1, nums2)]
        total = sum(d)
        k = k1 + k2
        if total <= k:
            return 0
        left, right = 0, max(d)
        while left < right:
            mid = (left + right) // 2
            need = sum(max(0, v - mid) for v in d)
            if need <= k:
                right = mid
            else:
                left = mid + 1
        for i in range(n):
            k -= max(0, d[i] - left)
            d[i] = min(d[i], left)
        for i in range(n):
            if k == 0:
                break
            if d[i] == left:
                d[i] -= 1
                k -= 1
        return sum(v * v for v in d)