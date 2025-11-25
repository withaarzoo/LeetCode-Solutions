class Solution:
    def smallestRepunitDivByK(self, k: int) -> int:
        # If k has factor 2 or 5, no repunit (111...) is divisible by k.
        if k % 2 == 0 or k % 5 == 0:
            return -1

        rem = 0  # current remainder
        # At most k different remainders before we repeat
        for length in range(1, k + 1):
            # Append 1: new remainder = (old * 10 + 1) % k
            rem = (rem * 10 + 1) % k
            if rem == 0:
                return length

        return -1
