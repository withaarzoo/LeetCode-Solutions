class Solution:
    def maximumTotalDamage(self, power: List[int]) -> int:
        from collections import Counter
        if not power:
            return 0
        cnt = Counter(power)
        vals = sorted(cnt.keys())  # sorted unique damage values
        m = len(vals)
        value_sum = [vals[i] * cnt[vals[i]] for i in range(m)]
        dp = [0] * m
        dp[0] = value_sum[0]
        import bisect
        for i in range(1, m):
            need = vals[i] - 3  # last allowed value <= need
            # find rightmost index <= need
            j = bisect.bisect_right(vals, need, 0, i) - 1
            take = value_sum[i] + (dp[j] if j >= 0 else 0)
            skip = dp[i - 1]
            dp[i] = max(skip, take)
        return dp[-1]
