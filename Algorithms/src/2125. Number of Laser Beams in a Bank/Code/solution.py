from typing import List

class Solution:
    def numberOfBeams(self, bank: List[str]) -> int:
        ans = 0         # total beams
        prev = 0        # device count in previous non-empty row
        
        for row in bank:
            cnt = row.count('1')   # count devices in current row
            if cnt > 0:
                ans += prev * cnt  # beams between previous non-empty row and current row
                prev = cnt         # update previous
        return ans
