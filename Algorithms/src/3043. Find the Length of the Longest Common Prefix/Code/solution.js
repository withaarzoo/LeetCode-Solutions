/**
 * Function to find the length of the longest common prefix between numbers in two arrays.
 * @param {number[]} arr1 - The first array of numbers.
 * @param {number[]} arr2 - The second array of numbers.
 * @return {number} - The length of the longest common prefix found.
 */
var longestCommonPrefix = function (arr1, arr2) {
  // Create a Map to store all the prefixes found in arr1 and their counts.
  const prefixMap = new Map();

  // Step 1: Build the prefix map for arr1
  // For each number in arr1, we will generate all possible prefixes of that number and store them in the map.
  for (let num of arr1) {
    // Convert the number to a string so we can extract prefixes.
    let strNum = num.toString();
    let prefix = ""; // Initialize an empty string to build the prefix.

    // For each character in the string representation of the number, build the prefix.
    for (let ch of strNum) {
      prefix += ch; // Add the current character to the prefix.

      // Store the prefix in the map and increment its count.
      // If the prefix already exists, increase the count by 1.
      // Otherwise, initialize the count to 1.
      prefixMap.set(prefix, (prefixMap.get(prefix) || 0) + 1);
    }
  }

  // Variable to keep track of the longest common prefix length found.
  let maxLength = 0;

  // Step 2: Check for common prefixes in arr2
  // Now, for each number in arr2, we'll check if any of its prefixes exist in the prefix map from arr1.
  for (let num of arr2) {
    // Convert the number to a string so we can extract prefixes.
    let strNum = num.toString();
    let prefix = ""; // Initialize an empty string to build the prefix.

    // For each character in the string representation of the number, build the prefix.
    for (let ch of strNum) {
      prefix += ch; // Add the current character to the prefix.

      // Check if the current prefix exists in the prefix map (i.e., it was also a prefix in arr1).
      if (prefixMap.has(prefix)) {
        // Update the maxLength to the length of the current prefix if it's the longest found so far.
        maxLength = Math.max(maxLength, prefix.length);
      }
    }
  }

  // Return the length of the longest common prefix found.
  return maxLength;
};
