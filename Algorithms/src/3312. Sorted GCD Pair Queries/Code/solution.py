class Solution:
    def gcdValues(self, nums: List[int], queries: List[int]) -> List[int]:

        # Maximum value in the array.
        mx = max(nums)

        # Frequency of each value.
        freq = [0] * (mx + 1)
        for x in nums:
            freq[x] += 1

        # exact[g] = pairs having GCD exactly g.
        exact = [0] * (mx + 1)

        # Process from largest divisor.
        for g in range(mx, 0, -1):

            # Count divisible numbers.
            cnt = 0
            for m in range(g, mx + 1, g):
                cnt += freq[m]

            # Total pairs with GCD multiple of g.
            pairs = cnt * (cnt - 1) // 2

            # Remove already computed larger GCDs.
            for m in range(g * 2, mx + 1, g):
                pairs -= exact[m]

            exact[g] = pairs

        # Prefix sums.
        prefix = [0] * (mx + 1)
        for g in range(1, mx + 1):
            prefix[g] = prefix[g - 1] + exact[g]

        ans = []

        for q in queries:

            # Binary search.
            l, r = 1, mx

            while l < r:
                mid = (l + r) // 2

                if prefix[mid] >= q + 1:
                    r = mid
                else:
                    l = mid + 1

            ans.append(l)

        return ans