// MyCalendar struct represents a calendar system where events can be booked
// The bookings field is a 2D slice (a slice of slices) which stores all the booked intervals.
type MyCalendar struct {
    bookings [][]int  // 2D slice to store booked time intervals [start, end)
}

// Constructor initializes the MyCalendar object
// It returns an empty MyCalendar instance with no bookings.
func Constructor() MyCalendar {
    // Return a MyCalendar object with an empty bookings list
    return MyCalendar{bookings: [][]int{}}
}

// Book method attempts to book a new event with a start and end time.
// It returns true if the event can be booked without overlapping any existing event.
// If there is an overlap, it returns false.
func (this *MyCalendar) Book(start int, end int) bool {
    // Iterate over all the existing bookings to check for any overlap.
    for _, event := range this.bookings {
        existingStart := event[0] // The start time of the existing booking
        existingEnd := event[1]   // The end time of the existing booking
        
        // Check for overlap condition:
        // Two intervals [start1, end1) and [start2, end2) overlap if:
        // start1 < end2 AND start2 < end1
        if start < existingEnd && end > existingStart {
            // If the new event's start time is before the existing event's end time 
            // and the new event's end time is after the existing event's start time,
            // there's an overlap, so return false.
            return false
        }
    }
    
    // If no overlap is found, append the new event [start, end) to the bookings list
    this.bookings = append(this.bookings, []int{start, end})
    
    // Return true indicating the booking was successful
    return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();  // Create a new MyCalendar object
 * param_1 := obj.Book(start,end);  // Book a new event with start and end times
 */
