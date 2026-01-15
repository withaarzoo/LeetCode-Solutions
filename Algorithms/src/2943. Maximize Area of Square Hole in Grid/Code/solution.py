class Solution:
    def maximizeSquareHoleArea(self, n: int, m: int, hBars: List[int], vBars: List[int]) -> int:

        def get_max_gap(bars):
            bars.sort()
            max_len = 1
            cur_len = 1

            for i in range(1, len(bars)):
                if bars[i] == bars[i - 1] + 1:
                    cur_len += 1
                else:
                    cur_len = 1
                max_len = max(max_len, cur_len)

            return max_len

        h_gap = get_max_gap(hBars) + 1
        v_gap = get_max_gap(vBars) + 1

        side = min(h_gap, v_gap)
        return side * side
