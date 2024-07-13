import java.util.*;

class Solution {
    public List<Integer> survivedRobotsHealths(int[] positions, int[] healths, String directions) {
        int n = positions.length; // Get the number of robots
        List<Integer> result = new ArrayList<>(); // To store the final health of survived robots
        int[] indices = new int[n]; // Array to hold the original indices of the robots
        Stack<Integer> stack = new Stack<>(); // Stack to manage the robots moving to the right

        // Initialize indices array with the original indices
        for (int i = 0; i < n; i++) {
            indices[i] = i;
        }

        // Sort the indices array based on the positions of robots
        Arrays.sort(indices, (a, b) -> Integer.compare(positions[a], positions[b]));

        // Process each robot based on the sorted indices
        for (int currentIndex : indices) {
            if (directions.charAt(currentIndex) == 'R') {
                // If the robot is moving to the right, push its index onto the stack
                stack.push(currentIndex);
            } else {
                // If the robot is moving to the left, handle collisions
                while (!stack.isEmpty() && healths[currentIndex] > 0) {
                    int topIndex = stack.pop(); // Get the last robot moving to the right from the stack

                    if (healths[topIndex] > healths[currentIndex]) {
                        // If the health of the right-moving robot is greater
                        healths[topIndex] -= 1; // Reduce the health of the right-moving robot
                        healths[currentIndex] = 0; // Set the health of the current left-moving robot to 0
                        stack.push(topIndex); // Push the right-moving robot back onto the stack
                    } else if (healths[topIndex] < healths[currentIndex]) {
                        // If the health of the right-moving robot is lesser
                        healths[currentIndex] -= 1; // Reduce the health of the current left-moving robot
                        healths[topIndex] = 0; // Set the health of the right-moving robot to 0
                    } else {
                        // If both robots have the same health
                        healths[currentIndex] = 0; // Set the health of both robots to 0
                        healths[topIndex] = 0;
                    }
                }
            }
        }

        // Collect the health of survived robots
        for (int i = 0; i < n; i++) {
            if (healths[i] > 0) {
                result.add(healths[i]);
            }
        }

        return result; // Return the list of healths of survived robots
    }
}
