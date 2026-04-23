class Solution:
    def distance(self, nums):
        from collections import defaultdict
        
        n = len(nums)
        mp = defaultdict(list)

        # Group indices
        for i in range(n):
            mp[nums[i]].append(i)

        res = [0] * n

        for idx in mp.values():
            k = len(idx)

            prefix_sum = 0
            total_sum = sum(idx)

            for i in range(k):
                curr = idx[i]

                left = curr * i - prefix_sum
                right = (total_sum - prefix_sum - curr) - curr * (k - i - 1)

                res[curr] = left + right

                prefix_sum += curr

        return res