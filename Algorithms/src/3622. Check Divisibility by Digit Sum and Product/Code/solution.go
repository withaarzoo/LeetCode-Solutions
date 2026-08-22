func checkDivisibility(n int) bool {
	original := n // I save the original number because n will be changed while extracting digits.
	digitSum := 0 // I start the digit sum at 0 for repeated addition.
	digitProduct := 1 // I start the digit product at 1 for repeated multiplication.

	for n > 0 { // I keep processing digits until the number becomes 0.
		digit := n % 10 // I extract the last digit of the current number.
		digitSum += digit // I add the current digit to the digit sum.
		digitProduct *= digit // I multiply the current digit into the digit product.
		n /= 10 // I remove the last digit using integer division.
	}

	divisor := digitSum + digitProduct // I calculate the required divisor from the sum and product.
	return original%divisor == 0 // I return true only when the original number has no remainder.
}