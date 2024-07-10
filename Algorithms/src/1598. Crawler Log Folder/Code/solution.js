var minOperations = function (logs) {
  // Initialize the depth of the current directory to 0
  let depth = 0;

  // Loop through each log in the logs array
  for (const log of logs) {
    // If the log indicates moving up one directory level
    if (log === "../") {
      // Decrease depth if it's greater than 0 (cannot go above the root directory)
      if (depth > 0) depth--;
    }
    // If the log is not a no-op ("./")
    else if (log !== "./") {
      // Move into a subdirectory, increase depth
      depth++;
    }
  }

  // Return the final depth, which represents the minimum number of operations needed
  return depth;
};
