class Solution:
    def countDaysTogether(self, arriveAlice: str, leaveAlice: str, arriveBob: str, leaveBob: str) -> int:
        days = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]
        def toDay(s):
            m = (ord(s[0]) - 48) * 10 + (ord(s[1]) - 48)
            d = (ord(s[3]) - 48) * 10 + (ord(s[4]) - 48)
            res = d
            for i in range(m - 1):
                res += days[i]
            return res
        aStart, aEnd = toDay(arriveAlice), toDay(leaveAlice)
        bStart, bEnd = toDay(arriveBob), toDay(leaveBob)
        start = max(aStart, bStart)
        end = min(aEnd, bEnd)
        return end - start + 1 if start <= end else 0