class Solution
{
public:
    int minOperations(vector<string> &logs)
    {
        int depth = 0; // Initialize depth to 0, representing the root directory
        for (const string &log : logs)
        { // Iterate over each log in the logs vector
            if (log == "../")
            { // If the log is "../", move one level up
                if (depth > 0)
                    depth--; // Only move up if depth is greater than 0
            }
            else if (log != "./")
            {            // If the log is not "./", move one level down
                depth++; // Increase depth to represent moving into a subdirectory
            }
            // No action is needed for "./" since it represents staying in the current directory
        }
        return depth; // Return the final depth, representing the folder level from the root
    }
};
