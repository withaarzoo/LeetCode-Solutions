class Solution:
    def minNumberOfSeconds(self, mountainHeight: int, workerTimes: List[int]) -> int:

        def can(time):
            total = 0

            for t in workerTimes:

                left, right = 0, mountainHeight

                while left <= right:

                    mid = (left + right) // 2

                    required = t * (mid * (mid + 1) // 2)

                    if required <= time:
                        left = mid + 1
                    else:
                        right = mid - 1

                total += right

                if total >= mountainHeight:
                    return True

            return False

        left, right = 1, 10**18
        ans = right

        while left <= right:

            mid = (left + right) // 2

            if can(mid):
                ans = mid
                right = mid - 1
            else:
                left = mid + 1

        return ans