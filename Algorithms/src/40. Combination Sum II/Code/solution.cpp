#include <vector>
#include <algorithm>

class Solution
{
public:
    // Main function to find all unique combinations that sum up to the target
    std::vector<std::vector<int>> combinationSum2(std::vector<int> &candidates, int target)
    {
        // Sort the input array to handle duplicates and make it easier to skip them later
        std::sort(candidates.begin(), candidates.end());

        // This will store all the valid combinations that sum up to the target
        std::vector<std::vector<int>> result;

        // Temporary vector to store the current combination being explored
        std::vector<int> current;

        // Start the backtracking process from the first index
        backtrack(candidates, target, 0, current, result);

        // Return all the found combinations
        return result;
    }

private:
    // Helper function to perform the backtracking
    void backtrack(std::vector<int> &candidates, int target, int start, std::vector<int> &current, std::vector<std::vector<int>> &result)
    {
        // Base case: If the remaining target is 0, we found a valid combination
        if (target == 0)
        {
            result.push_back(current); // Add the current combination to the result
            return;                    // Exit the function since we've found a valid combination
        }

        // Iterate through the candidates starting from the 'start' index
        for (int i = start; i < candidates.size(); i++)
        {
            // Skip duplicates: If the current candidate is the same as the previous one, skip it to avoid duplicate combinations
            if (i > start && candidates[i] == candidates[i - 1])
                continue;

            // If the current candidate is greater than the remaining target, stop exploring further (since array is sorted, all following numbers will be too large)
            if (candidates[i] > target)
                break;

            // Include the current candidate in the current combination
            current.push_back(candidates[i]);

            // Recursively call the backtrack function with the updated target (target - candidates[i]) and move to the next index (i + 1)
            backtrack(candidates, target - candidates[i], i + 1, current, result);

            // Backtrack: Remove the last added candidate to explore other potential combinations
            current.pop_back();
        }
    }
};
