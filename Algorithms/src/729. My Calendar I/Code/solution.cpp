#include <vector>
using namespace std;

class MyCalendar
{
private:
    // Vector to store pairs of start and end times for booked events
    vector<pair<int, int>> bookings;

public:
    // Constructor to initialize the MyCalendar object
    MyCalendar() {}

    // Function to book an event in the calendar
    bool book(int start, int end)
    {
        // Iterate through all the existing bookings to check for conflicts
        for (auto &event : bookings)
        {
            int existingStart = event.first; // Start time of an existing event
            int existingEnd = event.second;  // End time of an existing event

            // Check for overlap between the new event and the existing event.
            // Overlap happens if the new event's start is before the existing event's end,
            // and the new event's end is after the existing event's start.
            if (start < existingEnd && end > existingStart)
            {
                // If overlap is found, return false indicating that the booking cannot be made
                return false;
            }
        }

        // If no overlap is found with any of the existing bookings,
        // add the new event (start, end) to the bookings list
        bookings.push_back({start, end});

        // Return true indicating the event was successfully booked
        return true;
    }
};

/**
 * Example of how to use the MyCalendar class:
 * MyCalendar* obj = new MyCalendar();   // Instantiate the MyCalendar object
 * bool isBooked = obj->book(start, end);  // Call the book method with start and end times
 *
 * isBooked will be true if the booking is successful (no overlap), or false if there's a conflict.
 */
