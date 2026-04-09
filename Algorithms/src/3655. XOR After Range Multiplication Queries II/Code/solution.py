class Solution:
    def xorAfterQueries(self, nums: List[int], queries: List[List[int]]) -> int:
        MOD = 10**9 + 7
        n = len(nums)
        limit = int(n ** 0.5) + 1

        # Fast exponentiation
        def mod_pow(base, exp):
            result = 1

            while exp > 0:
                if exp & 1:
                    result = (result * base) % MOD

                base = (base * base) % MOD
                exp >>= 1

            return result

        # Modular inverse
        def mod_inverse(x):
            return mod_pow(x, MOD - 2)

        small_queries = {}

        for l, r, k, v in queries:
            # Large k -> process directly
            if k >= limit:
                i = l
                while i <= r:
                    nums[i] = (nums[i] * v) % MOD
                    i += k
            else:
                if k not in small_queries:
                    small_queries[k] = []

                small_queries[k].append((l, r, v))

        # Process grouped small-k queries
        for k, group in small_queries.items():
            diff = [1] * n

            for l, r, v in group:
                diff[l] = (diff[l] * v) % MOD

                steps = (r - l) // k
                next_pos = l + (steps + 1) * k

                if next_pos < n:
                    diff[next_pos] = (diff[next_pos] * mod_inverse(v)) % MOD

            for i in range(n):
                if i >= k:
                    diff[i] = (diff[i] * diff[i - k]) % MOD

                nums[i] = (nums[i] * diff[i]) % MOD

        answer = 0

        for num in nums:
            answer ^= num

        return answer