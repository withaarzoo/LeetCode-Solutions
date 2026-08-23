class Solution:
    def sumGame(self, num: str) -> bool:
        # Find the index where the right half begins.
        mid = len(num) // 2

        # Store the sums of known digits in both halves.
        left_sum = 0
        right_sum = 0

        # Count how many '?' characters are in both halves.
        left_question = 0
        right_question = 0

        # Scan every character and update the correct half.
        for i, char in enumerate(num):
            if i < mid:
                # This character belongs to the left half.
                if char == '?':
                    left_question += 1
                else:
                    left_sum += int(char)
            else:
                # This character belongs to the right half.
                if char == '?':
                    right_question += 1
                else:
                    right_sum += int(char)

        # Bob can force equality only in this exact situation.
        # Otherwise, Alice can always make the final sums different.
        return 2 * (left_sum - right_sum) != \
               9 * (right_question - left_question)