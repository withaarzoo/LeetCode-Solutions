// minOperations calculates the minimum number of operations to reach the main folder
// given a list of folder navigation logs.
func minOperations(logs []string) int {
    // Initialize depth to 0, representing the main folder
    depth := 0

    // Iterate through each log entry
    for _, log := range logs {
        // If the log is "../", move up one directory if not already at the root
        if log == "../" {
            if depth > 0 {
                depth-- // Decrease depth to move up one directory
            }
        // If the log is "./", stay in the current directory (do nothing)
        } else if log != "./" {
            depth++ // Increase depth to move into a subdirectory
        }
    }

    // Return the current depth, representing the minimum operations needed
    return depth
}
