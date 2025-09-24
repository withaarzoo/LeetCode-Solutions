/**
 * @param {number} numerator
 * @param {number} denominator
 * @return {string}
 */
var fractionToDecimal = function (numerator, denominator) {
  if (numerator === 0) return "0";

  // Use BigInt to avoid overflow for extreme values
  let n = BigInt(numerator);
  let d = BigInt(denominator);

  let res = "";
  // sign: different signs -> negative
  if (n < 0n !== d < 0n) res += "-";

  // work with absolute values
  if (n < 0n) n = -n;
  if (d < 0n) d = -d;

  // integer part
  res += (n / d).toString();
  let rem = n % d;
  if (rem === 0n) return res;

  res += ".";
  const seen = new Map(); // remainder (BigInt) -> index in res string

  while (rem !== 0n) {
    if (seen.has(rem)) {
      const pos = seen.get(rem);
      // insert parentheses around repeating part
      res = res.slice(0, pos) + "(" + res.slice(pos) + ")";
      break;
    }
    seen.set(rem, res.length);
    rem *= 10n;
    res += (rem / d).toString();
    rem = rem % d;
  }
  return res;
};
