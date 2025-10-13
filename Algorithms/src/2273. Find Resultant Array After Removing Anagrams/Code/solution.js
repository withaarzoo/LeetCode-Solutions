/**
 * @param {string[]} words
 * @return {string[]}
 */
var removeAnagrams = function (words) {
  const result = [];
  let prevSig = "";

  for (const w of words) {
    // signature by sorting characters
    const sig = w.split("").sort().join("");
    if (sig !== prevSig) {
      result.push(w);
      prevSig = sig;
    }
  }
  return result;
};
