import java.util.ArrayList;
import java.util.List;

class Solution {

    // Main method that returns a list of integers in lexical order from 1 to n
    public List<Integer> lexicalOrder(int n) {
        List<Integer> result = new ArrayList<>(); // List to store the result in lexical order

        // Iterate from 1 to 9 to start generating numbers in lexicographical order
        for (int i = 1; i <= 9; i++) {
            dfs(i, n, result); // Perform a depth-first search (DFS) starting with each digit from 1 to 9
        }

        return result; // Return the result list once all numbers are added
    }

    // Helper method to perform DFS to generate numbers in lexicographical order
    private void dfs(int curr, int n, List<Integer> result) {
        // Base condition: If the current number exceeds n, stop further exploration
        if (curr > n)
            return;

        // Add the current number to the result list
        result.add(curr);

        // Generate the next numbers by appending digits (0-9) to the current number
        // This will help in exploring numbers like 10, 100, 101, etc.
        for (int i = 0; i <= 9; i++) {
            int nextNum = curr * 10 + i; // Create the next number by appending digit i to the current number

            // If the generated number exceeds n, break the loop to avoid unnecessary
            // recursion
            if (nextNum > n)
                break;

            // Recursively explore the next number
            dfs(nextNum, n, result);
        }
    }
}
