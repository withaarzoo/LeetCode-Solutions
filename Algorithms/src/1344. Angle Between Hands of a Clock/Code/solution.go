import "math"

func angleClock(hour int, minutes int) float64 {
    // Convert 12 to 0 because both point to the same position
    hour %= 12

    // Minute hand moves 6 degrees per minute
    minuteAngle := float64(minutes) * 6.0

    // Hour hand moves 30 degrees per hour
    // and 0.5 degrees per minute
    hourAngle := float64(hour)*30.0 + float64(minutes)*0.5

    // Find the difference between both angles
    diff := math.Abs(hourAngle - minuteAngle)

    // Return the smaller angle
    return math.Min(diff, 360.0-diff)
}