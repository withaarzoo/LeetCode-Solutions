class Solution {
    public double averageWaitingTime(int[][] customers) {
        // Variable to store the total waiting time for all customers
        long totalWaitingTime = 0;

        // Variable to track the current time in the restaurant
        int currentTime = 0;

        // Iterate over each customer in the list of customers
        for (int[] customer : customers) {
            // Get the arrival time and cooking time for the current customer
            int arrivalTime = customer[0];
            int cookingTime = customer[1];

            // Update the current time to be the later of the current time or the customer's
            // arrival time,
            // then add the cooking time. This represents the time at which the current
            // customer's order is finished.
            currentTime = Math.max(currentTime, arrivalTime) + cookingTime;

            // Calculate the waiting time for the current customer by subtracting their
            // arrival time
            // from the time their order is finished, and add it to the total waiting time.
            totalWaitingTime += currentTime - arrivalTime;
        }

        // Calculate and return the average waiting time by dividing the total waiting
        // time
        // by the number of customers. Cast totalWaitingTime to double to ensure the
        // result is a double.
        return (double) totalWaitingTime / customers.length;
    }
}
