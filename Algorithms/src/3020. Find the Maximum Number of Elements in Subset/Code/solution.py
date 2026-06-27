class Solution:
    def maximumLength(self, nums: List[int]) -> int:
        # Store frequency of every number
        from collections import Counter

        freq = Counter(nums)
        ans = 1

        # Handle value 1 separately
        if 1 in freq:
            cnt = freq[1]

            # Only odd count of ones is valid
            ans = max(ans, cnt if cnt % 2 else cnt - 1)

        # Try every distinct starting value
        for start in freq:
            if start == 1:
                continue

            cur = start
            length = 0

            while cur in freq:
                # Use two copies if available
                if freq[cur] >= 2:
                    length += 2

                    # Move to the squared value
                    cur *= cur
                else:
                    # Single copy becomes the center
                    length += 1
                    break

            # No center found
            if length % 2 == 0:
                length -= 1

            ans = max(ans, length)

        return ans