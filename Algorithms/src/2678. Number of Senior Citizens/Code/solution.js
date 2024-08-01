/**
 * Function to count the number of seniors (age > 60) based on details provided.
 *
 * @param {string[]} details - An array of strings where each string contains personal details.
 * @return {number} - The count of people whose age is greater than 60.
 */
var countSeniors = function (details) {
  let count = 0; // Initialize the count of seniors to 0

  // Iterate over each detail in the details array
  for (let detail of details) {
    // Extract the substring representing the age from the detail
    // Assuming the age is always located at a fixed position in the string
    let age_str = detail.substring(11, 13);

    // Convert the extracted substring to an integer
    let age = parseInt(age_str);

    // Check if the extracted age is greater than 60
    if (age > 60) {
      count++; // Increment the count if the condition is true
    }
  }

  return count; // Return the total count of seniors
};
