/**
 * @param {string} s
 * @param {string} goal
 * @return {boolean}
 */
var rotateString = function (s, goal) {
  // If lengths differ, s cannot be rotated to match goal
  if (s.length !== goal.length) return false;

  // Concatenate s with itself
  let doubled = s + s;

  // Check if goal is a substring of doubled
  return doubled.includes(goal);
};
