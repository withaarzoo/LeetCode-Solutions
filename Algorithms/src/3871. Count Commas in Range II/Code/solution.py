class Solution:
    def countCommas(self, n: int) -> int:
        # Start at 1000 because numbers below 1000 contain no commas.
        start = 1000

        # Every number from 1000 onward in the first group has one comma.
        commas = 1

        # Store the total number of commas across the whole range.
        answer = 0

        # Process one comma group at a time.
        while start <= n:
            # If the next group boundary is beyond n, stop the group at n.
            if start > n // 1000:
                end = n
            else:
                # Otherwise, the group ends just before start * 1000.
                end = start * 1000 - 1

            # Count the numbers inside this group.
            count = end - start + 1

            # Add the number of commas contributed by this entire group.
            answer += count * commas

            # Move to the next group, which starts at a power of 1000.
            start *= 1000

            # The next group needs one more comma per number.
            commas += 1

        # Return the final total.
        return answer