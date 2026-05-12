/**
 * @param {number[][]} tasks
 * @return {number}
 */
var minimumEffort = function (tasks) {
  // Sort by (minimum - actual) descending
  tasks.sort((a, b) => b[1] - b[0] - (a[1] - a[0]));

  let answer = 0; // Minimum initial energy
  let energy = 0; // Current energy

  // Process all tasks
  for (let [actual, minimum] of tasks) {
    // If energy is insufficient,
    // increase it
    if (energy < minimum) {
      let need = minimum - energy;

      answer += need;
      energy += need;
    }

    // Complete the task
    energy -= actual;
  }

  return answer;
};
