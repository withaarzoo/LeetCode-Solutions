var MyCalendarTwo = function () {
  this.single = []; // Stores single booked intervals
  this.doubleBooked = []; // Stores double booked intervals
};

/**
 * @param {number} start
 * @param {number} end
 * @return {boolean}
 */
MyCalendarTwo.prototype.book = function (start, end) {
  // Check for triple booking by overlapping with double booked intervals
  for (let [s, e] of this.doubleBooked) {
    if (Math.max(start, s) < Math.min(end, e)) {
      return false; // Triple booking detected
    }
  }

  // Add overlapping parts to double bookings
  for (let [s, e] of this.single) {
    if (Math.max(start, s) < Math.min(end, e)) {
      this.doubleBooked.push([Math.max(start, s), Math.min(end, e)]);
    }
  }

  // Add the event to single bookings
  this.single.push([start, end]);
  return true;
};
