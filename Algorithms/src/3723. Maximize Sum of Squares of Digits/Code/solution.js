/**
 * @param {number} num
 * @param {number} sum
 * @return {string}
 */
var maxSumOfSquares = function(num, sum) {
    if (sum > 9 * num) return "";
    let q = Math.floor(sum / 9);
    let r = sum % 9;
    let ans = '9'.repeat(q);
    if (r) ans += String(r);
    ans += '0'.repeat(num - q - (r ? 1 : 0));
    return ans;
};