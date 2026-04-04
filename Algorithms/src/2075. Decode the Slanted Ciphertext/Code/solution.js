/**
 * @param {string} encodedText
 * @param {number} rows
 * @return {string}
 */
var decodeCiphertext = function (encodedText, rows) {
  // Edge case: empty string
  if (encodedText.length === 0) return "";

  const n = encodedText.length;
  const cols = Math.floor(n / rows);

  let result = [];

  // Start from every column in the first row
  for (let startCol = 0; startCol < cols; startCol++) {
    let row = 0;
    let col = startCol;

    // Move diagonally down-right
    while (row < rows && col < cols) {
      result.push(encodedText[row * cols + col]);
      row++;
      col++;
    }
  }

  // Convert array to string and remove trailing spaces
  return result.join("").replace(/\s+$/, "");
};
