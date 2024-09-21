class Solution
{
public:
    // Main function to return lexicographical order of numbers from 1 to n
    vector<int> lexicalOrder(int n)
    {
        vector<int> result; // To store the numbers in lexicographical order

        // Iterate over the numbers from 1 to 9 (as the root nodes for DFS)
        // These numbers will act as starting points for building other numbers
        for (int i = 1; i <= 9; ++i)
        {
            // Perform a Depth-First Search (DFS) starting with each number
            dfs(i, n, result);
        }
        return result; // Return the final result after DFS on all root nodes
    }

    // Helper function to perform DFS and generate numbers in lexicographical order
    void dfs(int curr, int n, vector<int> &result)
    {
        // If the current number exceeds the upper limit n, stop the recursion
        if (curr > n)
            return;

        // Add the current number to the result list
        result.push_back(curr);

        // Try to generate the next lexicographical numbers by appending digits 0 to 9
        for (int i = 0; i <= 9; ++i)
        {
            // Calculate the next number by appending i to the current number
            int nextNum = curr * 10 + i;

            // If the next number exceeds n, no need to continue with this branch
            if (nextNum > n)
                break;

            // Recursively perform DFS with the next number
            dfs(nextNum, n, result);
        }
    }
};
