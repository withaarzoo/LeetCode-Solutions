class Solution:
    def totalNumbers(self, digits: List[int]) -> int:
        # I store how many times each digit appears.
        freq = [0] * 10

        # I build the frequency table so duplicate digits are handled correctly.
        for digit in digits:
            freq[digit] += 1

        # I store the total number of distinct valid 3-digit even numbers.
        answer = 0

        # I choose the hundreds digit from 1 to 9 because zero cannot be the first digit.
        for first in range(1, 10):
            # I choose the tens digit from 0 to 9.
            for second in range(10):
                # I choose only even digits for the ones position.
                for third in range(0, 10, 2):
                    # I skip this combination if any required digit is unavailable.
                    if freq[first] == 0 or freq[second] == 0 or freq[third] == 0:
                        continue

                    # Three equal digits require three copies of that digit.
                    if first == second == third and freq[first] < 3:
                        continue

                    # Equal first and second digits require two copies.
                    if first == second and freq[first] < 2:
                        continue

                    # Equal first and third digits require two copies.
                    if first == third and freq[first] < 2:
                        continue

                    # Equal second and third digits require two copies.
                    if second == third and freq[second] < 2:
                        continue

                    # This combination forms one distinct valid number.
                    answer += 1

        # I return the final count.
        return answer