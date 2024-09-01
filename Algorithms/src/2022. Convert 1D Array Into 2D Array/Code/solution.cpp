#include <vector>

class Solution
{
public:
    std::vector<std::vector<int>> construct2DArray(std::vector<int> &original, int m, int n)
    {
        // Step 1: Check if the total number of elements in the 'original' vector matches the required size of the 2D array.
        // If the size does not match, it is impossible to form a m*n matrix, so we return an empty 2D array.
        if (original.size() != m * n)
        {
            return {}; // Return an empty vector if the condition is not met.
        }

        // Step 2: Initialize a 2D vector 'result' with 'm' rows and 'n' columns, filled with zeros initially.
        // This will store the final 2D array that we need to construct.
        std::vector<std::vector<int>> result(m, std::vector<int>(n));

        // Step 3: Iterate through each element in the 'original' vector.
        for (int i = 0; i < original.size(); ++i)
        {
            // Calculate the corresponding row index in the 2D array.
            // 'i / n' gives the row index as 'n' elements fit into one row.
            int row = i / n;

            // Calculate the corresponding column index in the 2D array.
            // 'i % n' gives the column index, which is the remainder when 'i' is divided by 'n'.
            int col = i % n;

            // Step 4: Assign the value from the 'original' vector to the correct position in the 2D array.
            result[row][col] = original[i];
        }

        // Step 5: Return the constructed 2D array 'result'.
        return result;
    }
};
