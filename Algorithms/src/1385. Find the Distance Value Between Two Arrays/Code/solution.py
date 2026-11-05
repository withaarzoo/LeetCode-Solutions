class Solution:
    def findTheDistanceValue(self, arr1: list[int], arr2: list[int], d: int) -> int:
        arr2.sort()
        ans = 0
        for x in arr1:
            left, right = 0, len(arr2)
            while left < right:
                mid = (left + right) // 2
                if arr2[mid] < x:
                    left = mid + 1
                else:
                    right = mid
            ok = True
            if left < len(arr2) and abs(arr2[left] - x) <= d:
                ok = False
            if left > 0 and abs(arr2[left - 1] - x) <= d:
                ok = False
            if ok:
                ans += 1
        return ans