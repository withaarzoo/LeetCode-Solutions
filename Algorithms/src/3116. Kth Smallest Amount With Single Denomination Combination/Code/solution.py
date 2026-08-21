from typing import List
from math import gcd

class Solution:

    def findKthSmallest(self, coins: List[int], k: int) -> int:
        # Sort coins so smaller denominations are checked first.
        coins.sort()

        # Keep only denominations that are not already covered.
        useful = []

        for coin in coins:
            redundant = False

            # If a smaller kept coin divides this coin, every multiple of
            # this coin is already produced by the smaller denomination.
            for prev in useful:
                if coin % prev == 0:
                    redundant = True
                    break

            # Keep the denomination only if it adds new multiples.
            if not redundant:
                useful.append(coin)

        # The kth multiple of the smallest coin gives a safe upper bound.
        low = 1
        high = useful[0] * k

        m = len(useful)
        total_masks = 1 << m

        # Store the LCM for every subset.
        lcms = [1] * total_masks

        # Store +1 for odd subsets and -1 for even subsets.
        signs = [1] * total_masks

        # Precompute LCM values and inclusion-exclusion signs.
        for mask in range(1, total_masks):
            current_lcm = 1
            bits = 0

            for i in range(m):
                # Include useful[i] when its bit is present in the subset.
                if mask & (1 << i):
                    # Divide first so the multiplication stays smaller.
                    current_lcm //= gcd(current_lcm, useful[i])

                    # If the LCM becomes larger than high, it can never
                    # contribute to any count during the binary search.
                    if current_lcm > high // useful[i]:
                        current_lcm = high + 1
                        break

                    current_lcm *= useful[i]
                    bits += 1

            # Save the final LCM for this subset.
            lcms[mask] = current_lcm

            # Inclusion-exclusion adds odd subsets and subtracts even ones.
            signs[mask] = 1 if bits % 2 == 1 else -1

        # Count valid amounts from 1 through x.
        def count(x: int) -> int:
            result = 0

            for mask in range(1, total_masks):
                # A larger LCM cannot divide any number up to x.
                if lcms[mask] <= x:
                    result += signs[mask] * (x // lcms[mask])

            return result

        # Find the smallest value containing at least k valid amounts.
        while low < high:
            mid = low + (high - low) // 2

            # Keep mid when it already reaches the kth position.
            if count(mid) >= k:
                high = mid
            else:
                # Otherwise, move to larger values.
                low = mid + 1

        # low is the kth smallest valid amount.
        return low