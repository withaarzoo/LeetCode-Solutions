/**
 * @param {number} turnedOn
 * @return {string[]}
 */
var readBinaryWatch = function (turnedOn) {
  const result = [];

  // Try all hours
  for (let hour = 0; hour < 12; hour++) {
    // Try all minutes
    for (let minute = 0; minute < 60; minute++) {
      // Count number of set bits
      const countBits = (n) => n.toString(2).split("0").join("").length;

      if (countBits(hour) + countBits(minute) === turnedOn) {
        const time = hour + ":" + (minute < 10 ? "0" + minute : minute);

        result.push(time);
      }
    }
  }

  return result;
};
