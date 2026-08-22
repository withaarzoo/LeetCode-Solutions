class Solution:
    def checkDivisibility(self, n: int) -> bool:
        original = n  # I save the original value because n will change while extracting digits.
        digit_sum = 0  # I start the digit sum at 0 for addition.
        digit_product = 1  # I start the digit product at 1 for multiplication.

        while n > 0:  # I continue until every digit has been processed.
            digit = n % 10  # I extract the last digit of the current number.
            digit_sum += digit  # I add the digit to the total digit sum.
            digit_product *= digit  # I multiply the digit into the total digit product.
            n //= 10  # I remove the last digit using integer division.

        divisor = digit_sum + digit_product  # I calculate the value that must divide the original number.
        return original % divisor == 0  # I return whether the original number is exactly divisible.