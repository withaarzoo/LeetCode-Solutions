class MyCalendarTwo:

    def __init__(self):
        self.single = []  # Stores single booked intervals
        self.double_booked = []  # Stores double booked intervals

    def book(self, start: int, end: int) -> bool:
        # Check if the event overlaps with any double booking, which would cause a triple booking
        for s, e in self.double_booked:
            if max(start, s) < min(end, e):
                return False  # Triple booking found
        
        # Add the overlapping parts to double bookings
        for s, e in self.single:
            if max(start, s) < min(end, e):
                self.double_booked.append((max(start, s), min(end, e)))
        
        # Add the event to single bookings
        self.single.append((start, end))
        return True