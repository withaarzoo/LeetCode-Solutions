/**
 * @param {string} sentence
 * @return {boolean}
 */
var isCircularSentence = function (sentence) {
  // Step 1: Split the sentence into words
  const words = sentence.split(" ");

  // Step 2: Check adjacent pairs and the circular condition
  for (let i = 0; i < words.length; i++) {
    let lastChar = words[i].charAt(words[i].length - 1);
    let firstChar = words[(i + 1) % words.length].charAt(0);
    if (lastChar !== firstChar) {
      return false;
    }
  }

  return true;
};
