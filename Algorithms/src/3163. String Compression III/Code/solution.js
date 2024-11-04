/**
 * @param {string} word
 * @return {string}
 */
var compressedString = function (word) {
  let comp = "";
  let count = 1;

  for (let i = 1; i <= word.length; i++) {
    if (i === word.length || word[i] !== word[i - 1] || count === 9) {
      comp += count + word[i - 1];
      count = 1;
    } else {
      count++;
    }
  }

  return comp;
};
