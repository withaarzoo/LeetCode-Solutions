class Solution:
    def sumAndMultiply(self, s: str, queries: List[List[int]]) -> List[int]:
        # I use the required modulo for all concatenated number calculations.
        MOD = 10**9 + 7
        n = len(s)

        # non_zero_count[i] = number of non-zero digits in s[0:i].
        non_zero_count = [0] * (n + 1)

        # I keep only non-zero digits because zeroes never enter x or its digit sum.
        digits = []

        # I build the position mapping and compressed sequence together.
        for i, ch in enumerate(s):
            # I copy the previous count first.
            non_zero_count[i + 1] = non_zero_count[i]

            # Only a non-zero digit increases the compressed sequence length.
            if ch != '0':
                non_zero_count[i + 1] += 1
                digits.append(int(ch))

        k = len(digits)

        # prefix_value[i] stores the first i compressed digits modulo MOD.
        prefix_value = [0] * (k + 1)

        # prefix_sum[i] stores the sum of the first i compressed digits.
        prefix_sum = [0] * (k + 1)

        # power10[i] stores 10^i modulo MOD.
        power10 = [1] * (k + 1)

        # I build all prefix arrays over the compressed digits.
        for i, digit in enumerate(digits):
            # I append the current digit to the previous prefix number.
            prefix_value[i + 1] = (
                prefix_value[i] * 10 + digit
            ) % MOD

            # I extend the digit-sum prefix.
            prefix_sum[i + 1] = (
                prefix_sum[i] + digit
            )

            # I compute the next power of 10 for range extraction.
            power10[i + 1] = (
                power10[i] * 10
            ) % MOD

        # I collect one answer for every query.
        answer = []

        # Every query now takes constant time.
        for l, r in queries:
            # left counts non-zero digits before l.
            left = non_zero_count[l]

            # right counts non-zero digits through r.
            right = non_zero_count[r + 1]

            # length is the number of digits in the compressed query range.
            length = right - left

            # I remove the earlier prefix after shifting it by length places.
            x = (
                prefix_value[right]
                - prefix_value[left] * power10[length]
            ) % MOD

            # I get the digit sum by subtracting two prefix sums.
            digit_sum = (
                prefix_sum[right]
                - prefix_sum[left]
            )

            # I multiply the compressed number by its digit sum.
            answer.append((x * digit_sum) % MOD)

        return answer