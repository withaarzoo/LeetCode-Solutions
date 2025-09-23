/**
 * @param {string} version1
 * @param {string} version2
 * @return {number}
 */
var compareVersion = function (version1, version2) {
  let i = 0,
    j = 0;
  const n = version1.length,
    m = version2.length;

  while (i < n || j < m) {
    let num1 = 0,
      num2 = 0;
    // parse next number from version1
    while (i < n && version1[i] !== ".") {
      num1 = num1 * 10 + (version1.charCodeAt(i) - 48); // '0' -> 48
      i++;
    }
    if (i < n && version1[i] === ".") i++;

    // parse next number from version2
    while (j < m && version2[j] !== ".") {
      num2 = num2 * 10 + (version2.charCodeAt(j) - 48);
      j++;
    }
    if (j < m && version2[j] === ".") j++;

    if (num1 < num2) return -1;
    if (num1 > num2) return 1;
  }
  return 0;
};
