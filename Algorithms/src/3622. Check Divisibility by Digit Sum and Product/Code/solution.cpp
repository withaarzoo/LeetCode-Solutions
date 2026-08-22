class Solution {
public:
    bool checkDivisibility(int n) {
        int original = n; // I save the original number because n will be changed while extracting digits.
        int digitSum = 0; // I start the sum at 0 because digits will be added to it.
        int digitProduct = 1; // I start the product at 1 because it is the neutral value for multiplication.

        while (n > 0) { // I keep processing until every digit has been removed.
            int digit = n % 10; // I extract the last digit of the current number.
            digitSum += digit; // I add the current digit to the total digit sum.
            digitProduct *= digit; // I multiply the current digit into the total digit product.
            n /= 10; // I remove the last digit using integer division.
        }

        int divisor = digitSum + digitProduct; // I calculate the required sum of digit sum and digit product.
        return original % divisor == 0; // I return true only when the original number is divisible by this value.
    }
};