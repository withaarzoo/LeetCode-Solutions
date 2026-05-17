/**
 * @param {number[]} arr
 * @param {number} start
 * @return {boolean}
 */
var canReach = function (arr, start) {
  // Array to track visited indexes
  const visited = new Array(arr.length).fill(false);

  // DFS function
  const dfs = (index) => {
    // Invalid index
    if (index < 0 || index >= arr.length) {
      return false;
    }

    // Skip visited indexes
    if (visited[index]) {
      return false;
    }

    // Found value 0
    if (arr[index] === 0) {
      return true;
    }

    // Mark current index as visited
    visited[index] = true;

    // Explore forward and backward
    return dfs(index + arr[index]) || dfs(index - arr[index]);
  };

  // Start DFS from given index
  return dfs(start);
};
