// Define the MyCalendar class constructor
var MyCalendar = function () {
  // Initialize an empty array to store all the booked events
  this.bookings = [];
};

/**
 * Function to book an event in the calendar
 * @param {number} start - The start time of the event
 * @param {number} end - The end time of the event
 * @return {boolean} - Returns true if the event can be booked, false if there is a conflict
 */
MyCalendar.prototype.book = function (start, end) {
  // Loop through all previously booked events to check for any overlaps
  for (let event of this.bookings) {
    let [existingStart, existingEnd] = event; // Destructure to get the start and end times of the booked event

    // Check if the new event overlaps with the current booked event
    // Overlap occurs if the new event's start time is before an existing event's end time
    // AND the new event's end time is after an existing event's start time
    if (start < existingEnd && end > existingStart) {
      // If overlap is detected, return false to indicate the event can't be booked
      return false;
    }
  }

  // If no overlap is found, add the new event (start, end) to the bookings list
  this.bookings.push([start, end]);

  // Return true to indicate the event was successfully booked
  return true;
};

/**
 * Example usage:
 * var obj = new MyCalendar();
 * var param_1 = obj.book(start, end); // Call the book method with event start and end times
 */
