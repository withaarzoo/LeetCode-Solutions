class Solution:
    def survivedRobotsHealths(self, positions: List[int], healths: List[int], directions: str) -> List[int]:
        # Step 1: Initialize necessary variables
        n = len(positions)  # Number of robots
        indices = list(range(n))  # List of indices for the robots
        stack = []  # Stack to keep track of robots moving to the right ('R')
        result = []  # List to store the final health of survived robots

        # Step 2: Sort indices based on robot positions
        indices.sort(key=lambda x: positions[x])

        # Step 3: Process each robot in the order of their positions
        for currentIndex in indices:
            if directions[currentIndex] == 'R':
                # If the current robot is moving to the right, push its index onto the stack
                stack.append(currentIndex)
            else:
                # If the current robot is moving to the left
                while stack and healths[currentIndex] > 0:
                    topIndex = stack.pop()  # Get the index of the top robot from the stack

                    # Compare the healths of the two robots
                    if healths[topIndex] > healths[currentIndex]:
                        # If the robot moving right has more health
                        healths[topIndex] -= 1  # Reduce its health by 1
                        healths[currentIndex] = 0  # The current robot is destroyed
                        stack.append(topIndex)  # Push the top robot back onto the stack
                    elif healths[topIndex] < healths[currentIndex]:
                        # If the current robot has more health
                        healths[currentIndex] -= 1  # Reduce its health by 1
                        healths[topIndex] = 0  # The top robot is destroyed
                    else:
                        # If both robots have the same health
                        healths[currentIndex] = 0  # Both robots destroy each other
                        healths[topIndex] = 0

        # Step 4: Collect the healths of the survived robots
        for i in range(n):
            if healths[i] > 0:
                result.append(healths[i])

        return result  # Return the list of healths of the survived robots
