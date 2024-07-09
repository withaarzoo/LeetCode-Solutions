// averageWaitingTime calculates the average waiting time for customers.
func averageWaitingTime(customers [][]int) float64 {
    // Initialize totalWaitingTime to accumulate the total waiting time of all customers.
    totalWaitingTime := 0
    // Initialize currentTime to keep track of the current time in the restaurant.
    currentTime := 0
    
    // Iterate over each customer in the customers array.
    for _, customer := range customers {
        // Extract the arrival time and cooking time for the current customer.
        arrivalTime := customer[0]
        cookingTime := customer[1]
        
        // Update currentTime to the maximum of currentTime and arrivalTime, then add the cooking time.
        // This ensures the chef starts cooking as soon as the customer arrives or as soon as the previous
        // order is finished, whichever is later.
        currentTime = max(currentTime, arrivalTime) + cookingTime
        
        // Calculate the waiting time for the current customer and add it to totalWaitingTime.
        // Waiting time is the difference between the time the order is completed (currentTime)
        // and the arrival time of the customer.
        totalWaitingTime += currentTime - arrivalTime
    }
    
    // Return the average waiting time by dividing the total waiting time by the number of customers.
    // Convert both totalWaitingTime and len(customers) to float64 to get a floating-point result.
    return float64(totalWaitingTime) / float64(len(customers))
}

// max returns the greater of two integers a and b.
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
