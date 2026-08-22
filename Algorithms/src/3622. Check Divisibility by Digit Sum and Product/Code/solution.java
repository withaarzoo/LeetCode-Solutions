class Solution {
    public boolean checkDivisibility(int n) {
        int original = n; // I save the original number because n will change while extracting digits.
        int digitSum = 0; // I start the digit sum at 0 for repeated addition.
        int digitProduct = 1; // I start the digit product at 1 for repeated multiplication.

        while (n > 0) { // I continue until all digits have been processed.
            int digit = n % 10; // I get the last digit of the current number.
            digitSum += digit; // I add the digit to the total sum.
            digitProduct *= digit; // I multiply the digit into the total product.
            n /= 10; // I remove the last digit from the number.
        }

        int divisor = digitSum + digitProduct; // I calculate the value that must divide the original number.
        return original % divisor == 0; // I return whether the remainder is 0.
    }
}