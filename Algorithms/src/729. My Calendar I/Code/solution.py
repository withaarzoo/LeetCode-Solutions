class MyCalendar:
    
    def __init__(self):
        # Initialize an empty list to store the booking intervals
        # Each booking will be a tuple (start, end) representing the start and end time of the event
        self.bookings = []

    def book(self, start: int, end: int) -> bool:
        """
        Attempt to book a new event. The event is represented by its start and end time.
        If the event does not overlap with any existing bookings, it is added to the list.
        If it overlaps, the booking is rejected.
        
        :param start: Start time of the event (inclusive)
        :param end: End time of the event (exclusive)
        :return: True if the event is successfully booked, False otherwise.
        """
        
        # Iterate through all the existing bookings to check for any overlap
        for existingStart, existingEnd in self.bookings:
            # Check if the new event overlaps with the existing one
            # Two intervals [start1, end1) and [start2, end2) overlap if:
            # - the start of one event is before the end of the other
            # - AND the end of one event is after the start of the other
            # This condition is a simplified overlap check:
            #   If (start < existingEnd and end > existingStart), it means the events overlap.
            if start < existingEnd and end > existingStart:
                # If the new event overlaps with an existing one, reject the booking
                return False
        
        # If no overlap is found, add the new event to the list of bookings
        # Store the event as a tuple (start, end)
        self.bookings.append((start, end))
        
        # Return True to indicate that the booking was successful
        return True


# Example usage:
# Create an instance of MyCalendar
# obj = MyCalendar()
# Try booking a new event with start and end times
# param_1 = obj.book(start, end)
