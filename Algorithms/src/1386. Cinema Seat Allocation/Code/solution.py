from typing import List

class Solution:

    def maxNumberOfFamilies(self, n: int, reservedSeats: List[List[int]]) -> int:
        # Store the reserved seats of each affected row as a bitmask.
        rows = {}

        # Process every reserved seat once.
        for row, col in reservedSeats:
            # Seats 1 and 10 cannot be part of any four-person group.
            if 2 <= col <= 9:
                # Get the current mask, or 0 if this row has not been seen yet.
                current_mask = rows.get(row, 0)

                # Set the bit corresponding to the reserved seat.
                rows[row] = current_mask | (1 << col)

        # Every row not stored in the map has no relevant reservations.
        # Such a row can always contain two groups.
        answer = 2 * (n - len(rows))

        # Create masks for the three possible blocks.
        # left = seats 2-5, middle = seats 4-7, right = seats 6-9.
        left = (1 << 2) | (1 << 3) | (1 << 4) | (1 << 5)
        middle = (1 << 4) | (1 << 5) | (1 << 6) | (1 << 7)
        right = (1 << 6) | (1 << 7) | (1 << 8) | (1 << 9)

        # Check only rows that contain relevant reserved seats.
        for mask in rows.values():
            # Check whether each possible block is completely free.
            can_left = (mask & left) == 0
            can_middle = (mask & middle) == 0
            can_right = (mask & right) == 0

            # Left and right do not overlap, so both groups can be placed.
            if can_left and can_right:
                answer += 2

            # Otherwise, any one available block allows one group.
            elif can_left or can_middle or can_right:
                answer += 1

            # If none is available, this row gets zero groups.

        # Return the maximum number of four-person groups.
        return answer