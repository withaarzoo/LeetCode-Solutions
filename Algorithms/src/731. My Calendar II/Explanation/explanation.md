# MyCalendarTwo - Step-by-Step Explanation (Multiple Languages)

In this problem, the goal is to implement a booking system where events can be added if they do not result in more than two overlapping events. If a triple booking (three overlapping events) occurs, the booking is rejected. Below is a step-by-step explanation of the logic behind the code for each language.

---

## **C++ Code**

### Step 1: Check for Triple Booking

- Loop through the `double_booked` intervals, which store already double-booked ranges.
- If the new event overlaps with any double-booked range, it would cause a triple booking.
- If such an overlap is found, return `false` to reject the booking.

### Step 2: Identify Overlaps for Double Booking

- Loop through the `single` booked intervals.
- If the new event overlaps with a single-booked event, calculate the overlap and store it in `double_booked` to keep track of double bookings.

### Step 3: Add the Event to Single Bookings

- If no triple booking is detected, add the new event to `single` booked intervals.

### Step 4: Confirm the Booking

- Return `true` to indicate the event has been successfully booked.

---

## **Java Code**

### Step 1: Check for Triple Booking

- Iterate through `doubleBooked`, the list storing intervals that have already been booked twice.
- If the new event overlaps with any interval in `doubleBooked`, return `false` to reject the booking due to potential triple booking.

### Step 2: Identify Overlaps for Double Booking

- Loop through `single` bookings, and if the new event overlaps with any single-booked interval, add the overlapping part to `doubleBooked`.

### Step 3: Add the Event to Single Bookings

- If the event does not result in a triple booking, add it to the list `single`.

### Step 4: Return Success

- Return `true` if the booking is successful.

---

## **JavaScript Code**

### Step 1: Check for Triple Booking

- Loop through the `doubleBooked` array.
- Check if the new event overlaps with any already double-booked interval.
- If an overlap is found, return `false` because it would result in a triple booking.

### Step 2: Identify Overlaps for Double Booking

- Loop through `single` bookings.
- For each single-booked interval that overlaps with the new event, add the overlapping part to `doubleBooked`.

### Step 3: Add the Event to Single Bookings

- Add the new event to the `single` array after ensuring no triple booking occurs.

### Step 4: Confirm the Booking

- Return `true` if the event is successfully booked.

---

## **Python Code**

### Step 1: Check for Triple Booking

- Loop through the `double_booked` intervals.
- If the new event overlaps with any double-booked interval, return `false` because a triple booking would occur.

### Step 2: Identify Overlaps for Double Booking

- Loop through `single` booked intervals.
- For any overlap between the new event and existing single bookings, add the overlapping part to `double_booked`.

### Step 3: Add the Event to Single Bookings

- If no triple booking is detected, add the event to `single` bookings.

### Step 4: Confirm the Booking

- Return `true` if the booking is successful.

---

## **Go Code**

### Step 1: Check for Triple Booking

- Iterate through `doubleBooked` intervals.
- If the new event overlaps with any double-booked interval, return `false` as it would cause a triple booking.

### Step 2: Identify Overlaps for Double Booking

- Loop through the `single` intervals.
- For any overlap with the new event, append the overlapping interval to `doubleBooked`.

### Step 3: Add the Event to Single Bookings

- Add the new event to the `single` bookings if no triple booking is detected.

### Step 4: Return the Booking Status

- Return `true` if the event is successfully booked without causing triple bookings.

---

### Conclusion

The logic across all languages remains consistent:

1. **Check for triple bookings** (overlaps in double-booked intervals).
2. **Store overlaps** of new events with single-booked intervals as double bookings.
3. **Add the new event** to single-booked intervals if no triple booking is found.
4. Return `true` if the booking is successful, otherwise return `false`.
