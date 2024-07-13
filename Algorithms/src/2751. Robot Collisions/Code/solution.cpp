class Solution
{
public:
    vector<int> survivedRobotsHealths(vector<int> &positions, vector<int> &healths, string directions)
    {
        int n = positions.size();       // Get the number of robots
        vector<int> indices(n), result; // Create indices array and result array
        stack<int> stack;               // Stack to keep track of robots moving to the right ('R')

        // Initialize the indices array with 0 to n-1
        for (int index = 0; index < n; ++index)
        {
            indices[index] = index;
        }

        // Sort the indices based on the positions of robots
        sort(indices.begin(), indices.end(),
             [&](int lhs, int rhs)
             { return positions[lhs] < positions[rhs]; });

        // Iterate over each robot in the sorted order of their positions
        for (int currentIndex : indices)
        {
            // If the current robot is moving to the right, push its index onto the stack
            if (directions[currentIndex] == 'R')
            {
                stack.push(currentIndex);
            }
            else
            {
                // If the current robot is moving to the left, check for collisions with robots in the stack
                while (!stack.empty() && healths[currentIndex] > 0)
                {
                    int topIndex = stack.top(); // Get the index of the robot at the top of the stack
                    stack.pop();                // Remove the top robot from the stack

                    // Determine the result of the collision based on healths
                    if (healths[topIndex] > healths[currentIndex])
                    {
                        healths[topIndex] -= 1;    // Decrease the health of the top robot
                        healths[currentIndex] = 0; // Set the current robot's health to 0 (it is destroyed)
                        stack.push(topIndex);      // Push the top robot back onto the stack
                    }
                    else if (healths[topIndex] < healths[currentIndex])
                    {
                        healths[currentIndex] -= 1; // Decrease the health of the current robot
                        healths[topIndex] = 0;      // Set the top robot's health to 0 (it is destroyed)
                    }
                    else
                    {
                        // Both robots have the same health, so they destroy each other
                        healths[currentIndex] = 0;
                        healths[topIndex] = 0;
                    }
                }
            }
        }

        // Collect the healths of the surviving robots
        for (int index = 0; index < n; ++index)
        {
            if (healths[index] > 0)
            {
                result.push_back(healths[index]);
            }
        }
        return result; // Return the healths of the surviving robots
    }
};
