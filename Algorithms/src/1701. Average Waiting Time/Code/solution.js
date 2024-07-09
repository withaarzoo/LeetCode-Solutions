var averageWaitingTime = function (customers) {
  // Initialize total waiting time to 0
  let totalWaitingTime = 0;
  // Initialize current time to 0
  let currentTime = 0;

  // Iterate through each customer in the customers array
  for (const customer of customers) {
    // Extract the arrival time and cooking time for the current customer
    let arrivalTime = customer[0];
    let cookingTime = customer[1];

    // Update the current time to be the maximum of the current time and the arrival time,
    // then add the cooking time of the current customer
    currentTime = Math.max(currentTime, arrivalTime) + cookingTime;

    // Calculate the waiting time for the current customer and add it to the total waiting time
    totalWaitingTime += currentTime - arrivalTime;
  }

  // Calculate and return the average waiting time by dividing the total waiting time
  // by the number of customers
  return totalWaitingTime / customers.length;
};
