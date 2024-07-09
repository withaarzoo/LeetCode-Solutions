class Solution
{
public:
    double averageWaitingTime(vector<vector<int>> &customers)
    {
        // Initialize a variable to store the total waiting time of all customers
        long long total_waiting_time = 0;

        // Initialize a variable to keep track of the current time in the restaurant
        int current_time = 0;

        // Iterate through each customer in the customers list
        for (const auto &customer : customers)
        {
            // Extract the arrival time and cooking time for the current customer
            int arrival_time = customer[0];
            int cooking_time = customer[1];

            // Update the current time to either the customer's arrival time or
            // the end of the previous customer's cooking time, whichever is later,
            // and add the current customer's cooking time
            current_time = max(current_time, arrival_time) + cooking_time;

            // Calculate the waiting time for the current customer by subtracting
            // the arrival time from the current time (which now includes their cooking time)
            total_waiting_time += current_time - arrival_time;
        }

        // Calculate and return the average waiting time by dividing the total
        // waiting time by the number of customers
        return (double)total_waiting_time / customers.size();
    }
};
