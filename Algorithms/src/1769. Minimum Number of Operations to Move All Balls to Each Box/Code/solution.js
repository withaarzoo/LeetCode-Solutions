/**
 * @param {string} boxes
 * @return {number[]}
 */
var minOperations = function (boxes) {
  const n = boxes.length;
  const answer = new Array(n).fill(0);

  // Left-to-right pass
  let balls = 0,
    operations = 0;
  for (let i = 0; i < n; i++) {
    answer[i] += operations;
    balls += boxes[i] === "1" ? 1 : 0; // Count balls
    operations += balls; // Add the current number of balls to operations
  }

  // Right-to-left pass
  balls = 0;
  operations = 0;
  for (let i = n - 1; i >= 0; i--) {
    answer[i] += operations;
    balls += boxes[i] === "1" ? 1 : 0; // Count balls
    operations += balls; // Add the current number of balls to operations
  }

  return answer;
};
