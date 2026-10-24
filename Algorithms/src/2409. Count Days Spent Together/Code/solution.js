/**
 * @param {string} arriveAlice
 * @param {string} leaveAlice
 * @param {string} arriveBob
 * @param {string} leaveBob
 * @return {number}
 */
var countDaysTogether = function(arriveAlice, leaveAlice, arriveBob, leaveBob) {
    const days = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31];
    const toDay = (s) => {
        const m = (s.charCodeAt(0) - 48) * 10 + (s.charCodeAt(1) - 48);
        const d = (s.charCodeAt(3) - 48) * 10 + (s.charCodeAt(4) - 48);
        let res = d;
        for (let i = 0; i < m - 1; i++) res += days[i];
        return res;
    };
    const aStart = toDay(arriveAlice), aEnd = toDay(leaveAlice);
    const bStart = toDay(arriveBob), bEnd = toDay(leaveBob);
    const start = Math.max(aStart, bStart);
    const end = Math.min(aEnd, bEnd);
    return start <= end ? end - start + 1 : 0;
};