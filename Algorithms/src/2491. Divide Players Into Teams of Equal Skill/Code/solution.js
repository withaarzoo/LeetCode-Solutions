/**
 * @param {number[]} skill
 * @return {number}
 */
var dividePlayers = function (skill) {
  // Step 1: Sort the skill array
  skill.sort((a, b) => a - b);

  let totalSkill = skill[0] + skill[skill.length - 1]; // Required sum for each pair
  let chemistrySum = 0;

  // Step 2: Pair players using two pointers
  for (let i = 0; i < skill.length / 2; i++) {
    // Check if the sum of current pair matches the required totalSkill
    if (skill[i] + skill[skill.length - 1 - i] !== totalSkill) {
      return -1; // Invalid configuration, return -1
    }
    // Calculate the chemistry (product of pair) and add it to the sum
    chemistrySum += skill[i] * skill[skill.length - 1 - i];
  }

  return chemistrySum; // Return total chemistry
};
