/**
 * @param {number} n
 * @return {boolean}
 */
var checkDivisibility = function(n) {
    const original = n; // I save the original number because n will be reduced digit by digit.
    let digitSum = 0; // I initialize the digit sum for repeated addition.
    let digitProduct = 1; // I initialize the digit product for repeated multiplication.

    while (n > 0) { // I continue until there are no digits left to process.
        const digit = n % 10; // I extract the last digit.
        digitSum += digit; // I add the current digit to the digit sum.
        digitProduct *= digit; // I multiply the current digit into the digit product.
        n = Math.floor(n / 10); // I remove the last digit and keep only the remaining whole-number part.
    }

    const divisor = digitSum + digitProduct; // I calculate the required divisor.
    return original % divisor === 0; // I return true only if the original number divides exactly.
};