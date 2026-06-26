class Solution:
    def countMajoritySubarrays(self, nums: List[int], target: int) -> int:

        n = len(nums)

        # Build prefix sums after transforming the array
        pref = [0]
        for x in nums:
            pref.append(pref[-1] + (1 if x == target else -1))

        # Coordinate compression
        values = sorted(set(pref))

        # Fenwick Tree
        bit = [0] * (len(values) + 2)

        # Insert one prefix sum
        def update(idx):
            while idx < len(bit):
                bit[idx] += 1
                idx += idx & -idx

        # Count prefix sums up to idx
        def query(idx):
            s = 0
            while idx > 0:
                s += bit[idx]
                idx -= idx & -idx
            return s

        ans = 0

        for x in pref:

            # Compressed index
            idx = bisect_left(values, x) + 1

            # Count smaller prefix sums
            ans += query(idx - 1)

            # Store current prefix sum
            update(idx)

        return ans