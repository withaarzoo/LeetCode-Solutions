class Solution:
    def averageWaitingTime(self, customers: List[List[int]]) -> float:
        # Initialize the total waiting time to 0
        total_waiting_time = 0
        
        # Initialize the current time to 0
        current_time = 0
        
        # Loop through each customer's arrival and cooking times
        for arrival_time, cooking_time in customers:
            # Update current time to either the current time or the arrival time, whichever is later
            # Then add the cooking time
            current_time = max(current_time, arrival_time) + cooking_time
            
            # Calculate the waiting time for the current customer
            # Waiting time is the difference between when the customer finishes (current_time)
            # and when they arrived (arrival_time)
            total_waiting_time += current_time - arrival_time
        
        # Calculate the average waiting time by dividing the total waiting time
        # by the number of customers
        return total_waiting_time / len(customers)
