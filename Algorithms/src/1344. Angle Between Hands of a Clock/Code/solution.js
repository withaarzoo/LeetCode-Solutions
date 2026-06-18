/**
 * @param {number} hour
 * @param {number} minutes
 * @return {number}
 */
var angleClock = function (hour, minutes) {
  // Convert 12 to 0 because both point to the same position
  hour %= 12;

  // Minute hand moves 6 degrees per minute
  const minuteAngle = minutes * 6;

  // Hour hand moves 30 degrees per hour
  // and 0.5 degrees per minute
  const hourAngle = hour * 30 + minutes * 0.5;

  // Find the difference between both angles
  const diff = Math.abs(hourAngle - minuteAngle);

  // Return the smaller angle
  return Math.min(diff, 360 - diff);
};
