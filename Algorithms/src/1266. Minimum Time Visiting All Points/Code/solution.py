class Solution:
    def minTimeToVisitAllPoints(self, points):
        total_time = 0

        for i in range(1, len(points)):
            dx = abs(points[i][0] - points[i - 1][0])
            dy = abs(points[i][1] - points[i - 1][1])

            total_time += max(dx, dy)

        return total_time
