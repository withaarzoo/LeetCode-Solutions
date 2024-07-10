class Solution:
    def minOperations(self, logs: List[str]) -> int:
        # Initialize depth to track the current folder level
        depth = 0
        
        # Iterate through each log entry
        for log in logs:
            # If the log is "../", move up one directory if not already at the root
            if log == "../":
                if depth > 0:
                    depth -= 1
            # If the log is "./", do nothing (stay in the same directory)
            elif log != "./":
                # Any other log means moving into a subdirectory
                depth += 1
        
        # Return the final depth, which represents the minimum number of operations to go back to the main folder
        return depth
