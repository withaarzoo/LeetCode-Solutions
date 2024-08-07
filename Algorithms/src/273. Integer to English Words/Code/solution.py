class Solution:
    def __init__(self):
        # Initialize arrays to store words for numbers below 20, tens, and large number groups (thousands, millions, billions)
        self.below_20 = ["", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"]
        self.tens = ["", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"]
        self.thousands = ["", "Thousand", "Million", "Billion"]

    def numberToWords(self, num: int) -> str:
        # If the number is zero, return "Zero"
        if num == 0:
            return "Zero"
        
        result = ""  # Initialize result string
        i = 0  # Initialize index for thousands array
        
        # Process each group of 3 digits (thousands, millions, etc.)
        while num > 0:
            # If the current group of 3 digits is not zero, convert it to words
            if num % 1000 != 0:
                result = self.helper(num % 1000) + self.thousands[i] + " " + result
            # Move to the next group of 3 digits
            num //= 1000
            i += 1  # Increment the index for thousands array
        
        # Return the final result after stripping any extra spaces
        return result.strip()

    def helper(self, num: int) -> str:
        # If the number is zero, return an empty string
        if num == 0:
            return ""
        # If the number is less than 20, get the corresponding word from below_20 array
        elif num < 20:
            return self.below_20[num] + " "
        # If the number is less than 100, get the corresponding tens word and process the remaining digits
        elif num < 100:
            return self.tens[num // 10] + " " + self.helper(num % 10)
        # If the number is 100 or greater, get the hundreds word and process the remaining digits
        else:
            return self.below_20[num // 100] + " Hundred " + self.helper(num % 100)
