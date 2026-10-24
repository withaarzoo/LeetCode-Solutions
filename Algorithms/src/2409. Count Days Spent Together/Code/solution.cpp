class Solution {
public:
    int countDaysTogether(string arriveAlice, string leaveAlice, string arriveBob, string leaveBob) {
        int days[] = {31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};
        auto toDay = [&](const string& s) {
            int m = (s[0] - '0') * 10 + (s[1] - '0');
            int d = (s[3] - '0') * 10 + (s[4] - '0');
            int res = d;
            for (int i = 0; i < m - 1; i++) res += days[i];
            return res;
        };
        int aStart = toDay(arriveAlice), aEnd = toDay(leaveAlice);
        int bStart = toDay(arriveBob), bEnd = toDay(leaveBob);
        int start = max(aStart, bStart);
        int end = min(aEnd, bEnd);
        return start <= end ? end - start + 1 : 0;
    }
};