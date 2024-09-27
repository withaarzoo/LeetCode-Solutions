#include <vector>
#include <algorithm>
using namespace std;

class MyCalendarTwo
{
    // Vector to store single-booked intervals
    vector<pair<int, int>> single;

    // Vector to store double-booked intervals
    vector<pair<int, int>> double_booked;

public:
    // Constructor to initialize the calendar
    MyCalendarTwo() {}

    // Method to book a new event
    bool book(int start, int end)
    {
        // Step 1: Check if the event would cause a triple booking
        // Iterate over the double_booked intervals to see if the new event overlaps with any of them
        for (const auto &[s, e] : double_booked)
        {
            // If the overlap condition is met, it means the new event will cause a triple booking
            if (max(start, s) < min(end, e))
            {
                // Triple booking detected, return false
                return false;
            }
        }

        // Step 2: Check for overlapping with single-booked intervals
        // If there's overlap, we need to store that overlapping part as a double booking
        for (const auto &[s, e] : single)
        {
            // If the intervals overlap, calculate the overlapping part and add it to double_booked
            if (max(start, s) < min(end, e))
            {
                // The overlapping interval is the max of the start times and the min of the end times
                double_booked.push_back({max(start, s), min(end, e)});
            }
        }

        // Step 3: Add the current event to the list of single-booked intervals
        // The event can safely be booked without causing triple bookings
        single.push_back({start, end});

        // Step 4: Return true indicating the booking is successful
        return true;
    }
};
