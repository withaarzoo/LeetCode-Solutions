/**
 * @param {string} s1
 * @param {string} s2
 * @return {boolean}
 */
var canBeEqual = function (s1, s2) {
  // Get even indexed characters and sort them
  let even1 = [s1[0], s1[2]].sort().join("");
  let even2 = [s2[0], s2[2]].sort().join("");

  // Get odd indexed characters and sort them
  let odd1 = [s1[1], s1[3]].sort().join("");
  let odd2 = [s2[1], s2[3]].sort().join("");

  // Both even and odd groups must be equal
  return even1 === even2 && odd1 === odd2;
};
