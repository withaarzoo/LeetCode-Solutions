var numberToWords = function (num) {
  // Arrays to store words for numbers below 20, tens, and thousands
  const below_20 = [
    "",
    "One",
    "Two",
    "Three",
    "Four",
    "Five",
    "Six",
    "Seven",
    "Eight",
    "Nine",
    "Ten",
    "Eleven",
    "Twelve",
    "Thirteen",
    "Fourteen",
    "Fifteen",
    "Sixteen",
    "Seventeen",
    "Eighteen",
    "Nineteen",
  ];
  const tens = [
    "",
    "",
    "Twenty",
    "Thirty",
    "Forty",
    "Fifty",
    "Sixty",
    "Seventy",
    "Eighty",
    "Ninety",
  ];
  const thousands = ["", "Thousand", "Million", "Billion"];

  // Edge case for zero
  if (num === 0) return "Zero";

  let result = ""; // To store the final result
  let i = 0; // Index to keep track of thousand, million, billion, etc.

  // Loop until num is greater than 0
  while (num > 0) {
    // Process the last three digits of num
    if (num % 1000 !== 0) {
      // Use the helper function to convert the last three digits
      // and prepend the appropriate thousands place
      result = helper(num % 1000) + thousands[i] + " " + result;
    }
    // Remove the last three digits from num
    num = Math.floor(num / 1000);
    i++; // Move to the next thousands place
  }

  return result.trim(); // Trim any extra spaces and return the result
};

function helper(num) {
  // Arrays to store words for numbers below 20 and tens
  const below_20 = [
    "",
    "One",
    "Two",
    "Three",
    "Four",
    "Five",
    "Six",
    "Seven",
    "Eight",
    "Nine",
    "Ten",
    "Eleven",
    "Twelve",
    "Thirteen",
    "Fourteen",
    "Fifteen",
    "Sixteen",
    "Seventeen",
    "Eighteen",
    "Nineteen",
  ];
  const tens = [
    "",
    "",
    "Twenty",
    "Thirty",
    "Forty",
    "Fifty",
    "Sixty",
    "Seventy",
    "Eighty",
    "Ninety",
  ];

  // Base case: if num is zero, return an empty string
  if (num === 0) return "";
  // If num is less than 20, return the corresponding word from below_20
  else if (num < 20) return below_20[num] + " ";
  // If num is less than 100, return the corresponding tens word and recurse for the remainder
  else if (num < 100)
    return tens[Math.floor(num / 10)] + " " + helper(num % 10);
  // If num is 100 or more, return the word for the hundreds place and recurse for the remainder
  else return below_20[Math.floor(num / 100)] + " Hundred " + helper(num % 100);
}
