/**
 * @param {string} s
 * @return {boolean}
 */
var checkOnesSegment = function (s) {
  // If "01" appears, return false
  return !s.includes("01");
};
