import java.util.ArrayList;

class MyCalendar {
    // This ArrayList stores all the bookings made in the calendar.
    // Each booking is represented by an array of two integers [start, end],
    // where 'start' is the start time of the event and 'end' is the end time.
    private ArrayList<int[]> bookings;

    // Constructor initializes the ArrayList to keep track of all bookings.
    public MyCalendar() {
        // Initialize the bookings list where all booked events will be stored.
        bookings = new ArrayList<>();
    }

    // Method to attempt booking a new event in the calendar.
    // Returns 'true' if the booking can be made (no overlap), and 'false'
    // otherwise.
    public boolean book(int start, int end) {
        // Iterate through the list of previously booked events.
        for (int[] event : bookings) {
            // Extract the start and end times of the current booked event.
            int existingStart = event[0]; // Start time of an existing event.
            int existingEnd = event[1]; // End time of an existing event.

            // Check if there is an overlap between the new event and the current event.
            // Two events overlap if the start time of the new event is less than
            // the end time of the current event AND the end time of the new event is
            // greater than the start time of the current event.
            if (start < existingEnd && end > existingStart) {
                // If there is an overlap, return 'false', indicating the booking cannot be
                // made.
                return false;
            }
        }

        // If no overlap is found, add the new event (start, end) to the bookings list.
        bookings.add(new int[] { start, end });

        // Return 'true', indicating the booking was successful.
        return true;
    }
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * MyCalendar obj = new MyCalendar();
 * boolean param_1 = obj.book(start,end);
 */
