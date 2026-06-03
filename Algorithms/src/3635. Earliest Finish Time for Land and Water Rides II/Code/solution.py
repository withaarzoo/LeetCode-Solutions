class Solution:
    def earliestFinishTime(self, landStartTime: List[int], landDuration: List[int], waterStartTime: List[int], waterDuration: List[int]) -> int:

        from bisect import bisect_right

        # Computes the best answer when category A is taken first
        def solve(startA, durA, startB, durB):

            # Store (start, duration)
            rides = sorted(zip(startB, durB))

            m = len(rides)

            starts = [0] * m
            prefix_min_dur = [0] * m
            suffix_min_finish = [0] * m

            for i in range(m):
                starts[i] = rides[i][0]

                if i == 0:
                    prefix_min_dur[i] = rides[i][1]
                else:
                    prefix_min_dur[i] = min(
                        prefix_min_dur[i - 1],
                        rides[i][1]
                    )

            for i in range(m - 1, -1, -1):
                finish = rides[i][0] + rides[i][1]

                if i == m - 1:
                    suffix_min_finish[i] = finish
                else:
                    suffix_min_finish[i] = min(
                        suffix_min_finish[i + 1],
                        finish
                    )

            ans = float("inf")

            for s, d in zip(startA, durA):

                # Finish time of first ride
                finish1 = s + d

                # First ride with start > finish1
                pos = bisect_right(starts, finish1)

                if pos > 0:
                    ans = min(
                        ans,
                        finish1 + prefix_min_dur[pos - 1]
                    )

                if pos < m:
                    ans = min(
                        ans,
                        suffix_min_finish[pos]
                    )

            return ans

        return min(
            solve(
                landStartTime,
                landDuration,
                waterStartTime,
                waterDuration
            ),
            solve(
                waterStartTime,
                waterDuration,
                landStartTime,
                landDuration
            )
        )