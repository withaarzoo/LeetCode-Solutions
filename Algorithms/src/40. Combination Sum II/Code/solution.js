var combinationSum2 = function (candidates, target) {
  // Step 1: Sort the candidates array in ascending order to manage duplicates easily
  candidates.sort((a, b) => a - b);

  // Step 2: Initialize the result array to store all unique combinations
  const result = [];

  // Step 3: Initialize the current array to keep track of the current combination being explored
  const current = [];

  // Step 4: Define the backtracking function
  // 'target' is the remaining sum we need to achieve
  // 'start' is the current index in the candidates array we're exploring
  const backtrack = (target, start) => {
    // Step 5: Base case - If the target becomes zero, it means we've found a valid combination
    if (target === 0) {
      // Add a copy of the current combination to the result array
      result.push([...current]);
      return; // Stop further exploration and backtrack
    }

    // Step 6: Iterate over the candidates starting from the 'start' index
    for (let i = start; i < candidates.length; i++) {
      // Step 7: Skip duplicate candidates to avoid redundant combinations
      // We skip a candidate if it's the same as the previous one and we're not at the starting index
      if (i > start && candidates[i] === candidates[i - 1]) continue;

      // Step 8: If the current candidate exceeds the target, we can stop further exploration
      // This works because the array is sorted, so all subsequent candidates will also be larger
      if (candidates[i] > target) break;

      // Step 9: Choose the current candidate by adding it to the current combination
      current.push(candidates[i]);

      // Step 10: Recursively explore further by reducing the target and moving to the next candidate
      backtrack(target - candidates[i], i + 1);

      // Step 11: Backtrack by removing the last added candidate from the current combination
      // This allows us to explore other potential combinations
      current.pop();
    }
  };

  // Step 12: Start the backtracking process from the first candidate
  backtrack(target, 0);

  // Step 13: Return the result array containing all unique combinations that sum up to the target
  return result;
};
