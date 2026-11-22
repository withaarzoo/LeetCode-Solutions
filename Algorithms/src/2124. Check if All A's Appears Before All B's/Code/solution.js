/**
 * @param {string} s
 * @return {boolean}
 */
var checkString = function(s) {
    let seenB = false;
    for (let c of s) {
        if (c === 'b') {
            seenB = true;
        } else if (seenB) {
            return false;
        }
    }
    return true;
};