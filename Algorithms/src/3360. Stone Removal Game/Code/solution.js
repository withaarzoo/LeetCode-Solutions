/**
 * @param {number} n
 * @return {boolean}
 */
var canAliceWin = function(n) {
    let remove = 10;
    let aliceTurn = true;
    while (n >= remove) {
        n -= remove;
        remove--;
        aliceTurn = !aliceTurn;
    }
    return !aliceTurn;
};