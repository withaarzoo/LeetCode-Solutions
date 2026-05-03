/**
 * @param {string} s
 * @param {string} goal
 * @return {boolean}
 */
var rotateString = function(s, goal) {
    // If lengths are different, rotation is impossible
    if (s.length !== goal.length) return false;

    // Concatenate s with itself
    let doubled = s + s;

    // Check if goal is a substring
    return doubled.includes(goal);
};