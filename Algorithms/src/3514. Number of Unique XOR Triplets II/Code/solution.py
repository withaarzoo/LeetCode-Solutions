class Solution:
    def uniqueXorTriplets(self, nums: List[int]) -> int:
        MAX_XOR = 2048

        # Store whether a value exists.
        present = [False] * MAX_XOR
        for x in nums:
            present[x] = True

        # Initially only XOR = 0 is reachable.
        dp = [False] * MAX_XOR
        dp[0] = True

        # Pick exactly 3 values.
        for _ in range(3):
            nxt = [False] * MAX_XOR

            # Extend every reachable XOR.
            for cur in range(MAX_XOR):
                if not dp[cur]:
                    continue

                # Try every existing value.
                for v in range(MAX_XOR):
                    if present[v]:
                        nxt[cur ^ v] = True

            dp = nxt

        # Count unique XOR values.
        return sum(dp)