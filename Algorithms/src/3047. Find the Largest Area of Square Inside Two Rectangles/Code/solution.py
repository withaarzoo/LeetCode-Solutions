class Solution:
    def largestSquareArea(self, bottomLeft, topRight):
        n = len(bottomLeft)
        ans = 0

        for i in range(n):
            for j in range(i + 1, n):

                left   = max(bottomLeft[i][0], bottomLeft[j][0])
                right  = min(topRight[i][0], topRight[j][0])
                bottom = max(bottomLeft[i][1], bottomLeft[j][1])
                top    = min(topRight[i][1], topRight[j][1])

                if right > left and top > bottom:
                    side = min(right - left, top - bottom)
                    ans = max(ans, side * side)

        return ans
