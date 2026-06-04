class Solution:
    def totalWaviness(self, num1: int, num2: int) -> int:
        answer = 0

        # Check every number in the range
        for num in range(num1, num2 + 1):
            s = str(num)

            # Numbers with fewer than 3 digits have waviness 0
            if len(s) < 3:
                continue

            # Check every middle digit
            for i in range(1, len(s) - 1):
                # Peak condition
                if s[i] > s[i - 1] and s[i] > s[i + 1]:
                    answer += 1

                # Valley condition
                elif s[i] < s[i - 1] and s[i] < s[i + 1]:
                    answer += 1

        return answer