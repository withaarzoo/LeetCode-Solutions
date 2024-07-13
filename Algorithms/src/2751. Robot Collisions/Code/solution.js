var survivedRobotsHealths = function (positions, healths, directions) {
  // Get the number of robots
  let n = positions.length;

  // Create an array of indices from 0 to n-1
  let indices = Array.from({ length: n }, (_, i) => i);

  // Stack to keep track of robots moving to the right
  let stack = [];

  // Result array to store healths of surviving robots
  let result = [];

  // Sort indices based on the positions of the robots
  indices.sort((a, b) => positions[a] - positions[b]);

  // Iterate through each robot's index in the sorted order
  for (let currentIndex of indices) {
    // If the current robot is moving to the right
    if (directions[currentIndex] === "R") {
      // Push its index onto the stack
      stack.push(currentIndex);
    } else {
      // The current robot is moving to the left
      // Resolve conflicts with robots in the stack (moving to the right)
      while (stack.length > 0 && healths[currentIndex] > 0) {
        // Get the index of the robot at the top of the stack
        let topIndex = stack.pop();

        // Compare the health of the two robots
        if (healths[topIndex] > healths[currentIndex]) {
          // The robot moving to the right is stronger
          // Reduce its health by 1
          healths[topIndex] -= 1;
          // The current robot moving to the left is destroyed
          healths[currentIndex] = 0;
          // Push the robot moving to the right back onto the stack
          stack.push(topIndex);
        } else if (healths[topIndex] < healths[currentIndex]) {
          // The current robot moving to the left is stronger
          // Reduce its health by 1
          healths[currentIndex] -= 1;
          // The robot moving to the right is destroyed
          healths[topIndex] = 0;
        } else {
          // Both robots have the same health
          // Both robots are destroyed
          healths[currentIndex] = 0;
          healths[topIndex] = 0;
        }
      }
    }
  }

  // Collect the healths of surviving robots
  for (let i = 0; i < n; i++) {
    if (healths[i] > 0) {
      result.push(healths[i]);
    }
  }

  // Return the result array
  return result;
};
