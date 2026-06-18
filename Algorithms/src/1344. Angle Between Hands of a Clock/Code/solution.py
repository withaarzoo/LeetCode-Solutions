class Solution:
    def angleClock(self, hour: int, minutes: int) -> float:
        # Convert 12 to 0 because both point to the same position
        hour %= 12

        # Minute hand moves 6 degrees per minute
        minute_angle = minutes * 6.0

        # Hour hand moves 30 degrees per hour
        # and 0.5 degrees per minute
        hour_angle = hour * 30.0 + minutes * 0.5

        # Find the difference between both angles
        diff = abs(hour_angle - minute_angle)

        # Return the smaller angle
        return min(diff, 360.0 - diff)