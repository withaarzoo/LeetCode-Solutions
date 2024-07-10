class Solution {
    public int minOperations(String[] logs) {
        int depth = 0; // Initialize depth to 0, representing the root directory

        // Iterate over each log entry in the logs array
        for (String log : logs) {
            // If the log entry is "../", move one directory up if not already at the root
            if (log.equals("../")) {
                if (depth > 0) {
                    depth--; // Decrease depth by 1
                }
            }
            // If the log entry is "./", it means stay in the same directory, do nothing
            else if (!log.equals("./")) {
                depth++; // For any other log entry (representing a folder), increase depth by 1
            }
        }

        // Return the final calculated depth, which is the minimum number of operations
        // to return to the main folder
        return depth;
    }
}
