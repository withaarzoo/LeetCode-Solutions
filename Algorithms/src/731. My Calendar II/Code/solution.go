type MyCalendarTwo struct {
    single []Interval
    doubleBooked []Interval
}

type Interval struct {
    start, end int
}

func Constructor() MyCalendarTwo {
    return MyCalendarTwo{}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
    // Check for triple booking by overlapping with double booked intervals
    for _, booking := range this.doubleBooked {
        if max(start, booking.start) < min(end, booking.end) {
            return false // Triple booking detected
        }
    }
    
    // Add overlapping parts to double bookings
    for _, booking := range this.single {
        if max(start, booking.start) < min(end, booking.end) {
            this.doubleBooked = append(this.doubleBooked, Interval{max(start, booking.start), min(end, booking.end)})
        }
    }
    
    // Add the event to single bookings
    this.single = append(this.single, Interval{start, end})
    return true
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}