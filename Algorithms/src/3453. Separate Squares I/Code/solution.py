class Solution:
    def separateSquares(self, squares):
        totalArea = 0.0
        low, high = 1e18, -1e18

        for x, y, l in squares:
            totalArea += l * l
            low = min(low, y)
            high = max(high, y + l)

        for _ in range(80):
            mid = (low + high) / 2
            areaBelow = 0.0

            for x, y, l in squares:
                if mid <= y:
                    continue
                elif mid >= y + l:
                    areaBelow += l * l
                else:
                    areaBelow += l * (mid - y)

            if areaBelow * 2 < totalArea:
                low = mid
            else:
                high = mid

        return low
