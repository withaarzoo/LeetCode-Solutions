func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
    days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
    toDay := func(s string) int {
        m := int(s[0]-'0')*10 + int(s[1]-'0')
        d := int(s[3]-'0')*10 + int(s[4]-'0')
        res := d
        for i := 0; i < m-1; i++ {
            res += days[i]
        }
        return res
    }
    aStart, aEnd := toDay(arriveAlice), toDay(leaveAlice)
    bStart, bEnd := toDay(arriveBob), toDay(leaveBob)
    start := max(aStart, bStart)
    end := min(aEnd, bEnd)
    if start <= end {
        return end - start + 1
    }
    return 0
}