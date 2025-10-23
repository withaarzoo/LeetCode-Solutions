class Solution:
    def hasSameDigits(self, s: str) -> bool:
        # convert to list of ints
        digits = [ord(c) - 48 for c in s]

        # reduce until exactly two digits remain
        while len(digits) > 2:
            next_digits = [(digits[i] + digits[i+1]) % 10 for i in range(len(digits)-1)]
            digits = next_digits

        return len(digits) == 2 and digits[0] == digits[1]
