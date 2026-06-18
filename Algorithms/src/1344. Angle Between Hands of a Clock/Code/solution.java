class Solution {
    public double angleClock(int hour, int minutes) {
        // Convert 12 to 0 because both point to the same position
        hour %= 12;

        // Minute hand moves 6 degrees per minute
        double minuteAngle = minutes * 6.0;

        // Hour hand moves 30 degrees per hour
        // and 0.5 degrees per minute
        double hourAngle = hour * 30.0 + minutes * 0.5;

        // Find the difference between the two angles
        double diff = Math.abs(hourAngle - minuteAngle);

        // Return the smaller angle
        return Math.min(diff, 360.0 - diff);
    }
}