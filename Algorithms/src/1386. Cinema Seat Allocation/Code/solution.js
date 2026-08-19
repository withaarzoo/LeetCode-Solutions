/**
 * @param {number} n
 * @param {number[][]} reservedSeats
 * @return {number}
 */
var maxNumberOfFamilies = function (n, reservedSeats) {
  // Store the reserved seats of every affected row as a bitmask.
  const rows = new Map();

  // Process every reserved seat once.
  for (const [row, col] of reservedSeats) {
    // Seats 1 and 10 never belong to a valid four-seat block.
    if (col >= 2 && col <= 9) {
      // Get the current mask, or use 0 if this is the first seat in the row.
      const currentMask = rows.get(row) || 0;

      // Set the bit corresponding to the reserved seat.
      rows.set(row, currentMask | (1 << col));
    }
  }

  // Every row without relevant reservations can hold two groups.
  let answer = 2 * (n - rows.size);

  // Create masks for seats 2-5, 4-7, and 6-9.
  const left = (1 << 2) | (1 << 3) | (1 << 4) | (1 << 5);
  const middle = (1 << 4) | (1 << 5) | (1 << 6) | (1 << 7);
  const right = (1 << 6) | (1 << 7) | (1 << 8) | (1 << 9);

  // Process only rows that contain relevant reserved seats.
  for (const mask of rows.values()) {
    // A block is free when none of its seats are reserved.
    const canLeft = (mask & left) === 0;
    const canMiddle = (mask & middle) === 0;
    const canRight = (mask & right) === 0;

    // Left and right blocks do not overlap, so both groups can fit.
    if (canLeft && canRight) {
      answer += 2;
    }
    // If at least one block is free, one group can fit.
    else if (canLeft || canMiddle || canRight) {
      answer += 1;
    }
    // Otherwise, no group can be placed in this row.
  }

  // Return the maximum number of groups.
  return answer;
};
